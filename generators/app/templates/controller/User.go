package controller

import (
  "<%= appName %>/dto"
  "<%= appName %>/dto/request"
  UserRequest "<%= appName %>/dto/request/user"
  "<%= appName %>/dto/response"
  UserResponse "<%= appName %>/dto/response/user"
  "<%= appName %>/helpers/constants"
  "<%= appName %>/service"

  "github.com/gin-gonic/gin"
)

type User interface {
  Register(c *gin.Context)
  Login(c *gin.Context)
  UserInfo(c *gin.Context)
}

type user struct {
  service service.User
}

func NewUser(service service.User) User {
  return &user{service: service}
}

// Register register
// @Summary Register
// @Tags PublicUser
// @Accept  json
// @Param RegisterDTO body requestuser.RegisterDTO true "RegisterDTO"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/public/user/register [post]
func (r *user) Register(c *gin.Context) {
  var registerDTO UserRequest.RegisterDTO
  err := request.BindAndValid(c, &registerDTO)
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  _, checkRegistered := r.service.FindOneByUsername(registerDTO.Username)
  if checkRegistered {
    response.CreateErrorResponse(c, constants.ErrorStringApi.USER_EXISTED)
    return
  }

  userDTO := dto.UserDTO{Name: registerDTO.Username, Password: registerDTO.Password}

  userDTO, isSuccess := r.service.Create(userDTO)
  if !isSuccess {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  tokenString, error := r.service.GenerateToken(userDTO.UserID, userDTO.Name, userDTO.GetRolesStr())
  if error != nil {
    response.CreateErrorResponse(c, constants.ErrorStringApi.USER_TOKEN_GEN_FAILED)
    return
  }

  response.CreateSuccesResponse(c, UserResponse.RegisterResponseDTO{
    Token: tokenString,
  })
}

// Login login
// @Summary Login
// @Tags PublicUser
// @Accept  json
// @Param LoginDTO body requestuser.LoginDTO true "LoginDTO"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/public/user/login [post]
func (r *user) Login(c *gin.Context) {
  var loginDTO UserRequest.LoginDTO
  err := request.BindAndValid(c, &loginDTO)
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  userDTO, isSuccess := r.service.FindOneLogin(loginDTO.Username, loginDTO.Password)
  if !isSuccess {
    response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
    return
  }

  tokenString, err := r.service.GetUserToken(userDTO)
  if err != nil {
    response.CreateErrorResponse(c, constants.ErrorStringApi.USER_TOKEN_GEN_FAILED)
    return
  }

  response.CreateSuccesResponse(c, UserResponse.LoginResponseDTO{
    Token: tokenString,
  })
}

func (r *user) UserInfo(c *gin.Context) {
}
