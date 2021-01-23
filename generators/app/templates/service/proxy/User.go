package service_proxy

import (
  "<%= appName %>/dto"
  "<%= appName %>/helpers/constants"
  "<%= appName %>/pkg/cache"
  "<%= appName %>/service"
  "encoding/json"
)

type userProxy struct {
  service service.User
  cache   cache.Cache
}

func NewUser(service service.User, cache cache.Cache) service.User {
  return &userProxy{service: service, cache: cache}
}

func (s *userProxy) Create(userDTO dto.UserDTO) (dto.UserDTO, bool) {
  return s.service.Create(userDTO)
}

func (s *userProxy) GetUserToken(userDTO dto.UserDTO) (string, error) {
  var token string
  key := s.cache.GenKey(constants.CACHE.USER, "token", userDTO.UserID)
  if s.cache.Exists(key) {
    data, err := s.cache.Get(key)
    if err != nil {
      return "", err
    }
    err = json.Unmarshal(data, &token)
    if err != nil {
      return "", err
    }
    return token, nil
  }

  token, err := s.service.GetUserToken(userDTO)
  if err != nil {
    return "", err
  }

  _ = s.cache.Set(key, token, 3600)

  return token, nil
}

func (s *userProxy) FindOneLogin(username string, password string) (dto.UserDTO, bool) {
  var userDTO dto.UserDTO

  key := s.cache.GenKey(constants.CACHE.USER, username, password)
  if s.cache.Exists(key) {
    data, err := s.cache.Get(key)
    if err != nil {
      return dto.UserDTO{}, false
    }
    err = json.Unmarshal(data, &userDTO)
    if err != nil {
      return dto.UserDTO{}, false
    }
    return userDTO, true
  }

  userDTO, exist := s.service.FindOneLogin(username, password)
  if !exist {
    return dto.UserDTO{}, false
  }

  _ = s.cache.Set(key, userDTO, 3600)

  return userDTO, true
}

func (s *userProxy) FindOneByUserID(userId string) (dto.UserDTO, bool) {
  return s.service.FindOneByUserID(userId)
}

func (s *userProxy) FindOneByUsername(username string) (dto.UserDTO, bool) {
  var userDTO dto.UserDTO

  key := s.cache.GenKey(constants.CACHE.USER, "name", username)
  if s.cache.Exists(key) {
    data, err := s.cache.Get(key)
    if err != nil {
      return dto.UserDTO{}, false
    }
    err = json.Unmarshal(data, &userDTO)
    if err != nil {
      return dto.UserDTO{}, false
    }
    return userDTO, true
  }

  userDTO, exist := s.service.FindOneByUsername(username)
  if !exist {
    return dto.UserDTO{}, false
  }

  _ = s.cache.Set(key, userDTO, 3600)

  return userDTO, true
}

func (s *userProxy) GenerateToken(userID string, username string, roles []string) (string, error) {
  return s.service.GenerateToken(userID, username, roles)
}
