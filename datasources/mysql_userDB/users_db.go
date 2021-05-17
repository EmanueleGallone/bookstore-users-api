package mysql_userDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
//mysql_user = "root"
//mysql_user_passwords = "root"
//mysql_host = "172.17.0.2"
)

var (
	DBClient *sql.DB
	// WARNING: replace values with environment variables!!!
	mysql_user     = "root"       //os.Getenv(mysql_user)
	mysql_password = "root"       //os.Getenv(mysql_password)
	mysql_host     = "172.17.0.2" // os.Getenv(mysql_host) point to docker MYSQL
	mysql_db       = "users_db"   //os.Getenv()
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
