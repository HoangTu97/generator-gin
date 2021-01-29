package service_impl

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/dto"
  "<%= appName %>/helpers/jwt"
  "<%= appName %>/helpers/service"

  "github.com/gin-gonic/gin"
)

type auth struct {
  jwtManager jwt.JwtManager
}

func NewAuth(jwtManager jwt.JwtManager) service.Auth {
  return &auth{jwtManager: jwtManager}
}

func (s *auth) GenerateToken(userDTO dto.UserDTO) (string, error) {
  return s.jwtManager.GenerateToken(userDTO.UserID, userDTO.Name, userDTO.GetRolesStr())
}

func (s *auth) GetUserInfo(c *gin.Context) *domain.Token {
  return c.MustGet("UserInfo").(*domain.Token)
}

func (s *auth) GetUserName(c *gin.Context) string {
  return s.GetUserInfo(c).GetUserName()
}

func (s *auth) GetUserID(c *gin.Context) string {
  return s.GetUserInfo(c).GetUserID()
}

func (s *auth) Check(c * gin.Context) bool {
  _, exists := c.Get("UserInfo")
  return exists
}
