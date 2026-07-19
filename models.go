package main

import ("time"
        "go.mongodb.org/mongo-driver/bson/primitive"
)


type ArticleResponse struct {
	Title     string    `json:"title"`      
	Tags      []string  `json:"tags"`       
	CreatedTime time.Time `json:"created_at"` 
}

type Article struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Body      string             `bson:"body" json:"body"`  
	Tags      []string           `bson:"tags" json:"tags"`  
	CreatedTime time.Time       `bson:"created_at" json:"created_at"`
}