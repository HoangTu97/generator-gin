package routers

import (
  "<%= appName %>/config"

  "github.com/gin-gonic/gin"
)

func registerPublicApi(apiRoutes *gin.RouterGroup) {
  publicRoutes := apiRoutes.Group("/public")
  // Api declare
  {
    publicFileRoutes := publicRoutes.Group("/file")
    publicFileRoutes.POST("/upload", config.FileController.Upload)
    publicFileRoutes.GET("/:id", config.FileController.FileDisplay)
    publicFileRoutes.GET("/:id/download", config.FileController.Download)
  }
  {
    publicUserRoutes := publicRoutes.Group("/auth")
    publicUserRoutes.POST("/register", config.AuthController.Register)
    publicUserRoutes.POST("/login", config.AuthController.Login)
  }
  // Api declare end : dont remove
}
