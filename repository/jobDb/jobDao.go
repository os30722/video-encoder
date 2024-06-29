package jobDb

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cloud/encoder/vo"
	"github.com/jackc/pgx/v5"
)

func (jo JobDao) CreateJob(ctx context.Context) (int, error) {
	var db = jo.db

	var jobId int
	err := db.QueryRow(ctx, "insert into job(status) values($1) returning job_id", "RUNNING").Scan(&jobId)
	if err != nil {
		return 0, err
	}

	return jobId, nil
}

func (jo JobDao) UpdateProcesses(ctx context.Context, jobId int, processes []vo.Process) error {
	var db = jo.db
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "update job set total_processes=$1 where job_id=$2", len(processes), jobId)
	if err != nil {
		return err
	}

	_, err = tx.CopyFrom(ctx, pgx.Identifier{"process"},
		[]string{"job_id", "part_name", "total_parts"},
		pgx.CopyFromSlice(len(processes), func(i int) ([]any, error) {
			return []any{processes[i].JobId, processes[i].PartName, processes[i].TotalPart}, nil
		}))

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return err
}

func (jo JobDao) UpdateAndReturnCompletion(ctx context.Context, jobId int, partName string) (bool, bool, error) {
	var db = jo.db
	var completed, jobCompleted bool

	tx, err := db.Begin(ctx)
	if err != nil {
		return completed, jobCompleted, err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, "update process set completed_parts=completed_parts+1 where job_id=$1 and part_name=$2 returning completed_parts=total_parts",
		jobId, partName).Scan(&completed)
	if err != nil {
		return completed, jobCompleted, err
	}

	if completed {
		err = tx.QueryRow(ctx, "update job set completed_processed=completed_processed+1 where job_id=$1' returning completed_processes=total_processes",
			jobId).Scan(&jobCompleted)
		if err != nil {
			return completed, jobCompleted, err
		}

		_, err = tx.Exec(ctx, "delete from process where job_id=$1 and part_name=$2", jobId, partName)
		if err != nil {
			return completed, jobCompleted, err
		}
	}

	return completed, jobCompleted, nil
}

func (jo JobDao) GetOutputs(ctx context.Context, templateId int) (*vo.EncodeOuputStruct, error) {
	var db = jo.db

	var info vo.EncodeOuputStruct
	var outputs string

	err := db.QueryRow(ctx, "select outputs from job_template where template_id=$1", templateId).
		Scan(&outputs)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(strings.NewReader(outputs)).Decode(&info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
