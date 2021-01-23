package service

import (
  "<%= appName %>/dto"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
)

type <%= entityCap %> interface {
  Save(<%= entityLower %>DTO dto.<%= entityCap %>DTO) (dto.<%= entityCap %>DTO, bool)
  FindOne(id uint) (dto.<%= entityCap %>DTO, bool)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
