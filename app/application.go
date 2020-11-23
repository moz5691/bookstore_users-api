package app

import (
	"github.com/gin-gonic/gin"
	"github.com/moz5691/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Starting app...")
	if err := router.Run(":8080"); err != nil {
		logger.Error("Fail to start", err)
		panic(err)
	}

}
