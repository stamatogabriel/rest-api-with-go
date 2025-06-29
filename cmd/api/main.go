package main

import (
	"backend/db"
	"backend/handlers"
	"backend/services"
	"context"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Models services.Models
}

func main() {
	mongoClient, err := db.ConnectToMongoDB()

	if err != nil {
		log.Panic("Failed to connect to MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Panic("Failed to disconnect from MongoDB:", err)
		}
	}()

	services.New(mongoClient)
	log.Println("Services running on port", 8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}
