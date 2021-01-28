package middlewares

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/dto/response"
  "<%= appName %>/helpers/constants"
  "regexp"

  "github.com/gin-gonic/gin"
)

var accessibleRoles map[string][]string

func init() {
  accessibleRoles = make(map[string][]string)
  // Security declare
  accessibleRoles["/api/private/.*"] = []string{constants.ROLE.USER}
  // Security declare end : dont remove
}

// Security is Security middleware
func Security(c *gin.Context) {
  var roles []string
  found := false
  pathBytes := []byte(c.Request.URL.Path)
  for path, _roles := range accessibleRoles {
    if regexp.MustCompile(path).Match(pathBytes) {
      roles = _roles
      found = true
      break
    }
  }

  if !found {
    c.Next()
    return
  }

  iUserInfo, exists := c.Get("UserInfo")
  if !exists {
    response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
    c.Abort()
    return
  }

  userInfo := iUserInfo.(*domain.Token)
  if err := userInfo.Valid(); err != nil {
    response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
    c.Abort()
    return
  }

  for _, authority := range roles {
    if !userInfo.HasAuthority(authority) {
      response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
      c.Abort()
      return
    }
  }

  c.Next()
}
