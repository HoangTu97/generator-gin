package converter

import (
  "encoding/json"
  "fmt"
  "strconv"
  "strings"
)

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
  if i >= 0 && i < len(a) {
    r = a[i]
  } else if len(args) > 0 {
    r = args[0]
  }
  return
}

// ToStr Convert any type to string.
func ToStr(value interface{}, args ...int) (s string) {
  switch v := value.(type) {
  case bool:
    s = strconv.FormatBool(v)
  case float32:
    s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
  case float64:
    s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
  case int:
    s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
  case int8:
    s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
  case int16:
    s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
  case int32:
    s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
  case int64:
    s = strconv.FormatInt(v, argInt(args).Get(0, 10))
  case uint:
    s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
  case uint8:
    s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
  case uint16:
    s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
  case uint32:
    s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
  case uint64:
    s = strconv.FormatUint(v, argInt(args).Get(0, 10))
  case string:
    s = v
  case []byte:
    s = string(v)
  case []string:
    s = strings.Join(v, ",")
  default:
    b, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
      fmt.Println(err)
    }
    s = string(b)
  }
  return s
}

func ToStrs(values []interface{}, args ...int) []string {
  arr := make([]string, len(values))
  for i, value := range values {
    arr[i] = ToStr(value, args...)
  }
  return arr
}

// Int2HexStr converts decimal number to hex format string.
func Int2HexStr(num int) (hex string) {
  if num == 0 {
    return "0"
  }

  for num > 0 {
    r := num % 16

    var c string
    if r >= 0 && r <= 9 {
      c = string(rune(r + '0'))
    } else {
      c = string(rune(r + 'a' - 10))
    }
    hex = c + hex
    num = num / 16
  }
  return hex
}
