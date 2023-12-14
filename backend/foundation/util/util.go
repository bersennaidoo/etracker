package util

import (
	"context"
	"database/sql"
	"log"

	"github.com/bersennaidoo/etracker/backend/application/rest/auth"
	"github.com/bersennaidoo/etracker/backend/infrastructure/storage/pgstore"
	"github.com/lib/pq"
)

func CreateUserInDb(db *sql.DB) {

	ctx := context.Background()
	querier := pgstore.New(db)

	log.Println("Creating user@user...")
	hashPwd := auth.HashPassword("bersen")

	_, err := querier.CreateUsers(ctx, pgstore.CreateUsersParams{
		UserName:     "bersen@g.com",
		PasswordHash: hashPwd,
		Name:         "bersen",
	})

	// This is interesting to look at, the sql/pq library recommends we use
	// this pattern to understand errors. We could use the ErrorCode directly
	// or look for the specific type. We know we'll be violating unique_violation
	// if our user already exists in the database
	if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
		log.Println("Dummy User already present")
		return
	}

	if err != nil {
		log.Println("Failed to create user:", err)
	}
}
