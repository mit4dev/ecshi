package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	api "github.com/mit4dev/ecshi/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDB() (*mongo.Database, error) {
	conn := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(conn)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database("ecshi")
	return db, nil
}

func main() {
	db, err := connectToDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	api := api.NewApi(db, r)
	api.RegisterRoutes()

	r.Run()
}
