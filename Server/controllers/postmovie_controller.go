package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostMovie(c *gin.Context, WatchlistCollection *mongo.Collection) {
	var movie interface{}

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := WatchlistCollection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	c.JSON(http.StatusBadRequest, res)

}
