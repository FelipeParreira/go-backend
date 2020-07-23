package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go-backend/controllers/ping"
	Servers "go-backend/controllers/servers"
	"go-backend/docs"
	"go-backend/utils/os_utils"
)

// @title RedCoins Swagger API
// @description Swagger API for RedCoins project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email felipeparreirafb@gmail.com

// @license.name MIT
// @license.url https://github.com/FelipeParreira/go-backend/blob/master/LICENSE

// @schemes http https
func setUpRouter() *gin.Engine {
	router := gin.Default()

	// programmatically set swagger info
	docs.SwaggerInfo.Version = os_utils.GetEnvOrDefault("API_VERSION", "1.0")
	docs.SwaggerInfo.BasePath = "/api/v" + docs.SwaggerInfo.Version
	docs.SwaggerInfo.Host = os_utils.GetEnvOrDefault("HOST", "localhost") + os_utils.GetEnvOrDefault("PORT", "8080")

	// Docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API version
	api := router.Group(docs.SwaggerInfo.BasePath)

	// health check
	api.GET("/health_check", ping.HealthCheck)

	// public endpoints
	servers := api.Group("/servers")
	servers.GET("/summary", Servers.HandleGetSummary)

	return router
}
