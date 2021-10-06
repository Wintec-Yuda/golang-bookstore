package main

import (
	"log"
	"net/http"

	"github.com/Wintec-Yuda/golang-bookstore.git/pkg/routes"
	"github.com/Wintec-Yuda/golang-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main()  {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}