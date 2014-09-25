package db

import (
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	user     = "root"
	password = "momoc1oZ"
	database = "totec"
)

var con *sqlx.DB
var dbmap gorp.DbMap

func init() {
	var err error
	con, _ = sqlx.Connect("mysql", user+":"+password+"@/"+database)

	if err != nil {
		log.Fatal(err.Error())
	}

}

// ExecWithArgs function insert record
func ExecWithArgs(sql string, args ...interface{}) error {
	_, err := con.Exec(sql, args)

	if err != nil {
		return err
	}

	return nil
}

// Exec function
func Exec(sql string) error {
	_, err := con.Exec(sql)

	if err != nil {
		return err
	}

	return nil
}

// MysqlRead function
func MysqlRead(sql string, args ...interface{}) (*sqlx.Rows, error) {
	rows, err := con.Queryx(sql, args)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
