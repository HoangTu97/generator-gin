package converter

func MustUint(f string) uint {
  v, _ := Uint(f)
  return v
}

func MustUint8(f string) uint8 {
  v, _ := Uint8(f)
  return v
}

func MustUint32(f string) uint32 {
  v, _ := Uint32(f)
  return v
}

func MustUint64(f string) uint64 {
  v, _ := Uint64(f)
  return v
}

func MustInt(f string) int {
  v, _ := Int(f)
  return v
}

func MustInt64(f string) int64 {
  v, _ := Int64(f)
  return v
}

func MustFloat64(f string) float64 {
  v, _ := Float64(f)
  return v
}

func MustString(f string) string {
  if Exist(f) {
    return f
  }
  return ""
}

func MustBool(f string) bool {
  v, _ := Bool(f)
  return v
}

func MustArrStr(f string) []string {
  if !Exist(f) {
    return []string{}
  }
  v, _ := ArrStr(f)
  return v
}
