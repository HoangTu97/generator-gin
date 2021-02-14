package main

import (
  "time"
  "fmt"
  "log"
  "net/http"

  "github.com/spf13/viper"
  "github.com/gin-gonic/gin"

  "<%= appName %>/config"
  "<%= appName %>/pkg/service/Cache"
  "<%= appName %>/pkg/service/Database"
  "<%= appName %>/pkg/service/Jwt"
  "<%= appName %>/pkg/service/Mail"
  "<%= appName %>/routers"
)

// @title <%= appName %> API
// @version 1.0
// @description An example of gin
// @host localhost:<%= serverPort %>
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
  viper.SetConfigName("config")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig() // Find and read the config file
  if err != nil { // Handle errors reading the config file
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
  }

  dbManager := Database.NewManager()
  defer dbManager.Shutdown()
  models := config.GetModelsNeedMigrate()
  dbManager.Connection("")
  dbManager.Migrate("", models)

  cacheManager := Cache.NewManager()
  mailManager := Mail.NewManager()
  jwtManager := Jwt.NewManager()

  controllers := config.Providers(dbManager, jwtManager, cacheManager, mailManager)

  gin.ForceConsoleColor()
  gin.SetMode(viper.GetString("app.runMode"))

  router := routers.InitRouter(jwtManager, controllers)

  readTimeout := viper.GetDuration("app.readTimeout")*time.Second
  writeTimeout := viper.GetDuration("app.writeTimeout")*time.Second
  endPoint := fmt.Sprintf(":%s", viper.GetString("app.port"))
  maxHeaderBytes := 1 << 20

  server := &http.Server{
    Addr:           endPoint,
    Handler:        router,
    ReadTimeout:    readTimeout,
    WriteTimeout:   writeTimeout,
    MaxHeaderBytes: maxHeaderBytes,
  }

  log.Printf("[info] start http server listening %s", endPoint)

  if viper.GetBool("app.SSL") {
    SSLKeys := &struct {
      CERT string
      KEY  string
    }{}

    //Generated using sh generate-certificate.sh
    SSLKeys.CERT = "./cert/myCA.cer"
    SSLKeys.KEY = "./cert/myCA.key"

    err := server.ListenAndServeTLS(SSLKeys.CERT, SSLKeys.KEY)
    if err != nil {
      log.Fatal("Web server (HTTPS): ", err)
    }
  } else {
    err := server.ListenAndServe()
    if err != nil {
      log.Fatal("Web server (HTTP): ", err)
    }
  }

}
