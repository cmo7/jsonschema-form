package main

import (
  "github.com/swaggest/jsonschema-go"
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/gin-contrib/cors"
)

type TestUser struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Friends []int `json:"friends,omitempty"`
}

func main() {
  router := gin.Default()
  router.Use(cors.Default())
  router.GET("/get-schema", getSchema)

  err := router.Run(":8080")
  if err != nil {
    panic(err)
  }
}

func getSchema(c *gin.Context) {
  reflector := jsonschema.Reflector{}
  schema, err := reflector.Reflect(TestUser{})
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.IndentedJSON(http.StatusOK, schema)
}
