package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result is the base payload
type Result struct {
	Status  string
	Message string
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, Result{"OK", "HELLOWORLD!"})
	})

	router.Run(":8080")
}
