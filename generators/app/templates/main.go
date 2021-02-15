package main

import (
  "context"
  "fmt"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/spf13/viper"
  "github.com/gin-gonic/gin"

  "<%= appName %>/config"
  "<%= appName %>/pkg/service/Cache"
  "<%= appName %>/pkg/service/Database"
  "<%= appName %>/pkg/service/Jwt"
  "<%= appName %>/pkg/service/Mail"
  "<%= appName %>/pkg/service/Log"
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
  models := config.GetModelsNeedMigrate()
  dbManager.Connection("")
  dbManager.Migrate("", models)

  logManager := Log.NewManager()
  logger := logManager.Driver("")
  cacheManager := Cache.NewManager(logger)
  mailManager := Mail.NewManager(logger)
  jwtManager := Jwt.NewManager()

  controllers := config.Providers(dbManager, jwtManager, cacheManager, mailManager, logManager)

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

  logger.Info("Start http server listening", endPoint)

  if viper.GetBool("app.SSL") {
    SSLKeys := &struct {
      CERT string
      KEY  string
    }{}

    //Generated using sh generate-certificate.sh
    SSLKeys.CERT = "./cert/myCA.cer"
    SSLKeys.KEY = "./cert/myCA.key"

    go func() {
      err := server.ListenAndServeTLS(SSLKeys.CERT, SSLKeys.KEY)
      if err != nil {
        logger.Fatal("Web server (HTTPS): ", err)
      }
    }()
  } else {
    go func() {
      err := server.ListenAndServe()
      if err != nil {
        logger.Fatal("Web server (HTTP): ", err)
      }
    }()
  }

  // Wait for interrupt signal to gracefully shutdown the server with
  // a timeout of 5 seconds.
  quit := make(chan os.Signal)
  // kill (no param) default send syscall.SIGTERM
  // kill -2 is syscall.SIGINT
  // kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  logger.Info("Shutting down server...")

  // The context is used to inform the server it has 5 seconds to finish
  // the request it is currently handling
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  logger.Info("Server exiting")
  logManager.Shutdown()
  dbManager.Shutdown()

  if err := server.Shutdown(ctx); err != nil {
    log.Fatal("Server forced to shutdown:", err)
  }
}
