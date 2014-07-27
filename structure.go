// Package structure contains various utilities functions to work with structs.
package structure

import (
	"reflect"
	"sort"
)

// ToMap converts the given s struct to a map[string]interface{}. The default
// map key names are the struct fieldnames but this can be changed by defining
// a "structure" tag key if needed. Note that only exported fields of a struct
// can be accessed, non exported fields will be neglected. It panics if s's
// kind is not struct.
func ToMap(s interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("not struct")
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

	return out
}

// ToSlice converts the given s struct's field values to a []interface{}.
// Values are inserted and sorted according to the field names. Note that only
// exported fields of a struct can be accessed, non exported fields  will be
// neglected.  It panics if s's kind is not struct.
func ToSlice(s interface{}) []interface{} {
	m := ToMap(s)

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

	return t

}

// Fields returns a sorted slice of field names. Note that only exported
// fields of a struct can be accessed, non exported fields  will be neglected.
func Fields(s interface{}) []string {
	m := ToMap(s)

	keys := make([]string, len(m))
	count := 0
	for k := range m {
		keys[count] = k
		count++
	}

	sort.Strings(keys)

	return keys
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
