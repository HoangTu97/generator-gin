package CacheStore

type nullStore struct {
}

func NewNull() *nullStore {
  return &nullStore{}
}

func (s *nullStore) Get(key string) interface{} {
  return nil
}

func (s *nullStore) Many(keys []string) []interface{} {
  return make([]interface{}, 0)
}

func (s *nullStore) Add(key string, value interface{}, ttl int) bool {
  return false
}

func (s *nullStore) Put(key string, value interface{}, ttl int) bool {
  return false
}

func (s *nullStore) Increment(key string, value int) bool {
  return false
}

func (s *nullStore) Decrement(key string, value int) bool {
  return false
}

func (s *nullStore) Forever(key string, value interface{}) bool {
  return false
}

func (s *nullStore) Forget(key string) bool {
  return true
}

func (s *nullStore) Flush() bool {
  return true
}
