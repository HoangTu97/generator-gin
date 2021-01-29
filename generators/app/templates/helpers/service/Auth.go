package service

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/dto"

  "github.com/gin-gonic/gin"
)

type Auth interface {
  GenerateToken(userDTO dto.UserDTO) (string, error)
  GetUserInfo(c *gin.Context) *domain.Token
  GetUserName(c *gin.Context) string
  GetUserID(c *gin.Context) string
  Check(c * gin.Context) bool
}
