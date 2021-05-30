package mapper_impl

import (
  "<%= appName %>/pkg/domain"
  "<%= appName %>/dto"
  "<%= appName %>/models"
  "<%= appName %>/pkg/converter"
  "<%= appName %>/service/mapper"

  uuid "github.com/satori/go.uuid"
  "gorm.io/gorm"
)

type user struct {}

func NewUser() mapper.User {
  return &user{}
}

func (m *user) ToDTO(entity models.User) dto.UserDTO {
  roleStrings, _ := converter.ArrStr(entity.Roles)
  roles := make([]domain.UserRole, len(roleStrings))
  for i, r := range roleStrings {
    ro, _ := domain.ParseUserRole(r)
    roles[i] = ro
  }

  return dto.UserDTO{
    ID:         entity.Model.ID,
    CreatedAt:  entity.Model.CreatedAt,
    UpdatedAt:  entity.Model.UpdatedAt,
    DeletedAt:  entity.Model.DeletedAt,
    UserID:     entity.UserID.String(),
    Name:       entity.Name,
    Password:   entity.Password,
    Roles:      roles,
    Address:    entity.Address,
    Age:        entity.Age,
    Gender:     entity.Gender,
    Occupation: entity.Occupation,
    Long:       entity.Long,
    Lat:        entity.Lat,
    ZipCode:    entity.ZipCode,
  }
}

func (m *user) ToEntity(dto dto.UserDTO) models.User {
  var id uuid.UUID
  if len(dto.UserID) > 0 {
    id = uuid.Must(uuid.FromString(dto.UserID))
  }

  roles := make([]string, len(dto.Roles))
  for i, r := range dto.Roles {
    roles[i] = r.String()
  }

  return models.User{
    Model: gorm.Model{
      ID:         dto.ID,
      CreatedAt:  dto.CreatedAt,
      UpdatedAt:  dto.UpdatedAt,
      DeletedAt:  dto.DeletedAt,
    },
    UserID:     id,
    Name:       dto.Name,
    Password:   dto.Password,
    Roles:      converter.ToStr(roles),
    Address:    dto.Address,
    Age:        dto.Age,
    Gender:     dto.Gender,
    Occupation: dto.Occupation,
    Long:       dto.Long,
    Lat:        dto.Lat,
    ZipCode:    dto.ZipCode,
  }
}

func (m *user) ToDTOS(entityList []models.User) []dto.UserDTO {
  dtos := make([]dto.UserDTO, len(entityList))

  for i, entity := range entityList {
    dtos[i] = m.ToDTO(entity)
  }

  return dtos
}

func (m *user) ToEntities(dtoList []dto.UserDTO) []models.User {
  entities := make([]models.User, len(dtoList))

  for i, dto := range dtoList {
    entities[i] = m.ToEntity(dto)
  }

  return entities
}
