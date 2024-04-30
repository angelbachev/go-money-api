package main

import (
	"log"
	"os"

	auth_infrastructure "github.com/angelbachev/go-money-api/infrastructure/domain/auth"
	"github.com/angelbachev/go-money-api/infrastructure/factory"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	addr := os.Getenv("LISTEN_ADDRESS")
	if addr == "" {
		addr = ":80"
	}

	authService := *auth_infrastructure.NewJWTAuth(os.Getenv("JWT_SECRET"))

	dbFactory, err := factory.NewMySQLRepositoryFactory(os.Getenv("MYSQL_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	commandHandlerFactory := factory.NewCommandHandlerFactory(dbFactory, authService)
	queryHandlerFactory := factory.NewQueryHandlerFactory(dbFactory)
	actionFactory := factory.NewActionFactory(commandHandlerFactory, queryHandlerFactory)

	server := rest.NewServer(addr, authService, actionFactory.All())
	server.Run()
}
