package service_impl

import (
  "<%= appName %>/dto"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/repository"
  "<%= appName %>/service"
  "<%= appName %>/service/mapper"
)

type <%= entityLower %> struct {
  repository repository.<%= entityCap %>
  mapper mapper.<%= entityCap %>
}

func New<%= entityCap %>(repository repository.<%= entityCap %>, mapper mapper.<%= entityCap %>) service.<%= entityCap %> {
  return &<%= entityLower %>{repository: repository, mapper: mapper}
}

func (s *<%= entityLower %>) Save(<%= entityLower %>DTO dto.<%= entityCap %>DTO) (dto.<%= entityCap %>DTO, bool) {
  <%= entityLower %> := s.mapper.ToEntity(<%= entityLower %>DTO)
  var err error
  <%= entityLower %>, err = s.repository.Save(<%= entityLower %>)
  if err != nil {
    return <%= entityLower %>DTO, false
  }
  return s.mapper.ToDTO(<%= entityLower %>), true
}

func (s *<%= entityLower %>) FindOne(id uint) (dto.<%= entityCap %>DTO, bool) {
  <%= entityLower %>, err := s.repository.FindOne(id)
  if err != nil {
    return dto.<%= entityCap %>DTO{}, false
  }
  return s.mapper.ToDTO(<%= entityLower %>), true
}

func (s *<%= entityLower %>) FindPage(pageable pagination.Pageable) page.Page {
  return s.repository.FindPage(pageable)
}

func (s *<%= entityLower %>) Delete(id uint) {
  s.repository.Delete(id)
}
