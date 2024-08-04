package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bndrmrtn/fate_fighter/database/queries"
	"github.com/bndrmrtn/fate_fighter/pkg/configs"
)

var DB *queries.Queries

func MustConnect() {
	db, err := sql.Open("mysql", configs.Env.DB.GetUrl())
	if err != nil {
		panic(err)
	}

	DB = queries.New(db)
}
