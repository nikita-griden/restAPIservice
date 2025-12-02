package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logs := log.New(os.Stdout, "", log.LstdFlags)
	log.Fatal(server.Router(logs).Serv.ListenAndServe())
}
