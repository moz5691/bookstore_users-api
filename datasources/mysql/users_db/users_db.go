package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	var err error
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)
	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	//mysql.SetLogger()

	log.Println("database successfully configured")
}
