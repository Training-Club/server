package controller

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"tc-server/config"
)

type GlobalController struct {
	Config *config.FullConfig
	Mongo  *mongo.Client
	Redis  *redis.Client
}
