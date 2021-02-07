package routers

import (
  "<%= appName %>/controller"

  "github.com/gin-gonic/gin"
)

// InitRouterApi InitRouterApi
func InitRouterApi(r *gin.Engine, controllers []controller.Base) {
  for _, controller := range controllers {
    registerController(r, controller)
  }
}

func registerController(r *gin.Engine, baseController controller.Base) {
  for _, route := range baseController.GetRoutes() {
    r.Handle(route.Method, route.Path, route.Handler)
  }
}
