package Hash

type Service interface {
  Make(value string) string
  Check(hashedValue, value string) bool
}

func NewService(driver string) Service {
  switch driver {
  case "bcrypt":
    return NewBcrypt()
  case "md5":
    return NewMd5()
  }
  return NewBcrypt()
}