package main

import (
	"log"
	"net/http"

	"github.com/school-sys-rest-api/routes"
	"github.com/school-sys-rest-api/services/db"
)

func main() {

	db.ConnectDB()
	//script.CompleteMigrateTables()
	log.Fatal(http.ListenAndServe("localhost:4000", routes.AllRoutes()))
}
