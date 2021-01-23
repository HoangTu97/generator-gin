package routers

import (
  "github.com/gin-gonic/gin"
)

// InitRouterApi InitRouterApi
func InitRouterApi(r *gin.Engine) {
  apiRoutes := r.Group("/api")
  registerPublicApi(apiRoutes)
  registerPrivateApi(apiRoutes)
}
