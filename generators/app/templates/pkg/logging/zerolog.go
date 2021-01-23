package logging

import (
  "github.com/gin-gonic/gin"
  "github.com/rs/zerolog"
)

func NewZeroLog() {
  zerolog.SetGlobalLevel(zerolog.InfoLevel)
    if gin.IsDebugging() {
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
  }
}