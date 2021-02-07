package config

import (
  "log"

  "gorm.io/gorm"

  "<%= appName %>/models"
)

// Setup initializes the database instance
func SetupDB(db *gorm.DB) *gorm.DB {
  sqlDB, errSqlDB := db.DB()
  if errSqlDB != nil {
    log.Fatalf("models.Setup db.DB() err: %v", errSqlDB)
  }

  sqlDB.SetMaxIdleConns(10)
  sqlDB.SetMaxOpenConns(100)
  sqlDB.Stats()

  migrateDB(db)

  return db
}

func migrateDB(db *gorm.DB) {
  _ = db.AutoMigrate(
    // Models declare
    &models.User{},
    // Models declare end : dont remove
  )
}
