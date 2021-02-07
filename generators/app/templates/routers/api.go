package routers

import (
  "<%= appName %>/config"
  "<%= appName %>/controller"

  "github.com/gin-gonic/gin"
)

// InitRouterApi InitRouterApi
func InitRouterApi(r *gin.Engine) {
  // Api declare
  registerController(r, config.AuthController)
  registerController(r, config.FileController)
  // Api declare end : dont remove
}

func registerController(r *gin.Engine, baseController controller.BaseController) {
  for _, route := range baseController.GetRoutes() {
    r.Handle(route.Method, route.Path, route.Handler)
  }
}
