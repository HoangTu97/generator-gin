package controller

import (
  "<%= appName %>/dto"
  "<%= appName %>/dto/request"
  <%= entityCap %>Request "<%= appName %>/dto/request/<%= entityLower %>"
  "<%= appName %>/dto/response"
  <%= entityCap %>Response "<%= appName %>/dto/response/<%= entityLower %>"
  "<%= appName %>/helpers/constants"
  "<%= appName %>/helpers/pagination"
  "<%= appName %>/pkg/converter"
  "<%= appName %>/service"

  "github.com/gin-gonic/gin"
)

type <%= entityCap %> interface {
  GetAll(c *gin.Context)
  GetDetails(c *gin.Context)
  Create(c *gin.Context)
  Update(c *gin.Context)
  Delete(c *gin.Context)
}

type <%= entityLower %> struct {
  service service.<%= entityCap %>
}

func New<%= entityCap %>(service service.<%= entityCap %>) <%= entityCap %> {
  return &<%= entityLower %>{service: service}
}

// <%= entityCap %> all
// @Summary GetAll
// @Tags Public<%= entityCap %>
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=<%= entityCap %>Response.ListResponseDTO} "desc"
// @Router /api/public/<%= entityLower %> [get]
func (r *<%= entityLower %>) GetAll(c *gin.Context) {
  pageable := pagination.GetPage(c)

  page := r.service.FindPage(pageable)

  response.CreateSuccesResponse(c, <%= entityCap %>Response.CreateListResponseDTOFromPage(page))
}

// <%= entityCap %> details
// @Summary GetDetails
// @Tags Public<%= entityCap %>
// @Accept json
// @Param id path int true "<%= entityCap %> ID"
// @Success 200 {object} response.APIResponseDTO{data=dto.<%= entityCap %>DTO} "desc"
// @Router /api/public/<%= entityLower %>/{id} [get]
func (r *<%= entityLower %>) GetDetails(c *gin.Context) {
  id := converter.MustUint(c.Param("id"))
  
  <%= entityLower %>DTO, success := r.service.FindOne(id)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  response.CreateSuccesResponse(c, <%= entityLower %>DTO)
}

// <%= entityCap %> create
// @Summary Create
// @Tags Private<%= entityCap %>
// @Accept json
// @Security ApiKeyAuth
// @Param body body <%= entityCap %>Request.CreateRequestDTO true "body"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/private/<%= entityLower %> [post]
func (r *<%= entityLower %>) Create(c *gin.Context) {
  var success bool

  var requestDTO <%= entityCap %>Request.CreateRequestDTO
  err := request.BindAndValid(c, &requestDTO)
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  <%= entityLower %>DTO := dto.<%= entityCap %>DTO{}

  _, success = r.service.Save(<%= entityLower %>DTO)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  response.CreateSuccesResponse(c, nil)
}

// <%= entityCap %> update
// @Summary Update
// @Tags Private<%= entityCap %>
// @Accept json
// @Security ApiKeyAuth
// @Param id path int true "<%= entityCap %> ID"
// @Param body body <%= entityCap %>Request.CreateRequestDTO true "body"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/private/<%= entityLower %>/{id} [put]
func (r *<%= entityLower %>) Update(c *gin.Context) {
  id := converter.MustUint(c.Param("id"))

  var requestDTO <%= entityCap %>Request.UpdateRequestDTO
  err := request.BindAndValid(c, &requestDTO)
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  var success bool
  <%= entityLower %>DTO, success := r.service.FindOne(id)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  _, success = r.service.Save(<%= entityLower %>DTO)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  response.CreateSuccesResponse(c, nil)
}

// <%= entityCap %> delete
// @Summary Delete
// @Tags Private<%= entityCap %>
// @Accept json
// @Security ApiKeyAuth
// @Param id path int true "<%= entityCap %> ID"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/private/<%= entityLower %>/{id} [delete]
func (r *<%= entityLower %>) Delete(c *gin.Context) {
  id := converter.MustUint(c.Param("id"))

  _, found := r.service.FindOne(id)
  if !found {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INVALID_REQUEST)
    return
  }

  r.service.Delete(id)
  response.CreateSuccesResponse(c, nil)
}
