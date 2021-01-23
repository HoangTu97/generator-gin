package security

import (
  "<%= appName %>/pkg/domain"

  "github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) *domain.Token {
  return c.MustGet("UserInfo").(*domain.Token)
}

func GetUserName(c *gin.Context) string {
  return GetUserInfo(c).GetUserName()
}

func GetUserID(c *gin.Context) string {
  return GetUserInfo(c).GetUserID()
}
