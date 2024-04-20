package main

import (
	"log"
	"os"

	"github.com/angelbachev/go-money-api/api"
	"github.com/angelbachev/go-money-api/storage"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	addr := os.Getenv("LISTEN_ADDRESS")
	if addr == "" {
		addr = ":8089"
	}

	connStr := os.Getenv("MYSQL_CONNECTION_STRING")
	store, err := storage.NewMySQLStore(connStr)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(addr, store)
	server.Run()
}
