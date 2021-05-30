package repository_proxy

import (
  "<%= appName %>/models"
  "<%= appName %>/repository"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
)

type user struct {
  repository repository.User
}

func NewUser(repository repository.User) repository.User {
  return &user{repository: repository}
}

func (r *user) Save(user models.User) (models.User, error) {
  return r.repository.Save(user)
}

func (r *user) FindOne(id uint) (models.User, error) {
  return r.repository.FindOne(id)
}

func (r *user) FineOneByUserId(userId string) (models.User, error) {
  return r.repository.FineOneByUserId(userId)
}

func (r *user) FindOneByName(name string) (models.User, error) {
  return r.repository.FindOneByName(name)
}

func (r *user) FindPage(pageable pagination.Pageable) page.Page {
  return r.repository.FindPage(pageable)
}

func (r *user) Delete(id uint) {
  r.repository.Delete(id)
}
