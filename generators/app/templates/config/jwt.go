package config

import (
  "<%= appName %>/helpers/jwt"
  "<%= appName %>/helpers/setting"
)

func SetupJWT(appSetting setting.App) jwt.JwtManager {
  return jwt.NewJwtManager(appSetting)
}