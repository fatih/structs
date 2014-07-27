// Package structure contains various utilities functions to work with structs.
package structure

import (
	"errors"
	"reflect"
	"sort"
)

// ErrNotStruct is returned when the passed value is not a struct
var ErrNotStruct = errors.New("not struct")

// ToMap converts a struct to a map[string]interface{}. The default map key
// string is the struct fieldname but this can be changed by defining a
// "structure" tag key if needed. Note that only exported fields of a struct
// can be accessed, non exported fields will be neglected.
func ToMap(s interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// we can't access the value of unexported fields
		if field.PkgPath != "" {
			continue
		}

		name := field.Name

		// override if the user passed a structure tag
		if tag := field.Tag.Get("structure"); tag != "" {
			name = tag
		}

		out[name] = v.Field(i).Interface()
	}

	return out, nil
}

// ToSlice converts a struct's field values to a []interface{}. Values are
// inserted and sorted according to the field names. Note that only exported
// fields of a struct can be accessed, non exported fields  will be neglected.
func ToSlice(s interface{}) ([]interface{}, error) {
	m, err := ToMap(s)
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(m))
	count := 0
	for k := range m {
		keys[count] = k
		count++
	}

	sort.Strings(keys)

	t := make([]interface{}, len(m))

	for i, key := range keys {
		t[i] = m[key]
	}

	return t, nil

}

// IsStruct returns true if the given variable is a struct or a pointer to
// struct.
func IsStruct(s interface{}) bool {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Kind() == reflect.Struct
}
