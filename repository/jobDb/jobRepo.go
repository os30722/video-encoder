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
	UpdateAndReturnCompletion(ctx context.Context, jobId int) (bool, error)
	GetJobFileAndOptions(ctx context.Context, jobId int) (*vo.JobFileAndOpts, error)
}
