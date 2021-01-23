package repository

import (
  "<%= appName %>/models"
)

type User interface {
  Save(user models.User) models.User
  FineOneByUserId(userId string) (models.User, error)
  FindOneByName(name string) (models.User, error)
}
