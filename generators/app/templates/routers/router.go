package routers

import (
  _ "<%= appName %>/docs"
  "<%= appName %>/middlewares"
  "<%= appName %>/controller"
  "<%= appName %>/pkg/service/Jwt"

  // "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
  swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter(jwtManager Jwt.Manager, controllers []controller.Base) *gin.Engine {
  r := gin.New()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  r.Use(middlewares.JWT(jwtManager))
  r.Use(middlewares.Security)

  // r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

  r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  InitRouterApi(r, controllers)

  return r
}
