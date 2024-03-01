package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tc-server/config"
	"tc-server/controller"
	"tc-server/db"
)

// Init will initialize the Gin server and all
// accompanying databases like Redis Cache and MongoDB
// In addition this function handles the application
// of routes used on this Gin instance
func Init(config *config.FullConfig) {
	gin.SetMode(config.Gin.Env)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.Gin.Origins
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("GET", "POST")
	corsConfig.AddAllowHeaders(
		"Content-Type", "X-XSRF-TOKEN", "Accept",
		"Origin", "X-Requested-With", "Authorization",
		"Set-Cookie", "Access-Control-Allow-Origin")

	router := gin.New()
	// middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(corsConfig))

	// db & cache
	redis, err := db.InitRedis(&config.Cache)
	if err != nil {
		panic("failed to establish connection with redis cache: " + err.Error())
	}
	mongo, err := db.InitMongo(&config.Mongo)
	if err != nil {
		panic("failed to establish connection with mongo database: " + err.Error())
	}

	gc := controller.GlobalController{
		Config: config,
		Mongo:  mongo,
		Redis:  redis,
	}

	// apply routes
	gc.ApplyAccountRoutes(router)

	if err := router.Run(":" + config.Gin.Port); err != nil {
		panic("failed to start gin: " + err.Error())
	}
}
