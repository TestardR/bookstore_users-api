package app

import (
	"github.com/TestardR/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapUrls()
	logger.Info("about to start the application")
	router.Run(":8080")
}
