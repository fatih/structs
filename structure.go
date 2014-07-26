// Package structure contains various utilities functions to work with structs.
package structure

import (
	"errors"
	"reflect"
)

// ErrNotStruct is returned when the passed value is not a struct
var ErrNotStruct = errors.New("not struct")

// ToMap converts a struct to a map[string]interface{}. The default map key
// string is the struct fieldname but this can be changed by defining a
// "structure" tag key if needed. Note that only exported fields of a struct
// can be accessed, non exported fields will be neglected.
func ToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	t := reflect.TypeOf(in)

	// if pointer get the underlying element≤
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	v := reflect.ValueOf(in)

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
