package jobDb

import (
	"context"

	"github.com/cloud/encoder/vo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type JobDao struct {
	db *pgxpool.Pool
}

func GetJobDao(db *pgxpool.Pool) *JobDao {
	return &JobDao{
		db: db,
	}
}

type JobRepo interface {
	CreateJob(ctx context.Context, templateId int) (int, error)
	UpdateProcesses(ctx context.Context, jobId int, processes []vo.Process) error
	UpdateAndReturnCompletion(ctx context.Context, jobId int, partName string) (bool, bool, error)
	GetOutputs(ctx context.Context, templateId int) (*vo.EncodeOutputStruct, error)
	GetTemplate(ctx context.Context, jobId int) (*vo.JobTemplate, error)
}
