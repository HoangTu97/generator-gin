package repository_impl

import (
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/models"
  "<%= appName %>/repository"

  "gorm.io/gorm"
)

type <%= entityLower %> struct {
  db *gorm.DB
}

func New<%= entityCap %>(db *gorm.DB) repository.<%= entityCap %> {
  return &<%= entityLower %>{db: db}
}

func (r *<%= entityLower %>) Save(<%= entityLower %> models.<%= entityCap %>) (models.<%= entityCap %>, error) {
  result := r.db.Save(&<%= entityLower %>)
  if result.Error != nil {
    return <%= entityLower %>, result.Error
  }
  return <%= entityLower %>, nil
}

func (r *<%= entityLower %>) FindOne(id uint) (models.<%= entityCap %>, error) {
  var <%= entityLower %> models.<%= entityCap %>

  result := r.db.First(&<%= entityLower %>, id)
  if result.Error != nil {
    return models.<%= entityCap %>{}, result.Error
  }

  return <%= entityLower %>, nil
}

func (r *<%= entityLower %>) FindPage(pageable pagination.Pageable) page.Page {
  var <%= entityLower %>s []models.<%= entityCap %>

  paginator := pagination.Paging(&pagination.Param{
    DB:      r.db.Joins("User").Joins("Recipe"),
    Page:    pageable.GetPageNumber(),
    Limit:   pageable.GetPageSize(),
    ShowSQL: true,
  }, &<%= entityLower %>s)

  return page.From(r.toInterfacesFrom<%= entityCap %>(<%= entityLower %>s), paginator.TotalRecord)
}

func (r *<%= entityLower %>) toInterfacesFrom<%= entityCap %>(<%= entityLower %>s []models.<%= entityCap %>) []interface{} {
  content := make([]interface{}, len(<%= entityLower %>s))
  for i, v := range <%= entityLower %>s {
    content[i] = v
  }
  return content
}

func (r *<%= entityLower %>) Delete(id uint) {
  r.db.Delete(&models.<%= entityCap %>{}, id)
}
