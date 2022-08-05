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
	log.Fatal(http.ListenAndServe("192.168.10.40:4000", routes.AllRoutes()))
}
