package config

import (
  "<%= appName %>/models"
)

func GetModelsNeedMigrate() []interface{} {
  return []interface{}{
    // Models declare
    &models.User{},
    // Models declare end : dont remove
  }
}
