package jwt

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/pkg/util"
  "time"

  "github.com/dgrijalva/jwt-go"
)

type JwtManager interface {
  GenerateToken(userID string, username string, roles []string) (string, error)
  ParseToken(token string) (*domain.Token, error)
}

type jwtManager struct {
  jwtSecret []byte
}

// Setup Initialize the util
func NewJwtManager(jwtSecretKey string) JwtManager  {
  return &jwtManager{
    jwtSecret: []byte(jwtSecretKey),
  }
}

// GenerateToken generate tokens used for auth
func (m *jwtManager) GenerateToken(userID string, username string, roles []string) (string, error) {
  nowTime := time.Now()
  expireTime := nowTime.Add(3 * time.Hour)

  claims := domain.Token{
    UserID: userID,
    Name:   util.EncodeMD5(username),
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

// ParseToken parsing token
func (m *jwtManager) ParseToken(token string) (*domain.Token, error) {
  tokenClaims, err := jwt.ParseWithClaims(token, &domain.Token{}, func(token *jwt.Token) (interface{}, error) {
    return m.jwtSecret, nil
  })
  if err != nil {
    return nil, err
  }

  if tokenClaims != nil {
    if claims, ok := tokenClaims.Claims.(*domain.Token); ok && tokenClaims.Valid {
      return claims, nil
    }
  }

  return nil, err
}
