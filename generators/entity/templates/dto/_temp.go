package dto

import (
  "time"
)

// <%= entityCap %>DTO godoc
type <%= entityCap %>DTO struct {
  ID uint

  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}
