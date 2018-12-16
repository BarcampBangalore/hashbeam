package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

func Init(connString string) (*DatabaseContext, error) {
	database, err := sqlx.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &DatabaseContext{database}, nil
}

type DatabaseContext struct {
	DB *sqlx.DB
}

func (ctx *DatabaseContext) SetupTables() error {
	sqlFile, err := ioutil.ReadFile("./db/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read db/schema.sql: %v", err)
	}

	sqlString := string(sqlFile)
	_, err = ctx.DB.Exec(sqlString)
	if err != nil {
		return fmt.Errorf("failed to execute table setup SQL: %v", err)
	}

	return nil
}
