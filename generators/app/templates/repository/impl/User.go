package repository_impl

import (
  "<%= appName %>/models"
  "<%= appName %>/repository"

  "gorm.io/gorm"
)

type user struct {
  db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
  return &user{ db: db}
}

func (r *user) Save(user models.User) models.User {
  r.db.Create(&user)
  return user
}

func (r *user) FineOneByUserId(userId string) (models.User, error) {
  user := models.User{}
  result := r.db.Where("user_id = ?", userId).First(&user)
  if result.Error != nil {
    return models.User{}, result.Error
  }
  return user, nil
}

func (r *user) FindOneByName(name string) (models.User, error) {
  user := models.User{}
  result := r.db.Where("name = ?", name).First(&user)
  if result.Error != nil {
    return models.User{}, result.Error
  }
  return user, nil
}
