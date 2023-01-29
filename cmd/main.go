package main

import (
	"log"
	"net/http"

	"github.com/TonyPham-Dev/Golang-Crud/pkg/configs"
	"github.com/TonyPham-Dev/Golang-Crud/pkg/routers"
	"github.com/gorilla/mux"
)

func main() {
	var port string = ":8001"
	client, ctx, context, err := configs.ConnectDatabase("mongodb+srv://TonyPham:Cong2001!@cluster0.daknur6.mongodb.net/?retryWrites=true&w=majority")

	if err != nil {
		panic(err)
	}

	defer configs.CloseDatabase(client, ctx, context)
	configs.PingDatabases(client, ctx)

	router := mux.NewRouter()
	log.Println("Server running is port", port)
	log.Fatal(http.ListenAndServe(port, routers.RegisterRouterBooks(router)))
}
