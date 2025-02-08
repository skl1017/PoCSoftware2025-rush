package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	Router              *gin.Engine       `bson:"router"`
	Client              *mongo.Client     `bson:"client"`
	WatchlistCollection *mongo.Collection `bson:"watchlist"`
	TMDBApiKey          string
}

func initializeClient() (*mongo.Client, error) {
	godotenv.Load()
	MONGO_URI := os.Getenv("MONGODB_URI")
	println("MONGODB_URI: " + MONGO_URI)
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to build client")
		return nil, err
	}
	fmt.Println("Client built!, Trying to connect to MongoDB...")

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
		return nil, err
	} else {
		fmt.Println("SUCCESS: Connection established!")
	}
	return client, nil
}

func initializeRouter() *gin.Engine {
	return gin.Default()
}

func initializeWatchlistCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("db").Collection("watchlist")
}

func initializeTMDBApiKey() string {
	godotenv.Load()
	return os.Getenv("TMDB_API_KEY")
}

func InitializeServer() (*Server, error) {
	var server Server

	server.Router = initializeRouter()
	client, err := initializeClient()
	if err != nil {
		log.Fatalf("ERROR: Error during router initialization: %v", err)
		return nil, err
	}
	server.WatchlistCollection = initializeWatchlistCollection(client)
	server.Client = client
	server.TMDBApiKey = initializeTMDBApiKey()
	return &server, nil
}
