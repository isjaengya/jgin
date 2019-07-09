package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	Config "tebu_go/api/config"
)

var mysqldb *sql.DB

func MysqlInit(){

 	var err error
	config := Config.Conf
	mysqlUser := config.GetString("mysql.user")
	mysqlPassword := config.GetString("mysql.password")
	mysqlHost := config.GetString("mysql.host")
	mysqlPort := config.GetString("mysql.port")
	mysqlDatabase := config.GetString("mysql.database")

	mysqlMaxOpenConns := config.GetInt("mysql.maxopenconns")
	mysqlMaxIdleConns := config.GetInt("mysql.maxidleconns")
	// user:password@(host:port)/dbname
	mysqlS := fmt.Sprintf("%s:%s@(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
    //Mysqldb, err = sql.Open("mysql", "root:root@/yinyu_dev")
    mysqldb, err = sql.Open("mysql", mysqlS)

    if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}

    mysqldb.SetMaxOpenConns(mysqlMaxOpenConns)
    mysqldb.SetMaxIdleConns(mysqlMaxIdleConns)

    err = mysqldb.Ping()
    if err != nil {
		log.Fatal(err)
	}

}

func GetMysqlClient() *sql.DB {
	return mysqldb
}
