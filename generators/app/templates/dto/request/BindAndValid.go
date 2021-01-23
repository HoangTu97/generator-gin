package request

import (
  "<%= appName %>/helpers/constants"
  "errors"
  "fmt"

  "github.com/astaxie/beego/validation"
  "github.com/gin-gonic/gin"
)

// BindAndValid bind data and return error if exist.
// @Return int errCode
func BindAndValid(c *gin.Context, form interface{}) (error) {
  err := c.Bind(form)
  if err != nil {
    return errors.New(constants.ErrorStringApi.INVALID_REQUEST)
  }

  valid := validation.Validation{}
  check, err := valid.Valid(form)
  if err != nil {
    return errors.New(constants.ErrorStringApi.INVALID_REQUEST)
  }
  if !check {
    MarkErrors(valid.Errors)
    return errors.New(constants.ErrorStringApi.INVALID_REQUEST)
  }

  return nil
}

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
  for _, err := range errors {
    // logging.Info(err.Key, err.Message)
    fmt.Printf("MarkErrors : %v %v", err.Key, err.Message)
  }
}
