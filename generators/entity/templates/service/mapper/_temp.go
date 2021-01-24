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
