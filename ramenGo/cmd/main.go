package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ramenGo/api"
	databases "github.com/ramenGo/domain/database"
	"github.com/ramenGo/domain/usecases"
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

	brothDb := databases.NewBrothDB(db)
	proteinDb := databases.NewProteinDB(db)
	orderDb := databases.NewOrderDB(db)

	listBrothsUc, err := usecases.NewListBrothsUseCase(brothDb)
	if err != nil {
		panic(err)
	}
	listProteinsUc, err := usecases.NewListProteinsUseCase(proteinDb)
	if err != nil {
		panic(err)
	}
	createOrderUc, err := usecases.NewCreateOrderUseCase(orderDb)
	if err != nil {
		panic(err)
	}

	wb := api.NewWebServer()
	wb.Router.Use(api.AuthMiddleware(config.API_KEY))
	wb.AddNewGetBrothsRoute("/broths", listBrothsUc)
	wb.AddNewGetProteinsRoute("/proteins", listProteinsUc)
	wb.AddNewCreateOrderRoute("/orders", createOrderUc)
	wb.Router.Run() // Default port 8080
}
