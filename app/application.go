package app

import (
	"fmt"
	"go-backend/logger"
	"go-backend/utils/os_utils"
)

func StartApplication() {
	router := setUpRouter()

	port := os_utils.GetEnvOrDefault("PORT", ":8080")
	logger.Info(fmt.Sprintf("About to start the application on port %s...", port))
	router.Run(port)
}
