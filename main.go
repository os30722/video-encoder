package main

import (
	"context"
	"log"

	"github.com/cloud/encoder/database"
	"github.com/cloud/encoder/repository/jobDb"
	"github.com/cloud/encoder/runner"
)

func failOnError(err error) {
	if err != nil {
		log.Panicf(err.Error())
	}
}

func main() {
	ctx := context.Background()

	db, err := database.GetPostgres()
	if err != nil {
		panic(err)
	}

	jobDao := jobDb.GetJobDao(db)

	if err = runner.Start(ctx, jobDao); err != nil {
		panic(err)
	}
}
