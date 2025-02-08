package main

import (
	"Server/server"
	"fmt"
	"log"
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
	s.Router.Run()
}
