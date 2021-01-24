package mapper_impl

import (
  "<%= appName %>/dto"
  "<%= appName %>/models"
  "<%= appName %>/service/mapper"
  "gorm.io/gorm"
)

type <%= entityLower %> struct {}

func New<%= entityCap %>() mapper.<%= entityCap %> {
  return &<%= entityLower %>{}
}

func (m *<%= entityLower %>) ToDTO(entity models.<%= entityCap %>) dto.<%= entityCap %>DTO {
  return dto.<%= entityCap %>DTO{
    ID:          entity.Model.ID,
    CreatedAt:   entity.Model.CreatedAt,
    UpdatedAt:   entity.Model.UpdatedAt,
  }
}

func (m *<%= entityLower %>) ToEntity(dto dto.<%= entityCap %>DTO) models.<%= entityCap %> {
  return models.<%= entityCap %>{
    gorm.Model{
      ID:          dto.ID,
      CreatedAt:   dto.CreatedAt,
      UpdatedAt:   dto.UpdatedAt,
    },
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