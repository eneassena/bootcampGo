package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func lendoRequest(c *gin.Context) {
	//O corpo, cabeçalho e método estão contidos no contexto gin.
	body := c.Request.Body
	header := c.Request.Header
	method := c.Request.Method

	fmt.Println("Eu recebi algo!")
	fmt.Printf("\tMétodo: %s\n", method)
	fmt.Printf("\tConteúdo do cabeçalho:\n")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}

	fmt.Printf("\tO body é um io.ReadCloser:(%s), e para trabalhar com ele teremos que ler depois\n", body)
	fmt.Println("¡Yay!")
	c.String(200, "Eu recebi!") //Respondemos ao cliente com 200 OK e uma mensagem.
}

func lendoParamentroQuery(c *gin.Context) {
	id := c.Param("id")

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "id is not a number",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s", id),
	})
}

func main() {
	gin.SetMode("debug")

	router := gin.Default()

	router.GET("/", lendoRequest)

	// define grupo de rotas
	group := router.Group("/api/v1")
	{
		group.GET("/hello-world/:id", lendoParamentroQuery)
	}

	router.Run()
}
