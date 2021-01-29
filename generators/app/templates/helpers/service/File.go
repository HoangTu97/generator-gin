package service

type File interface {
  GenBaseName(extension string) string
  GetPath(fileName string) string
  GetPathDir(fileName string) string
}
