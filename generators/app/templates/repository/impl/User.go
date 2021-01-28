package repository_impl

import (
  "<%= appName %>/models"
  "<%= appName %>/repository"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"

  "gorm.io/gorm"
)

type user struct {
  db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
  return &user{ db: db}
}

func (r *user) Save(user models.User) (models.User, error) {
  result := r.db.Save(&user)
  if result.Error != nil {
    return user, result.Error
  }
  return user, nil
}

func (r *user) FindOne(id uint) (models.User, error) {
  var user models.User

  result := r.db.First(&user, id)
  if result.Error != nil {
    return models.User{}, result.Error
  }

  return user, nil
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

func (r *user) FindPage(pageable pagination.Pageable) page.Page {
  var users []models.User

  paginator := pagination.Paging(&pagination.Param{
    DB:      r.db.Joins("User").Joins("Recipe"),
    Page:    pageable.GetPageNumber(),
    Limit:   pageable.GetPageSize(),
    ShowSQL: true,
  }, &users)

  return page.From(r.toInterfacesFromUser(users), paginator.TotalRecord)
}

func (r *user) Delete(id uint) {
  r.db.Delete(&models.User{}, id)
}



func (r *user) toInterfacesFromUser(users []models.User) []interface{} {
  content := make([]interface{}, len(users))
  for i, v := range users {
    content[i] = v
  }
  return content
}
