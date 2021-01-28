package service_impl

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/dto"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/repository"
  "<%= appName %>/service"
  "<%= appName %>/service/mapper"

  "golang.org/x/crypto/bcrypt"
)

type user struct {
  repository repository.User
  mapper     mapper.User
}

func NewUser(repository repository.User, mapper mapper.User) service.User {
  return &user{repository: repository, mapper: mapper}
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

func (s *user) Save(userDTO dto.UserDTO) (dto.UserDTO, bool) {
  user := s.mapper.ToEntity(userDTO)
  var err error
  user, err = s.repository.Save(user)
  if err != nil {
    return userDTO, false
  }
  return s.mapper.ToDTO(user), true
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

func (s *user) FindPage(pageable pagination.Pageable) page.Page {
  return s.repository.FindPage(pageable)
}

func (s *user) Delete(id uint) {
  s.repository.Delete(id)
}
