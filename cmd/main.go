package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/murali-muniyan/go-book-store/pkg/config"
	"github.com/murali-muniyan/go-book-store/pkg/router"
)

func main() {
	config.InitDB()
	// models.Migrate()

	rtr := mux.NewRouter()

	router.RegisterRouter("/books", rtr)

	log.Fatal(http.ListenAndServe(":8000", rtr))
}
