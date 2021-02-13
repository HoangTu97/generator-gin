package Database

import (
  "fmt"
  "log"

  "github.com/spf13/viper"
  "gorm.io/driver/postgres"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Manager interface {
  Connection(name string) *gorm.DB
  Migrate(name string, entities []interface{})
  Shutdown()
}

type manager struct {
  connections map[string]*gorm.DB
}

func NewManager() Manager {
  return &manager{
    connections: make(map[string]*gorm.DB),
  }
}

func (m *manager) Connection(name string) *gorm.DB {
  if name == "" {
    name = m.getDefault()
  }
  m.connections[name] = m.get(name)
  return m.connections[name]
}

func (m *manager) get(name string) *gorm.DB {
  if m.connections[name] == nil {
    log.Printf("Connecting DB %s", name)

    db, _ := gorm.Open(m.resolve(name), &gorm.Config{})
    switch name {
    case "postgres": {
      sqlDB, err := db.DB()
      if err != nil {
        log.Fatalf("models.Setup db.DB() err: %v", err)
      }

      sqlDB.SetMaxIdleConns(10)
      sqlDB.SetMaxOpenConns(100)
      sqlDB.Stats()
    }
    }
    return db
  }
  return m.connections[name]
}

func (m *manager) resolve(name string) gorm.Dialector {
  switch name {
  case "sqlite3":
    return sqlite.Open(viper.GetString("db.driver.sqlite3.path"))
  case "postgres":
    return postgres.New(postgres.Config{
      DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s ",
        viper.GetString("db.driver.pgsql.host"),
        viper.GetString("db.driver.pgsql.username"),
        viper.GetString("db.driver.pgsql.password"),
        viper.GetString("db.driver.pgsql.database"),
      ),
    })
  }
  return nil
}

func (m *manager) getDefault() string {
  return viper.GetString("db.default")
}

func (m *manager) Migrate(name string, entities []interface{}) {
  if name == "" {
    name = m.getDefault()
  }

  switch name {
  case "sqlite3", "postgres":
    db := m.connections[name]
    _ = db.AutoMigrate(entities...)
  }
}

func (m *manager) Shutdown() {
  for key, connection := range m.connections {
    log.Printf("Disconnecting DB connection : %s \n", key)
    sqlDB, _ := connection.DB()
    sqlDB.Close()
  }
}
