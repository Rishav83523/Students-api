package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mr-raj2001/students-api/internal/config"
	"github.com/mr-raj2001/students-api/internal/http/handlers/student"
	"github.com/mr-raj2001/students-api/internal/storage/sqlite"
)

func main() {
	//load config

	cfg := config.MustLoad()  //calling the MustLoad function from config package to get the configuration
	//databse setup

	storage, err := sqlite.New(cfg)  //we can switch to any other storage by changing this line only as we are using interface
	if err != nil {
		log.Fatalf("failed to setup storage: %s", err.Error())
	}

	slog.Info("Storage setup completed")
	//setup router
	//net http package for setting up http server
	router := http.NewServeMux()  //creating a new HTTP request multiplexer

	router.HandleFunc("POST /api/students",student.New(storage))//handler function with response writer and request pointer
	router.HandleFunc("GET /api/students/{id}",student.GetById(storage))
	router.HandleFunc("GET /api/students",student.GetStudentList(storage))
	//we cleaned it by moving the handler function to internal/http/handlers/student/student.go file
    
    
	//setup http server

	server := http.Server {
		Addr : cfg.Addr,        //setting the address from the config
		Handler: router,        //setting the handler to the router we created
	}

	slog.Info("Server started ", slog.String("address", cfg.Addr))

	//graceful shutdown using os signal package goroutine and channel 

	done := make(chan os.Signal, 1)   //channel to listen for termination signals with size 1 buffer channel
     
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)  //notifying the channel on receiving interrupt or termination signals
	go func() {
        err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server")
	}
	}() 



	<- done   //blocking main goroutine until a signal is received

	//we will write code for graceful shutdown here after receiving the signal and unblocking the main goroutine

	slog.Info("Server Stopped")

	//we are using context package to create a context with timeout for graceful shutdown otherwise it may wait indefinitely

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)  //creating a context with timeout for graceful shutdown , empty starting point 
    defer cancel()

	err = server.Shutdown(ctx)  //gracefully shutting down the server and recieve context if not cleased in 5 second and throw error for us 
	//sometimes infinite wait so we use context with timeout
	

	if err != nil {
		slog.Error("Server Shutdown Failed", slog.String("error", err.Error()))
	}

	slog.Info("Server Exited Properly")  //slog is structured logging package in go
}