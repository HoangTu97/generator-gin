package service

type File interface {
  GenFileBaseFileName(extension string) string
  GetFilePath(fileName string) string
  GetFilePathDir(fileName string) string
}
