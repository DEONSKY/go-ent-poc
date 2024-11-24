package database

import (
	"database/sql"
	"log"

	"ent-mysql/ent"
	_ "ent-mysql/ent/runtime"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DBConn *ent.Client

// Open new connection
func ConnectDb() {
	db, err := sql.Open("pgx", "postgresql://myuser:mypassword@127.0.0.1/mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	DBConn = ent.NewClient(ent.Driver(drv))
}
