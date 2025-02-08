package main

import (
	router "Server/routes"
	"Server/server"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
)

func main() {

	s, err := server.InitializeServer()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(s)
	}
	s.Router.Use(cors.Default())
	router.ApplyRoutes(s)
	s.Router.Run()
}
