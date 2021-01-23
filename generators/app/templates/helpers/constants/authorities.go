package constants

var ROLE role

type role struct {
  USER string
  ADMIN string
}

func init() {
  ROLE = role{
    USER: "USER",
    ADMIN: "ADMIN",
  }
}
