package data

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func randomUUID() uuid.UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		id = uuid.Nil
	}
	log.Printf("UUID: %s\n", id)
	return id
}

func query(db *sql.DB, queryStr string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(queryStr, args...)
	if err != nil {
		rows = nil
	}
	log.Println("Query:", queryStr)
	return rows
}

func prepSQLStmt(db *sql.DB, queryStr string) *sql.Stmt {
	stmt, err := db.Prepare(queryStr)
	if err != nil {
		stmt = nil
	}
	log.Println("Prepped SQL queryStr", queryStr, stmt)
	return stmt
}

func execSQLStmt(stmt *sql.Stmt, args ...interface{}) {
	if stmt == nil {
		return
	}
	result, _ := stmt.Exec(args...)
	log.Println("Exec'd stmt", stmt, result)
}
