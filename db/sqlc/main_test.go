package db

import (
	"context"
	_ "database/sql"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQuery *Queries

const dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
const dbDriver = "postgres"

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource) //
	//conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQuery = New(conn)
	os.Exit(m.Run())
}
