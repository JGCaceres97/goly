package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jgcaceres97/goly/app/database"
	"github.com/jgcaceres97/goly/app/server"
	"github.com/jgcaceres97/goly/app/settings"
)

func init() {
	err := settings.Setup()
	if err != nil {
		panic(err)
	}

	database.Connect()
}

func main() {
	app := server.SetupAndListen()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("\nGracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	fmt.Println("\t- Closing DB connection...")
	DB, _ := database.DB.DB()
	defer DB.Close()

	fmt.Println("\nServer was successfully shutdown.")
}
