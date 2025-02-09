package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetWatchlist(c *gin.Context, WatchlistCollection *mongo.Collection) {
	curs, err := WatchlistCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	var movies []interface{}

	for curs.Next(context.Background()) {
		var movie interface{}
		if err := curs.Decode(&movie); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		movies = append(movies, movie)
	}
	c.JSON(http.StatusOK, movies)
}
