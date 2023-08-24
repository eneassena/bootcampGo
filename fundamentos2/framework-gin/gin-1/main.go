package main

import "github.com/gin-gonic/gin"

//	trabalhando com o framework gin
func main() {
	router := gin.Default()

	router.GET("/hello", helloHandlerGin)

	router.Run()
}

func helloHandlerGin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World",
	})
}
