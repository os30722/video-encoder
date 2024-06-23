package jobDb

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cloud/encoder/vo"
)

func (jo JobDao) UpdateAndReturnCompletion(ctx context.Context, jobId int) (bool, error) {
	var db = jo.db

	var completed bool
	err := db.QueryRow(ctx, "update process set completed_parts=completed_parts+1 where job_id=$1 returning completed_parts=total_parts",
		jobId).Scan(&completed)
	if err != nil {
		return false, err
	}

	return completed, nil
}

func (jo JobDao) GetJobFileAndOptions(ctx context.Context, jobId int) (*vo.JobFileAndOpts, error) {
	var db = jo.db

	var info vo.JobFileAndOpts
	var optStr string
	err := db.QueryRow(ctx, "select opts,name,file_name,output from process join job on job.job_id = process.job_id where process.job_id = $1", jobId).
		Scan(&optStr, &info.JobName, &info.FileName, &info.OutputFormats)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(strings.NewReader(optStr)).Decode(&info.Opts)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
