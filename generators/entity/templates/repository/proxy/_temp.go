package repository_proxy

import (
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/models"
  "<%= appName %>/repository"
)

type <%= entityLower %> struct {
  repository repository.<%= entityCap %>
}

func New<%= entityCap %>(repository repository.<%= entityCap %>) repository.<%= entityCap %> {
  return &<%= entityLower %>{repository: repository}
}

func (r *<%= entityLower %>) Save(<%= entityLower %> models.<%= entityCap %>) (models.<%= entityCap %>, error) {
  return r.repository.Save(<%= entityLower %>)
}

func (r *<%= entityLower %>) FindOne(id uint) (models.<%= entityCap %>, error) {
  return r.repository.FindOne(id)
}

func (r *<%= entityLower %>) FindPage(pageable pagination.Pageable) page.Page {
  return r.repository.FindPage(pageable)
}

func (r *<%= entityLower %>) Delete(id uint) {
  r.repository.Delete(id)
}
