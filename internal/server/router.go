package server
 
import (
  "net/http"
  
  "github.com/gin-gonic/gin"
)
 
func setRouter() *gin.Engine {
  // Creates default gin router with Logger and Recovery middleware already attached
  router := gin.Default()
 
  // Create API route group
  api := router.Group("/api")
  {
    api.GET("/hello", func(ctx *gin.Context) {
      ctx.JSON(200, gin.H{"msg": "world"})
    })
    api.POST("/signup", signUp)
    api.POST("/signin", signIn)
    api.POST("/upload", processFile)
  }
 
  router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
 
  return router
}