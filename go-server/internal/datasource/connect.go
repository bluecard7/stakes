package datasource

import (
	"database/sql"
	"fmt"
	"stakes/internal/config"

	_ "github.com/lib/pq" // psql driver
)

var _ = sql.Drivers

func PrepDB() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=verify-full",
		config.Get("psql.user"),
		config.Get("psql.password"),
		config.Get("psql.host"),
		config.Get("psql.dbName"),
	)

	fmt.Println(connStr)
}
