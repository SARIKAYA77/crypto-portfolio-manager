package main

import (
	"crypto-portfolio/app/api"
	"crypto-portfolio/app/db"
	"crypto-portfolio/app/service"
	"database/sql"
	"log"
	"net/http"
)

func main() {
	currencyService := service.NewCurrencyService()
	router := api.InitRoutes(currencyService)

	dbCon := db.InitDB()
	defer func(dbCon *sql.DB) {
		err := dbCon.Close()
		if err != nil {
			log.Println(err)
		}
	}(dbCon)

	log.Fatal(http.ListenAndServe(":8090", router))
}
