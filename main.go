package main

import (
	"Restaurante/routers"
	"Restaurante/utils"
	"log"
	"net/http"
)

func main() {
	utils.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
