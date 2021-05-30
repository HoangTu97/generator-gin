package converter

import (
  "strconv"
  "strings"
  "time"
)

func Exist(f string) bool {
  return f != string(rune(0x1E))
}

func Uint(f string) (uint, error) {
  v, err := strconv.ParseUint(MustString(f), 10, 32)
  return uint(v), err
}

func Uint8(f string) (uint8, error) {
  v, err := strconv.ParseUint(MustString(f), 10, 8)
  return uint8(v), err
}

func Uint16(f string) (uint16, error) {
  v, err := strconv.ParseUint(MustString(f), 10, 16)
  return uint16(v), err
}

func Uint32(f string) (uint32, error) {
  v, err := strconv.ParseUint(MustString(f), 10, 32)
  return uint32(v), err
}

func Uint64(f string) (uint64, error) {
  v, err := strconv.ParseUint(MustString(f), 10, 64)
  return uint64(v), err
}

func Int(f string) (int, error) {
  v, err := strconv.ParseInt(MustString(f), 10, 0)
  return int(v), err
}

func Int8(f string) (int8, error) {
  v, err := strconv.ParseInt(MustString(f), 10, 8)
  return int8(v), err
}

func Int16(f string) (int16, error) {
  v, err := strconv.ParseInt(MustString(f), 10, 16)
  return int16(v), err
}

func Int32(f string) (int32, error) {
  v, err := strconv.ParseInt(MustString(f), 10, 32)
  return int32(v), err
}

func Int64(f string) (int64, error) {
  v, err := strconv.ParseInt(MustString(f), 10, 64)
  return int64(v), err
}

func Float64(f string) (float64, error) {
  v, err := strconv.ParseFloat(MustString(f), 64)
  return float64(v), err
}

func Bool(f string) (bool, error) {
  v, err := strconv.ParseBool(MustString(f))
  return v, err
}

func ArrStr(f string) ([]string, error) {
  return strings.Split(f, ","), nil
}

const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)

func Time(f string, layout string) (time.Time, error) {
  t, err := time.Parse(layout, f)
  if err != nil {
    return time.Time{}, err
  }
  return t, nil
}

// // HexStr2int converts hex format string to decimal number.
// func HexStr2int(hexStr string) (int, error) {
// 	num := 0
// 	length := len(hexStr)
// 	for i := 0; i < length; i++ {
// 		char := hexStr[length-i-1]
// 		factor := -1

// 		switch {
// 		case char >= '0' && char <= '9':
// 			factor = int(char) - '0'
// 		case char >= 'a' && char <= 'f':
// 			factor = int(char) - 'a' + 10
// 		default:
// 			return -1, fmt.Errorf("invalid hex: %s", string(char))
// 		}

// 		num += factor * PowInt(16, i)
// 	}
// 	return num, nil
// }
