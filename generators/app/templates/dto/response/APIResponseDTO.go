package response

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// APIResponseDTO godoc
type APIResponseDTO struct {
  Success bool        `json:"success"`
  Error   string      `json:"error"`
  Data    interface{} `json:"data"`
}

// CreateSuccesResponse godoc
func CreateSuccesResponse(c *gin.Context, data interface{}) {
  c.JSON(http.StatusOK, APIResponseDTO{
    Success: true,
    Error:   "",
    Data:    data,
  })
}

// CreateErrorResponse godoc
func CreateErrorResponse(c *gin.Context, err string) {
  c.JSON(http.StatusOK, APIResponseDTO{
    Success: false,
    Error:   err,
    Data:    nil,
  })
}
