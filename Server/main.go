package main

import (
	router "Server/routes"
	"Server/server"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
)

type Movie struct {
	Title string `bson:"title"`
	Date  int    `bson:"date"`
}

func main() {

	s, err := server.InitializeServer()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(s)
	}
	s.Router.Use(cors.Default())
	router.ApplyRoutes(s.Router)
	s.Router.Run()
}
