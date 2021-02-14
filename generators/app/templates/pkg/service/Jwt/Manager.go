package Jwt

import (
  "<%= appName %>/pkg/service/Hash"

  "time"

  "github.com/dgrijalva/jwt-go"
  "github.com/spf13/viper"
)

type Manager interface {
  GenerateToken(userID string, username string, roles []string) (string, error)
  ParseToken(token string) (*Token, error)
}

type manager struct {
  jwtSecret   []byte
  hashService Hash.Service
}

func NewManager() Manager {
  return &manager{
    jwtSecret:   []byte(viper.GetString("app.jwtSecretKey")),
    hashService: Hash.NewService("md5"),
  }
}

// GenerateToken generate tokens used for auth
func (m *manager) GenerateToken(userID string, username string, roles []string) (string, error) {
  nowTime := time.Now()
  expireTime := nowTime.Add(3 * time.Hour)

  claims := Token{
    UserID: userID,
    Name:   m.hashService.Make(username),
    Roles:  roles,
    StandardClaims: &jwt.StandardClaims{
      ExpiresAt: expireTime.Unix(),
      Issuer:    "gin-blog",
    },
  }

  tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  token, err := tokenClaims.SignedString(m.jwtSecret)

  return token, err
}

func (m *manager) ParseToken(token string) (*Token, error) {
  tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(token *jwt.Token) (interface{}, error) {
    return m.jwtSecret, nil
  })
  if err != nil {
    return nil, err
  }

  if tokenClaims != nil {
    if claims, ok := tokenClaims.Claims.(*Token); ok && tokenClaims.Valid {
      return claims, nil
    }
  }

  return nil, err
}
