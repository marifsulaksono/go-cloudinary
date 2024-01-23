package model

import (
	"context"
	"go-cloudinary/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	ImageURL string             `json:"image_url" bson:"image:url" form:"image"`
}

func (i *Image) InsertImage(ctx context.Context) error {
	data := bson.D{
		{Key: "image_url", Value: i.ImageURL},
	}
	_, err := config.DB.Collection("images").InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
