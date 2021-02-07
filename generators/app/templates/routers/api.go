package routers

import (
  "<%= appName %>/config"
  "<%= appName %>/controller"

  "github.com/gin-gonic/gin"
)

// InitRouterApi InitRouterApi
func InitRouterApi(r *gin.Engine) {
  for _, controller := range config.Controllers {
    registerController(r, controller)
  }
}

func registerController(r *gin.Engine, baseController controller.Base) {
  for _, route := range baseController.GetRoutes() {
    r.Handle(route.Method, route.Path, route.Handler)
  }
}
