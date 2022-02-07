package main

import (
	"flag"

	"github.com/tecnologer/dicegame/server/models"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 8088, "Port to expose the server")
	flag.Parse()

	models.InitGames()
	models.NewServer(port)
}
