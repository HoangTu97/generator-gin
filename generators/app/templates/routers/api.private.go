package routers

import (
  "<%= appName %>/config"

  "github.com/gin-gonic/gin"
)

func registerPrivateApi(apiRoutes *gin.RouterGroup) {
  privateRoutes := apiRoutes.Group("/private")
  // Api declare
  {
    privateUserRoutes := privateRoutes.Group("/user")
    privateUserRoutes.GET("/userinfo", config.UserController.UserInfo)
  }
  // Api declare end : dont remove
}