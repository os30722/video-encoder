package database

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "oishik"
	dbname   = "encoder"
)

var db *pgxpool.Pool = nil

// To get an instance of postgres connection

func GetPostgres() (*pgxpool.Pool, error) {

	var err error
	db, err = pgxpool.New(context.Background(), "postgresql://"+host+":"+strconv.Itoa(port)+"/"+dbname+"?user="+user+"&password="+password)
	if err != nil {
		return nil, err
	}

	log.Println("Established Database Connection...")
	return db, nil
}
