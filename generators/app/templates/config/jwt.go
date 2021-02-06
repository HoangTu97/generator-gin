package config

import (
  "<%= appName %>/helpers/jwt"

  "github.com/spf13/viper"
)

func SetupJWT() jwt.JwtManager {
  return jwt.NewJwtManager(viper.GetString("app.jwtSecretKey"))
}