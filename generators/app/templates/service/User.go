package service

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/dto"
  "<%= appName %>/helpers/jwt"
  "<%= appName %>/repository"
  "<%= appName %>/service/mapper"

  "golang.org/x/crypto/bcrypt"
)

type User interface {
  Create(userDTO dto.UserDTO) (dto.UserDTO, bool)
  GetUserToken(userDTO dto.UserDTO) (string, error)
  FindOneLogin(username string, password string) (dto.UserDTO, bool)
  FindOneByUserID(userId string) (dto.UserDTO, bool)
  FindOneByUsername(username string) (dto.UserDTO, bool)
  GenerateToken(userID string, username string, roles []string) (string, error)
}

type user struct {
  repository repository.User
  mapper     mapper.User
  jwtManager jwt.JwtManager
}

func NewUser(repository repository.User, mapper mapper.User, jwtManager jwt.JwtManager) User {
  return &user{repository: repository, mapper: mapper, jwtManager: jwtManager}
}

func (s *user) Create(userDTO dto.UserDTO) (dto.UserDTO, bool) {
  pass, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
  if err != nil {
    return dto.UserDTO{}, false
  }
  userDTO.Password = string(pass)
  userDTO.Roles = append(userDTO.Roles, domain.ROLE_USER)

  user := s.mapper.ToEntity(userDTO)
  s.repository.Save(user)

  return s.mapper.ToDTO(user), true
}

func (s *user) GetUserToken(userDTO dto.UserDTO) (string, error) {
  tokenString, err := s.jwtManager.GenerateToken(userDTO.UserID, userDTO.Name, userDTO.GetRolesStr())
  if err != nil {
    return "", err
  }

  return tokenString, nil
}

func (s *user) FindOneLogin(username string, password string) (dto.UserDTO, bool) {
  user, err := s.repository.FindOneByName(username)
  if err != nil {
    return dto.UserDTO{}, false
  }

  errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
    return dto.UserDTO{}, false
  }

  return s.mapper.ToDTO(user), true
}

func (s *user) FindOneByUserID(userId string) (dto.UserDTO, bool) {
  user, err := s.repository.FineOneByUserId(userId)
  if err != nil {
    return dto.UserDTO{}, false
  }

  return s.mapper.ToDTO(user), true
}

func (s *user) FindOneByUsername(username string) (dto.UserDTO, bool) {
  user, err := s.repository.FindOneByName(username)
  if err != nil {
    return dto.UserDTO{}, false
  }

  return s.mapper.ToDTO(user), true
}

func (s *user) GenerateToken(userID string, username string, roles []string) (string, error) {
  return s.jwtManager.GenerateToken(userID, username, roles)
}