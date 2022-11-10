package main

import (
	"JET_DEV_TASK/JetDev/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// installing newe gin router
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		//Loading Env variable
		infrastructure.LoadEnv()
		//New DB connection
		infrastructure.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "hello worRRkd!!"})
	})
	router.Run(":8000")
}
