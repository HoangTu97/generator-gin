package Hash

type Service interface {
  Make(value string) string
  Check(value string, hashedValue string) bool
}

func NewService(driver string) Service {
  if driver == "bcrypt" {
    return NewBcrypt()
  }
  return NewBcrypt()
}