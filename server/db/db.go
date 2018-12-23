package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"server/conf"
)

type DBContext struct {
	db *sqlx.DB
}

func NewDBContext(config conf.MySQL) (*DBContext, error) {
	connString := config.User + ":" + config.Password + "@(" +
		config.Host + ":" + config.Port + ")/" +
		config.Database + "?parseTime=true"

	database, err := sqlx.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &DBContext{database}, nil
}

func (ctx *DBContext) SetupTables() error {
	sqlFile, err := ioutil.ReadFile("./db/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read db/schema.sql: %v", err)
	}

	sqlString := string(sqlFile)
	_, err = ctx.db.Exec(sqlString)
	if err != nil {
		return fmt.Errorf("failed to execute table setup SQL: %v", err)
	}

	return nil
}

func (ctx *DBContext) Close() error {
	return ctx.db.Close()
}