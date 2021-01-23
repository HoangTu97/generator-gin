package repository_proxy

import (
  "<%= appName %>/models"
  "<%= appName %>/repository"
)

type user struct {
  repository repository.User
}

func NewUser(repository repository.User) repository.User {
  return &user{repository: repository}
}

func (r *user) Save(user models.User) models.User {
  return r.repository.Save(user)
}

func (r *user) FineOneByUserId(userId string) (models.User, error) {
  return r.repository.FineOneByUserId(userId)
}

func (r *user) FindOneByName(name string) (models.User, error) {
  return r.repository.FindOneByName(name)
}
