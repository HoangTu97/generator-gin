package repository

import (
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/models"
)

type <%= entityCap %> interface {
  Save(<%= entityLower %> models.<%= entityCap %>) (models.<%= entityCap %>, error)
  FindOne(id uint) (models.<%= entityCap %>, error)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
