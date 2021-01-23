package domain

import (
  "<%= appName %>/helpers/constants"
  "bytes"
  "database/sql/driver"
  "encoding/json"
  "fmt"
)

type UserRole int

const (
  ROLE_USER UserRole = iota
  ROLE_ADMIN
)

var toString = []string{
  constants.ROLE.USER,
  constants.ROLE.ADMIN,
}

func (role UserRole) String() string {
  return toString[role]
}

func (role UserRole) MarshalJSON() ([]byte, error) {
  buffer := bytes.NewBufferString(`"`)
  buffer.WriteString(toString[role])
  buffer.WriteString(`"`)
  return buffer.Bytes(), nil
}

func (role *UserRole) UnmarshalJSON(data []byte) error {
  var str string
  err1 := json.Unmarshal(data, &str)
  if err1 != nil {
    return err1
  }
  _role, err := ParseUserRole(str)
  if err != nil {
    return err
  }
  *role = _role
  return nil
}

func (role *UserRole) Scan(value interface{}) error {
  _role, err := ParseUserRole(value.(string))
  if err != nil {
    return err
  }
  *role = _role
  return nil
}

func (role *UserRole) Value() (driver.Value, error) {
  return role.String(), nil
}

func ParseUserRole(value string) (UserRole, error) {
  for i, v := range toString {
    if v == value {
      return UserRole(i), nil
    }
  }
  return ROLE_USER, fmt.Errorf("%s is not a valid UserRole", value)
}
