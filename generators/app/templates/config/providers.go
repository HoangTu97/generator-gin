package config

import (
  "<%= appName %>/controller"
  "<%= appName %>/pkg/service/Auth"
  "<%= appName %>/pkg/service/Database"
  "<%= appName %>/pkg/service/File"
  "<%= appName %>/pkg/service/Cache"
  "<%= appName %>/pkg/service/Mail"
  "<%= appName %>/pkg/service/Jwt"
  "<%= appName %>/pkg/service/Hash"
  // "<%= appName %>/pkg/service/Schedule"
  "<%= appName %>/repository/impl"
  "<%= appName %>/repository/proxy"
  "<%= appName %>/service/impl"
  "<%= appName %>/service/proxy"
  "<%= appName %>/service/mapper/impl"
)

func Providers(
  dbManager Database.Manager, 
  jwtManager Jwt.Manager,
  cacheManager Cache.Manager,
  mailManager Mail.Manager,
) []controller.Base {
  db := dbManager.Connection("")

  // Mappers declare
  userMapper := mapper_impl.NewUser()
  // Mappers declare end : dont remove

  // Repositories declare
  userRepo := repository_impl.NewUser(db)
  // Repositories declare end : dont remove

  // Proxy Repositories declare
  userRepoProxy := repository_proxy.NewUser(userRepo)
  // Proxy Repositories declare end : dont remove

  // Services declare
  cacheService := cacheManager.Driver("")
  fileService := File.NewService()
  // mailService := mailManager.Mailer("smtp")
  hashService := Hash.NewService("")
  // scheduleService := Schedule.NewService()
  authService := Auth.NewService(jwtManager)
  userService := service_impl.NewUser(userRepoProxy, userMapper, hashService)
  // Services declare end : dont remove

  // Proxy Services declare
  userServiceProxy := service_proxy.NewUser(userService, cacheService)
  // Proxy Services declare end : dont remove

  // Controllers declare
  fileController := controller.NewFile(fileService)
  authController := controller.NewAuth(authService, userServiceProxy)
  userController := controller.NewUser(userServiceProxy)
  // Controllers declare end : dont remove

  return []controller.Base{
    // Register controller declare
    fileController,
    authController,
    userController,
    // Register controller declare end : dont remove
  }
}
