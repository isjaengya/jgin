package service

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"log"
	tebuConfig "tebu_go/api/config"
)

var Mysqldb *sql.DB

func MysqlInit(){

 	var err error
	config := tebuConfig.LoadConfig()
	mysqlMaxOpenConns := config.GetInt("mysql.maxopenconns")
	mysqlMaxIdleConns := config.GetInt("mysql.maxidleconns")
	// user:password@/dbname
    Mysqldb, err = sql.Open("mysql", "root:root@/yinyu_dev")

    if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}


    Mysqldb.SetMaxOpenConns(mysqlMaxOpenConns)

    Mysqldb.SetMaxIdleConns(mysqlMaxIdleConns)

    err = Mysqldb.Ping()
    if err != nil {
		log.Fatal(err)
	}

}
