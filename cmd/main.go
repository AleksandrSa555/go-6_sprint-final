package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Server: ", log.LstdFlags)

	serv := server.NewServer(logger)

	logger.Println("Сервер запускается на порту :8080")

	err := serv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
		return
	}
}
