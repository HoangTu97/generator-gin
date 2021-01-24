package service_proxy

import (
  "<%= appName %>/dto"
  "<%= appName %>/helpers/page"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/service"
)

// <%= entityLower %> use for async,...
type <%= entityLower %> struct {
  service service.<%= entityCap %>
}

func New<%= entityCap %>(service service.<%= entityCap %>) service.<%= entityCap %> {
  return &<%= entityLower %>{service: service}
}

func (p *<%= entityLower %>) Save(<%= entityLower %>DTO dto.<%= entityCap %>DTO) (dto.<%= entityCap %>DTO, bool) {
  return p.service.Save(<%= entityLower %>DTO)
}

func (p *<%= entityLower %>) FindOne(id uint) (dto.<%= entityCap %>DTO, bool) {
  return p.service.FindOne(id)
}

func (p *<%= entityLower %>) FindPage(pageable pagination.Pageable) page.Page {
  return p.service.FindPage(pageable)
}

func (p *<%= entityLower %>) Delete(id uint) {
  p.service.Delete(id)
}
