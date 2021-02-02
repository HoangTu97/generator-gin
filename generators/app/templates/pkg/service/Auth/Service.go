package Auth

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/dto"
  "<%= appName %>/helpers/jwt"

  "github.com/gin-gonic/gin"
)

type Service interface {
  GenerateToken(userDTO dto.UserDTO) (string, error)
  GetUserInfo(c *gin.Context) *domain.Token
  GetUserName(c *gin.Context) string
  GetUserID(c *gin.Context) string
  Check(c * gin.Context) bool
}

type service struct {
  jwtManager jwt.JwtManager
}

func NewService(jwtManager jwt.JwtManager) Service {
  return &service{jwtManager: jwtManager}
}

func (s *service) GenerateToken(userDTO dto.UserDTO) (string, error) {
  return s.jwtManager.GenerateToken(userDTO.UserID, userDTO.Name, userDTO.GetRolesStr())
}

func (s *service) GetUserInfo(c *gin.Context) *domain.Token {
  return c.MustGet("UserInfo").(*domain.Token)
}

func (s *service) GetUserName(c *gin.Context) string {
  return s.GetUserInfo(c).GetUserName()
}

func (s *service) GetUserID(c *gin.Context) string {
  return s.GetUserInfo(c).GetUserID()
}

func (s *service) Check(c * gin.Context) bool {
  _, exists := c.Get("UserInfo")
  return exists
}
