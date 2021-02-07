package controller

import (
  "<%= appName %>/dto/response"
  "<%= appName %>/pkg/converter"
  FileService "<%= appName %>/pkg/service/File"

  "fmt"
  "io/ioutil"
  "net/http"

  "github.com/gabriel-vasile/mimetype"
  "github.com/gin-gonic/gin"
)

type File interface {
  GetRoutes() []RouteController
  Upload(c *gin.Context)
  Download(c *gin.Context)
  FileDisplay(c *gin.Context)
}

type file struct {
  service FileService.Service
}

func NewFile(service FileService.Service) File {
  return &file{service: service}
}

func (r *file) GetRoutes() []RouteController {
  return []RouteController{
    RouteController{Method:http.MethodPost,Path:"/api/public/file/upload",Handler:r.Upload},
    RouteController{Method:http.MethodGet,Path:"/api/public/file/:id/download",Handler:r.Download},
    RouteController{Method:http.MethodGet,Path:"/api/public/file/:id",Handler:r.FileDisplay},
  }
}

// Upload upload file
// @Summary Upload
// @Tags PublicFile
// @Accept mpfd
// @Param file formData file true "Body with file"
// @Success 200 {object} response.APIResponseDTO{data=string} "desc"
// @Router /api/public/file/upload [post]
func (r *file) Upload(c *gin.Context) {
  // Source
  file, err := c.FormFile("file")
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  // baseFilename := filepath.Base(file.Filename)
  ext := r.service.Extension(file.Filename)

  filename := r.service.GenBaseName(ext)
  path := r.service.GetPath(filename)

  _ = r.service.MakeDirectory(r.service.GetPathDir(filename))
  if err := c.SaveUploadedFile(file, path); err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  response.CreateSuccesResponse(c, filename)
}

// Download download file
// @Summary Download
// @Tags PublicFile
// @Produce octet-stream
// @Param id path string true "File id"
// @Router /api/public/file/{id}/download [get]
func (r *file) Download(c *gin.Context) {
  filename := converter.MustString(c.Param("id"))

  filePath := r.service.GetPath(filename)

  c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
  // fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
  c.Writer.Header().Add("Content-Type", "application/octet-stream")
  c.File(filePath)
}

// Display display file
// @Summary Display
// @Tags PublicFile
// @Produce octet-stream
// @Param id path string true "File id"
// @Router /api/public/file/{id} [get]
func (r *file) FileDisplay(c *gin.Context) {
  filename := converter.MustString(c.Param("id"))

  filePath := r.service.GetPath(filename)

  b, err := ioutil.ReadFile(filePath) // just pass the file name
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  mime := mimetype.Detect(b)

  c.Data(200, mime.String(), b)
}
