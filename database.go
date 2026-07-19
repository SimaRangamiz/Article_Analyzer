package main

import ("fmt"
        "context"
		"log"
		"time"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Collection

func ConnectDB() {
	mongoAddres := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddres))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("tag_extractor").Collection("articles")
	
	fmt.Println("Successfully connected to MongoDB!")
	
}