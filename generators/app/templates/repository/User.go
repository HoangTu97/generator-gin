package repository

import (
  "<%= appName %>/models"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
)

type User interface {
  Save(user models.User) (models.User, error)
  FindOne(id uint) (models.User, error)
  FineOneByUserId(userId string) (models.User, error)
  FindOneByName(name string) (models.User, error)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
