package exemplos

/*
import (
  "log"
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
  // Querystring (o que vem na URL ?, &)
  name := c.Query("name")
  lastName := c.Query("lastname")

  c.JSON(http.StatusOK, gin.H{
    "message": "Hello " + name + " " + lastName,
  })
}

func helloHandler1(c *gin.Context) {
  // URL Param (o que vem depois do ":" -> :id)
  id := c.Param("id")

  _, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "message": "id is not a number",
    })
    log.Println(err)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Hello: " + id,
  })
}

func main() {
  gin.SetMode("debug")

  router := gin.Default()

  router.GET("/", rootHandler)

  group := router.Group("/api/v1")
  {
    group.GET("/hello-world/:id", helloHandler1)
  }

  router.Run()
}

*/
