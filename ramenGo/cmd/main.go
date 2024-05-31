package main

import (
	"database/sql"
	"fmt"

	"github.com/ramenGo/infra/util"
)

func connectWithDb(configs *util.Configs) *sql.DB {
	conURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", configs.MYSQL_ROOT_PASSWORD,
		configs.MYSQL_PASSWORD, "mysql", configs.MYSQL_PORT, configs.MYSQL_DATABASE)

	db, err := sql.Open("mysql", conURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	config := util.GetConfigs()
	db := connectWithDb(config)
	defer db.Close()

}
