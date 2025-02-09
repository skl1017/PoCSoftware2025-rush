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
		// Utilise un log mais n'arrête pas le programme
		log.Println("Error binding JSON:", err)
		// Répond avec une erreur HTTP
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insère le film dans la collection MongoDB
	res, err := WatchlistCollection.InsertOne(context.Background(), movie)
	if err != nil {
		// Log l'erreur mais ne termine pas le programme
		log.Println("Error inserting movie:", err)
		// Répond avec une erreur HTTP
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to insert movie"})
		return
	}

	// Optionnellement, tu peux envoyer une réponse de succès avec des données
	c.JSON(http.StatusOK, gin.H{"message": "Movie added successfully", "id": res.InsertedID})
}
