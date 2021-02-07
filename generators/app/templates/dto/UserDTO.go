package dto

import (
  "<%= appName %>/pkg/domain"
  "time"
  "gorm.io/gorm"
)

// UserDTO godoc
type UserDTO struct {
  ID       uint              `json:"id"`
  UserID   string            `json:"userId"`
  Name     string            `json:"name"`
  Password string            `json:"password"`
  Roles    []domain.UserRole `json:"roles"`
  Address  string            `json:"address"`

  // features
  Age        uint8   `json:"age"`
  Gender     string  `json:"gender"`
  Occupation string  `json:"occupation"`
  Long       float32 `json:"long"`
  Lat        float32 `json:"lat"`
  ZipCode    uint16  `json:"zipCode"`

  CreatedAt time.Time  `json:"createdAt"`
  UpdatedAt time.Time  `json:"updatedAt"`
  DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (dto *UserDTO) GetRolesInterface() []interface{} {
  arr := make([]interface{}, len(dto.Roles))

  for i, role := range dto.Roles {
    arr[i] = role
  }

  return arr
}

func (dto *UserDTO) GetRolesStr() []string {
  arr := make([]string, len(dto.Roles))

  for i, role := range dto.Roles {
    arr[i] = role.String()
  }

  return arr
}
