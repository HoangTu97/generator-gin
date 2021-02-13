package constants

var CACHE cache

type cache struct {
  USER string
  IMAGE string
}

func init() {
  CACHE = cache{
    USER: "USER",
    IMAGE: "IMAGE",
  }
}
