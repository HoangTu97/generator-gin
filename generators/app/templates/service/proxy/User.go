package service_proxy

import (
  "<%= appName %>/dto"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/helpers/constants"
  "<%= appName %>/pkg/service/Cache"
  "<%= appName %>/service"
)

type user struct {
  service service.User
  cache   Cache.Service
}

func NewUser(service service.User, cache Cache.Service) service.User {
  return &user{service: service, cache: cache}
}

func (s *user) Create(userDTO dto.UserDTO) (dto.UserDTO, bool) {
  return s.service.Create(userDTO)
}

func (p *user) Save(userDTO dto.UserDTO) (dto.UserDTO, bool) {
  return p.service.Save(userDTO)
}

func (s *user) FindOneLogin(username string, password string) (dto.UserDTO, bool) {
  var userDTO dto.UserDTO

  key := s.cache.GenKey(constants.CACHE.USER, username, password)
  if s.cache.Has(key) {
    data := s.cache.Get(key)
    if data == nil {
      return dto.UserDTO{}, false
    }
    return data.(dto.UserDTO), true
  }

  userDTO, exist := s.service.FindOneLogin(username, password)
  if !exist {
    return dto.UserDTO{}, false
  }

  _ = s.cache.Set(key, userDTO, 3600)

  return userDTO, true
}

func (s *user) FindOneByUserID(userId string) (dto.UserDTO, bool) {
  return s.service.FindOneByUserID(userId)
}

func (s *user) FindOneByUsername(username string) (dto.UserDTO, bool) {
  var userDTO dto.UserDTO

  key := s.cache.GenKey(constants.CACHE.USER, "name", username)
  if s.cache.Has(key) {
    data := s.cache.Get(key)
    if data == nil {
      return dto.UserDTO{}, false
    }
    return data.(dto.UserDTO), true
  }

  userDTO, exist := s.service.FindOneByUsername(username)
  if !exist {
    return dto.UserDTO{}, false
  }

  _ = s.cache.Set(key, userDTO, 3600)

  return userDTO, true
}

func (p *user) FindPage(pageable pagination.Pageable) page.Page {
  return p.service.FindPage(pageable)
}

func (p *user) Delete(id uint) {
  p.service.Delete(id)
}
