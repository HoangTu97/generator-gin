package config

import (
  "<%= appName %>/controller"
  "<%= appName %>/helpers/jwt"
  "<%= appName %>/helpers/service/Auth"
  "<%= appName %>/helpers/service/File"
  "<%= appName %>/helpers/service/Cache"
  // "<%= appName %>/helpers/service/Mail"
  "<%= appName %>/helpers/service/Hash"
  "<%= appName %>/helpers/service/Schedule"
  "<%= appName %>/repository/impl"
  "<%= appName %>/repository/proxy"
  "<%= appName %>/service/impl"
  "<%= appName %>/service/proxy"
  "<%= appName %>/service/mapper/impl"

  "gorm.io/gorm"
)

var (
  // Controllers globale declare
  FileController controller.File
  UserController controller.User
  AuthController controller.Auth
  // Controllers globale declare end : dont remove
)

func SetupController(
  db *gorm.DB, 
  jwtManager jwt.JwtManager,
  cacheManager Cache.Manager,
) {
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
  cacheService := cacheManager.Driver("Memcached")
  fileService := File.NewService()
  // mailService := Mail.NewService()
  hashService := Hash.NewService("")
  scheduleService := Schedule.NewService()
  authService := Auth.NewService(jwtManager)
  userService := service_impl.NewUser(userRepoProxy, userMapper, hashService)
  // Services declare end : dont remove

  // Proxy Services declare
  userServiceProxy := service_proxy.NewUser(userService, cacheService)
  // Proxy Services declare end : dont remove

  // Controllers declare
  FileController = controller.NewFile(fileService)
  AuthController = controller.NewAuth(authService, userServiceProxy)
  UserController = controller.NewUser(userServiceProxy)
  // Controllers declare end : dont remove
}
