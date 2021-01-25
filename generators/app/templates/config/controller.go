package config

import (
  "<%= appName %>/controller"
  "<%= appName %>/helpers/jwt"
  "<%= appName %>/pkg/cache"
  "<%= appName %>/repository/impl"
  "<%= appName %>/repository/proxy"
  "<%= appName %>/service/impl"
  "<%= appName %>/service/proxy"
  "<%= appName %>/service/mapper/impl"

  "gorm.io/gorm"
)

var (
  // Controllers globale declare
  UserController controller.User
  FileController controller.File
  // Controllers globale declare end : dont remove
)

func SetupController(
  db *gorm.DB, 
  jwtManager jwt.JwtManager,
  cache cache.Cache,
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
  userService := service_impl.NewUser(userRepoProxy, userMapper, jwtManager)
  fileService := service_impl.NewFile()
  // Services declare end : dont remove

  // Proxy Services declare
  userServiceProxy := service_proxy.NewUser(userService, cache)
  // Proxy Services declare end : dont remove

  // Controllers declare
  UserController = controller.NewUser(userServiceProxy)
  FileController = controller.NewFile(fileService)
  // Controllers declare end : dont remove
}
