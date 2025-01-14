package main

import (
	"echo_ifElse"
	"echo_ifElse/pkg/repository"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	_, err := repository.NewPostgresDB(repository.Config{
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

	if err := echo_ifElse.Run(); err != nil {
		log.Fatal("shutDowning")
	}
}
