package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBHost = "DB_HOST"
	DBPort = "DB_PORT"
	DBName = "DB_NAME"
)

type DBConfig struct {
	DatabaseURL  string
	DatabaseName string
}

var DB *mongo.Database

func getConfig() *DBConfig {
	return &DBConfig{
		DatabaseURL:  fmt.Sprintf("%v:%v", os.Getenv(DBHost), os.Getenv(DBPort)),
		DatabaseName: os.Getenv(DBName),
	}
}

func ConnectDB() error {
	conf := getConfig()
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.DatabaseURL))
	if err != nil {
		return err
	}

	DB = db.Database(conf.DatabaseName)
	return nil
}
