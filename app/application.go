package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-backend/logger"
	"go-backend/utils/os_utils"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURLs()

	port := os_utils.GetEnvOrDefault("PORT", ":8080")
	logger.Info(fmt.Sprintf("About to start the application on port %s...", port))
	router.Run(port)
}
