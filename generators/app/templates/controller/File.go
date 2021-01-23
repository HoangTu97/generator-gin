package controller

import (
  "<%= appName %>/dto/response"
  fileUtil "<%= appName %>/helpers/file"
  "<%= appName %>/pkg/converter"
  "<%= appName %>/service"

  "fmt"
  "io/ioutil"
  "path/filepath"

  "github.com/gabriel-vasile/mimetype"
  "github.com/gin-gonic/gin"
)

type File interface {
  Upload(c *gin.Context)
  Download(c *gin.Context)
  FileDisplay(c *gin.Context)
}

type file struct {
  service service.File
}

func NewFile(service service.File) File {
  return &file{service: service}
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
  ext := filepath.Ext(file.Filename)

  filename := r.service.GenFileBaseFileName(ext)
  path := r.service.GetFilePath(filename)

  _ = fileUtil.MkDir(r.service.GetFilePathDir(filename))
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

  filePath := r.service.GetFilePath(filename)

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

  filePath := r.service.GetFilePath(filename)

  b, err := ioutil.ReadFile(filePath) // just pass the file name
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  mime := mimetype.Detect(b)

  c.Data(200, mime.String(), b)
}
