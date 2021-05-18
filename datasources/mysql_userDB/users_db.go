package mysql_userDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	DBClient *sql.DB
	// WARNING: replace values with environment variables!!!
	mysql_user     = os.Getenv("MYSQL_USER")
	mysql_password = os.Getenv("MYSQL_PASSWORD")
	mysql_host     = os.Getenv("MYSQL_HOST") //point to docker MYSQL
	mysql_db       = os.Getenv("MYSQL_DB")

	//WARNING do not log these values anywhere
)

func init() {

	//user:password@TCP(host)/schema
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysql_user, mysql_password, mysql_host, mysql_db,
	)
	var err error

	DBClient, err = sql.Open("mysql", datasource)
	if err != nil {
		panic("Error while connecting to DB")
	}

	// mysql.SetLogger() TODO set logger
	errPing := DBClient.Ping()
	if errPing != nil {
		fmt.Printf("Error in ping: %v\n", errPing)
	}
}
