package database

import (
  "fmt"
  "log"

  "gorm.io/driver/postgres"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "github.com/spf13/viper"
)

func NewDB() (*gorm.DB, func()) {
  var dialector gorm.Dialector

  switch viper.GetString("db.default") {
  case "sqlite3":
    dialector = sqlite.Open(viper.GetString("db.driver.sqlite3.path"))
  case "postgres":
    dialector = postgres.New(postgres.Config{
      DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s ",
        viper.GetString("db.driver.pgsql.host"),
        viper.GetString("db.driver.pgsql.username"),
        viper.GetString("db.driver.pgsql.password"),
        viper.GetString("db.driver.pgsql.database"),
      ),
    })
  default:
    log.Fatalln("NewDB -> config invalid -> ", viper.GetString("db.default"))
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
