package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mr-raj2001/students-api/internal/config"
)

func main() {
	//load config

	cfg := config.MustLoad()  //calling the MustLoad function from config package to get the configuration
	//databse setup
	//setup router
	//net http package for setting up http server
	router := http.NewServeMux()  //creating a new HTTP request multiplexer

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api"))
	})//handler function with response writer and request pointer
    

	//setup http server

	server := http.Server {
		Addr : cfg.Addr,        //setting the address from the config
		Handler: router,        //setting the handler to the router we created
	}

	fmt.Printf("Server started %s\n", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server")
	}

	
}