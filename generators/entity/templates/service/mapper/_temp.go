package mapper

import (
  "<%= appName %>/dto"
  "<%= appName %>/models"
)

type <%= entityCap %> interface {
  ToDTO(entity models.<%= entityCap %>) dto.<%= entityCap %>DTO
  ToEntity(dto dto.<%= entityCap %>DTO) models.<%= entityCap %>
  ToDTOS(entityList []models.<%= entityCap %>) []dto.<%= entityCap %>DTO
  ToEntities(dtoList []dto.<%= entityCap %>DTO) []models.<%= entityCap %>
}

type <%= entityLower %> struct {}

func New<%= entityCap %>() <%= entityCap %> {
  return &<%= entityLower %>{}
}

func (m *<%= entityLower %>) ToDTO(entity models.<%= entityCap %>) dto.<%= entityCap %>DTO {
  return dto.<%= entityCap %>DTO{
    ID:          entity.model.ID,
    CreatedAt:   entity.model.CreatedAt,
    UpdatedAt:   entity.model.UpdatedAt,
    DeletedAt:   entity.model.DeletedAt,
  }
}

func (m *<%= entityLower %>) ToEntity(dto dto.<%= entityCap %>DTO) models.<%= entityCap %> {
  return models.<%= entityCap %>{
    ID:          dto.ID,
    CreatedAt:   dto.CreatedAt,
    UpdatedAt:   dto.UpdatedAt,
    DeletedAt:   dto.DeletedAt,
  }
}

func (m *<%= entityLower %>) ToDTOS(entityList []models.<%= entityCap %>) []dto.<%= entityCap %>DTO {
  dtos := make([]dto.<%= entityCap %>DTO, len(entityList))

  for i, entity := range entityList {
    dtos[i] = m.ToDTO(entity)
  }

  return dtos
}

func (m *<%= entityLower %>) ToEntities(dtoList []dto.<%= entityCap %>DTO) []models.<%= entityCap %> {
  entities := make([]models.<%= entityCap %>, len(dtoList))

  for i, dto := range dtoList {
    entities[i] = m.ToEntity(dto)
  }

  return entities
}
