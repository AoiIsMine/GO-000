package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	username = "root"
	password = "xiao13579"
	network  = "tcp"
	server   = "localhost"
	port     = 3306
	database = "test"
)

//Query db query
func Query(id int) (interface{}, error) {
	fmt.Println("mydb.Query start")
	db, initErr := dbInit()
	if initErr != nil {
		// return nil, initErr	//常规err无堆栈,附加信息
		return nil, errors.Wrap(initErr, "db init error")
	}

	var name string
	row := db.QueryRow("select name from users where id =?", id)
	queryErr := row.Scan(&name)

	return name, queryErr
	// return name, errors.Wrap(queryErr, "query error")
}

func getMysqlConfig() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, network, server, port, database)
}

func dbInit() (*sql.DB, error) {
	return sql.Open("mysql", getMysqlConfig())
	// db, err := sql.Open("mysql", getMysqlConfig())
	// if err != nil {
	// 	//附加msg
	// 	return nil, err
	// }
	// return db, nil
}

func dbClose(db *sql.DB) error {
	return db.Close()
	// err := db.Close()
	// if err != nil {
	// 	return errors.New("dbClose err ")
	// }
	// return nil
}

// type User struct {
// 	ID   int
// 	Name string
// }
