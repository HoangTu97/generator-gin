package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/spf13/viper"
  "github.com/gin-gonic/gin"

  "<%= appName %>/config"
  "<%= appName %>/pkg/database"
  "<%= appName %>/pkg/logging"
  "<%= appName %>/pkg/service/Cache"
  "<%= appName %>/pkg/service/Mail"
  "<%= appName %>/routers"
)

// @title <%= appName %> API
// @version 1.0
// @description An example of gin
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

  database, closeDB := database.NewDB()
  defer closeDB()
  database = config.SetupDB(database)

  // logging.NewLogger(*config.LoggerSetting)
  logging.NewZeroLog()

  cacheManager := Cache.NewManager()
  mailManager := Mail.NewManager()

  jwtManager := config.SetupJWT()

  controllers := config.SetupController(database, jwtManager, cacheManager, mailManager)

  gin.ForceConsoleColor()
  gin.SetMode(viper.GetString("app.runMode"))

  router := routers.InitRouter(jwtManager, controllers)

  readTimeout := viper.GetDuration("app.readTimeout")
  writeTimeout := viper.GetDuration("app.writeTimeout")
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
