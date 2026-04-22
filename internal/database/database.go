package database

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
	"github.com/rms-diego/book-my-session/pkg/config"
)

var Db *goqu.Database

func Init() error {
	dialect := goqu.Dialect("postgres")
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Env.DB_HOST,
		config.Env.DB_PORT,
		config.Env.DB_USER,
		config.Env.DB_PASSWORD,
		config.Env.DB_NAME,
	)

	pgDb, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	Db = dialect.DB(pgDb)
	if _, err := Db.Exec(`SELECT 'PING'`); err != nil {
		return err
	}

	return nil
}
