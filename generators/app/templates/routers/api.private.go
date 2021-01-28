package routers

import (
  "<%= appName %>/config"

  "github.com/gin-gonic/gin"
)

func registerPrivateApi(apiRoutes *gin.RouterGroup) {
  privateRoutes := apiRoutes.Group("/private")
  // Api declare
  {
    privateUserRoutes := privateRoutes.Group("/auth")
    privateUserRoutes.GET("/userinfo", config.AuthController.UserInfo)
  }
  // Api declare end : dont remove
}