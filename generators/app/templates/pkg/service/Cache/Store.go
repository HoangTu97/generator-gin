package Cache

type Store interface {
  Get(key string) interface{}
  Many(keys []string) []interface{}
  Add(key string, value interface{}, ttl int) bool
  Put(key string, value interface{}, ttl int) bool
  Increment(key string, value int) bool
  Decrement(key string, value int) bool
  Forever(key string, value interface{}) bool
  Forget(key string) bool
  Flush() bool
}