package routers

import (
  "<%= appName %>/config"

  "github.com/gin-gonic/gin"
)

func registerPublicApi(apiRoutes *gin.RouterGroup) {
  publicRoutes := apiRoutes.Group("/public")
  // Api declare
  {
    publicUserRoutes := publicRoutes.Group("/user")
    publicUserRoutes.POST("/register", config.UserController.Register)
    publicUserRoutes.POST("/login", config.UserController.Login)
  }
  {
    publicFileRoutes := publicRoutes.Group("/file")
    publicFileRoutes.POST("/upload", config.FileController.Upload)
    publicFileRoutes.GET("/:id", config.FileController.FileDisplay)
    publicFileRoutes.GET("/:id/download", config.FileController.Download)
  }
  // Api declare end : dont remove
}
