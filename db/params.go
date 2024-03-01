package db

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoParams struct {
	Client         *mongo.Client
	DBName         string
	CollectionName string
}

type RedisParams struct {
	RedisClient *redis.Client
}
