package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbURL string) {

	var err error
	if Dbdriver == "postgres" {
		server.DB, err = gorm.Open(Dbdriver, DbURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error:", err)

		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port " + addr)
	log.Fatal(http.ListenAndServe(":"+addr, server.Router))
}
