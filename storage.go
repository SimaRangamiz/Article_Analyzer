package main

import (
	"context"
	"time"
)

func SaveArticle(article Article) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := DB.InsertOne(ctx, article)
	if err != nil {
		return err
	}
	return nil
}