package main

import (
	"log"
	"net/http"

	"my-sqlite-app/db"
	"my-sqlite-app/handlers"
)

func main() {
	database := db.InitDB("my.db")
	defer database.Close()

	http.HandleFunc("/query", handlers.UserHandler(database))

	log.Println("サーバ起動中 :8080")
	http.ListenAndServe(":8080", nil)
}
