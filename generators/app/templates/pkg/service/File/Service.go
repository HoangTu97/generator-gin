package File

import (
  "fmt"
  "hash/fnv"
  "path"
  "path/filepath"
  "mime/multipart"
  "os"
  "io/ioutil"

  "github.com/satori/go.uuid"
)

type Service interface {
  GenBaseName(extension string) string
  GetPath(fileName string) string
  GetPathDir(fileName string) string
  Size(f multipart.File) (int, error)
  Extension(fileName string) string
  CheckNotExist(src string) bool
  CheckPermission(src string) bool
  // IsNotExistMkDir(src string) error
  Open(name string, flag int, perm os.FileMode) (*os.File, error)
  // MustOpen(fileName, filePath string) (*os.File, error)
  MakeDirectory(path string) bool
}

type service struct {
  location string
}

func NewService() Service {
  return &service{
    location: "./data/Files",
  }
}

func (s *service) GenBaseName(extension string) string {
  return uuid.Must(uuid.NewV4()).String() + extension
}

func (s *service) hash(str string) uint32 {
  h := fnv.New32a()
  _, _ = h.Write([]byte(str))
  return h.Sum32()
}

func (s *service) GetPath(fileName string) string {
  hash := s.hash(fileName)
  var mask uint32 = 255
  firstDir := hash & mask
  secondFir := (hash >> 8) & mask
  return filepath.Join(s.location, fmt.Sprintf("%02x", firstDir), fmt.Sprintf("%02x", secondFir), fileName)
}

func (s *service) GetPathDir(fileName string) string {
  return filepath.Dir(s.GetPath(fileName))
}

// Size get the file size
func (s *service) Size(f multipart.File) (int, error) {
  content, err := ioutil.ReadAll(f)
  return len(content), err
}

// Extension get the file ext
func (s *service) Extension(fileName string) string {
  return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func (s *service) CheckNotExist(src string) bool {
  _, err := os.Stat(src)
  return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func (s *service) CheckPermission(src string) bool {
  _, err := os.Stat(src)
  return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
// func (s *service) IsNotExistMkDir(src string) error {
//   if notExist := s.CheckNotExist(src); notExist == true {
//     if err := s.MakeDirectory(src); err != nil {
//       return err
//     }
//   }
//   return nil
// }

// Open a file according to a specific mode
func (s *service) Open(name string, flag int, perm os.FileMode) (*os.File, error) {
  f, err := os.OpenFile(name, flag, perm)
  if err != nil {
    return nil, err
  }
  return f, nil
}

func (s *service) MakeDirectory(path string) bool {
  err := os.MkdirAll(path, os.ModePerm)
  if err != nil {
    return false
  }
  return true
}

// MustOpen maximize trying to open the file
// func (s *service) MustOpen(fileName, filePath string) (*os.File, error) {
//   dir, err := os.Getwd()
//   if err != nil {
//     return nil, fmt.Errorf("os.Getwd err: %v", err)
//   }
//   src := dir + "/" + filePath
//   perm := s.CheckPermission(src)
//   if perm == true {
//     return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
//   }
//   err = s.IsNotExistMkDir(src)
//   if err != nil {
//     return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
//   }
//   f, err := s.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
//   if err != nil {
//     return nil, fmt.Errorf("Fail to OpenFile :%v", err)
//   }
//   return f, nil
// }
