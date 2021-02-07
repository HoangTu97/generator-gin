package controller

import (
  "github.com/gin-gonic/gin"
)

type BaseController interface {
  GetRoutes() []RouteController
}

type RouteController struct {
  Method string
  Path string
  Handler gin.HandlerFunc
}