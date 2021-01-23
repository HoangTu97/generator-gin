package middlewares

import (
  "<%= appName %>/helpers/jwt"
  "strings"

  "github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT(jwtManager jwt.JwtManager) gin.HandlerFunc {
  return func(c *gin.Context) {
    token := c.Request.Header.Get("Authorization")
    if len(token) == 0 {
      c.Next()
      return
    }
    token = strings.Fields(token)[1]

    parsedToken, err := jwtManager.ParseToken(token)
    if err != nil {
      c.Next()
      return
    }

    c.Set("UserInfo", parsedToken)

    c.Next()
  }
}
