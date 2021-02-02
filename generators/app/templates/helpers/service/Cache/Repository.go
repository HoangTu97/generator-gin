package Cache

type Repository interface {
  Has(key string) bool
  Missing(key string) bool
  Get(key string) interface{}
  Set(key string, value interface{}, ttl int) bool
  Delete(key string) bool
  Clear() bool
  Many(keys []string) map[string]interface{}
  PutMany(values map[string]interface{}, ttl int) bool
  DeleteMultiple(keys []string) bool

  Pull(key string) interface{}
  Put(key string, value interface{}, ttl int) bool
  Add(key string, value interface{}, ttl int) bool
  Increment(key string, value int) bool
  Increment1(key string) bool
  Decrement(key string, value int) bool
  Decrement1(key string) bool
  Forever(key string, value interface{}) bool
  Forget(key string) bool

  // GetStore() Store
}

type repository struct {
  store Store
}

func NewRepository(store Store) Repository {
  return &repository{store: store}
}

func (r *repository) Has(key string) bool {
  return r.Get(key) != nil
}

func (r *repository) Missing(key string) bool {
  return !r.Has(key)
}

func (r *repository) Get(key string) interface{} {
  return r.store.Get(key)
}

func (r *repository) Set(key string, value interface{}, ttl int) bool {
  return r.Put(key, value, ttl)
}

func (r *repository) Delete(key string) bool {
  return r.Forget(key)
}

func (r *repository) Clear() bool {
  return r.store.Flush()
}

func (r *repository) Many(keys []string) map[string]interface{} {
  return make(map[string]interface{}, 0)
}

func (r *repository) PutMany(values map[string]interface{}, ttl int) bool {
  return true
}

func (r *repository) DeleteMultiple(keys []string) bool {
  result := true
  for _, key := range keys {
    if !r.Forget(key) {
      result = false
    }
  }
  return result
}

func (r *repository) Pull(key string) interface{} {
  value := r.Get(key)
  r.Forget(key)
  return value
}

func (r *repository) Put(key string, value interface{}, ttl int) bool {
  if ttl == 0 {
    return r.Forever(key, value)
  }
  if ttl < 0 {
    return r.Forget(key)
  }
  return r.store.Put(key, value, ttl)
}

func (r *repository) Add(key string, value interface{}, ttl int) bool {
  if ttl == 0 {
    if r.Get(key) == nil {
      return r.Put(key, value, ttl)
    }
    return false
  }
  if ttl < 0 {
    return false
  }
  return r.store.Add(key, value, ttl)
}

func (r *repository) Increment(key string, value int) bool {
  return r.store.Increment(key, value)
}

func (r *repository) Increment1(key string) bool {
  return r.Increment(key, 1)
}

func (r *repository) Decrement(key string, value int) bool {
  return r.store.Decrement(key, value)
}

func (r *repository) Decrement1(key string) bool {
  return r.Decrement(key, 1)
}

func (r *repository) Forever(key string, value interface{}) bool {
  return r.store.Forever(key, value)
}

func (r *repository) Forget(key string) bool {
  return r.store.Forget(key)
}

// func (r *repository) GetStore() Store {
//   return r.store
// }
