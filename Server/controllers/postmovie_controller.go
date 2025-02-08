package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	MovieId int `json:"movie"`
}

func PostMovie(c *gin.Context) {
	var body RequestBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
