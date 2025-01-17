package main

import (
	"echo_ifElse/pkg/handler"
	"echo_ifElse/pkg/repository"
	"echo_ifElse/pkg/service"
	"echo_ifElse/server"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5434",
		Username: "postgres",
		Password: "5432",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := server.NewServer(handlers)

	if err := server.Run(); err != nil {
		log.Fatal("shutDowning")
	}
}
