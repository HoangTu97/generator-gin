package database

import (
  "fmt"
  "log"

  "gorm.io/driver/postgres"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

func NewDB(config Config) (*gorm.DB, func()) {
  var dialector gorm.Dialector

  switch config.Type {
  case "sqlite3":
    dialector = sqlite.Open(config.Name)
  case "postgres":
    dialector = postgres.New(postgres.Config{
      DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s ",
        config.Host,
        config.User,
        config.Password,
        config.Name),
    })
  default:
    log.Fatalln("NewDB -> config invalid -> ", config)
  }

  db, err := gorm.Open(dialector, &gorm.Config{})
  if err != nil {
    log.Fatalf("models.Setup err: %v", err)
  }

  teardown := func() {
    sqlDB, _ := db.DB()
    sqlDB.Close()
  }
  return db, teardown
}
