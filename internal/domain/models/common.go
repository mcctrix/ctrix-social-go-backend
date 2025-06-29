package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// StringArray is a custom type to properly handle string arrays with GORM and PostgreSQL
type StringArray []string

// Value implements the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	// Escape single quotes and backslashes
	escaped := make([]string, len(a))
	for i, s := range a {
		escaped[i] = strings.Replace(strings.Replace(s, "\\", "\\\\", -1), "'", "\\'", -1)
	}

	// Format as PostgreSQL array literal
	return fmt.Sprintf("{%s}", strings.Join(escaped, ",")), nil
}

// Scan implements the sql.Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case nil:
		*a = StringArray{}
		return nil
	default:
		return fmt.Errorf("unsupported type for StringArray: %T", value)
	}

	// Handle empty array
	if str == "{}" {
		*a = StringArray{}
		return nil
	}

	// Remove curly braces and split by comma
	trimmed := str[1 : len(str)-1]
	elements := strings.Split(trimmed, ",")

	// Unescape each element
	result := make([]string, len(elements))
	for i, e := range elements {
		result[i] = strings.Replace(strings.Replace(e, "\\'", "'", -1), "\\\\", "\\", -1)
	}

	*a = result
	return nil
}
