package main

import ("time"
        "go.mongodb.org/mongo-driver/bson/primitive"
)


type Article_Response struct {
	Title     string    `json:"title"`      
	Tags      []string  `json:"tags"`       
	Created_Time time.Time `json:"created_at"` 
}

type Article struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Body      string             `bson:"body" json:"body"`  
	Tags      []string           `bson:"tags" json:"tags"`  
	Created_Time time.Time       `bson:"created_at" json:"created_at"`
}