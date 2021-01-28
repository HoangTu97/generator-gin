package controller

import (
  // "<%= appName %>/dto"
  // "<%= appName %>/dto/request"
  // UserRequest "<%= appName %>/dto/request/user"
  // "<%= appName %>/dto/response"
  // UserResponse "<%= appName %>/dto/response/user"
  // "<%= appName %>/helpers/constants"
  "<%= appName %>/service"

  // "github.com/gin-gonic/gin"
)

type User interface {
  // Register(c *gin.Context)
  // Login(c *gin.Context)
  // UserInfo(c *gin.Context)
}

type user struct {
  service service.User
}

func NewUser(service service.User) User {
  return &user{service: service}
}
