package Jwt

import (
  "github.com/dgrijalva/jwt-go"
)

type Token struct {
  UserID string   `json:"userId"`
  Name   string   `json:"name"`
  Roles  []string `json:"roles"`
  *jwt.StandardClaims
}

func (t *Token) GetUserID() string {
  return t.UserID
}

func (t *Token) GetUserName() string {
  return t.Name
}

func (t *Token) GetRoles() []string {
  return t.Roles
}

func (t *Token) HasAuthority(authority string) bool {
  for _, s := range t.Roles {
    if s == authority {
      return true
    }
  }
  return false
}
