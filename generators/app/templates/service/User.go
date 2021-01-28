package service

import (
  "<%= appName %>/dto"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
)

type User interface {
  Create(userDTO dto.UserDTO) (dto.UserDTO, bool)
  Save(userDTO dto.UserDTO) (dto.UserDTO, bool)
  FindOneLogin(username string, password string) (dto.UserDTO, bool)
  FindOneByUserID(userId string) (dto.UserDTO, bool)
  FindOneByUsername(username string) (dto.UserDTO, bool)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
