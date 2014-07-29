// Package structure contains various utilities functions to work with structs.
package structure

import (
	"reflect"
	"sort"
)

// Map converts the given s struct to a map[string]interface{}, where the keys
// of the map are the field names and the values of the map the associated
// values of the fields. The default key string is the struct field name but
// can be changed in the struct field's tag value. The "structure" key in the
// struct's field tag value is the key name. Example:
//
//   // Field appears in map as key "myName".
//   Name string `structure:"myName"`
//
// A value with the content of "-" ignores that particular field. Example:
//
//   // Field is ignored by this package.
//   Field bool `structure:"-"`
//
// Note that only exported fields of a struct can be accessed, non exported
// fields will be neglected. It panics if s's kind is not struct.
func Map(s interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	v, fields := strctInfo(s)

	for i, field := range fields {
		name := field.Name

		// override if the user passed a structure tag value
		// ignore if the user passed the "-" value
		if tag := field.Tag.Get("structure"); tag != "" {
			if tag == "-" {
				continue
			}

			name = tag
		}

		out[name] = v.Field(i).Interface()
	}

	return out
}

// Values converts the given s struct's field values to a []interface{}.
// Values are inserted and sorted according to the field names. A struct tag
// with the content of "-" ignores the that particular field. Example:
//
//   // Field is ignored by this package.
//   Field int `structure:"-"`
//
// Note that only exported fields of a struct can be accessed, non exported
// fields  will be neglected.  It panics if s's kind is not struct.
func Values(s interface{}) []interface{} {
	m := Map(s)

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

// IsValid returns true if all fields in a struct are initialized (non zero
// value). A struct tag with the content of "-" ignores the checking of that
// particular field. Example:
//
//   // Field is ignored by this package.
//   Field bool `structure:"-"`
//
// Note that only exported fields of a struct can be accessed, non exported
// fields  will be neglected. It panics if s's kind is not struct.
func IsValid(s interface{}) bool {
	v, fields := strctInfo(s)

	for i, field := range fields {
		// don't check if it's omitted
		if tag := field.Tag.Get("structure"); tag == "-" {
			continue
		}

		// zero value of the given field, such as "" for string, 0 for int
		zero := reflect.Zero(v.Field(i).Type()).Interface()

		//  current value of the given field
		current := v.Field(i).Interface()

		if reflect.DeepEqual(current, zero) {
			return false
		}
	}

	return true
}

// Fields returns a sorted slice of field names. A struct tag with the content
// of "-" ignores the checking of that particular field. Example:
//
//   // Field is ignored by this package.
//   Field bool `structure:"-"`
//
// Note that only exported fields of a struct can be accessed, non exported
// fields  will be neglected.
func Fields(s interface{}) []string {
	m := Map(s)

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

// strctInfo returns the struct value and the exported struct fields for a
// given s struct. This is a convenient helper method to avoid duplicate code
// in some of the functions.
func strctInfo(s interface{}) (reflect.Value, []reflect.StructField) {
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	t := v.Type()

	f := make([]reflect.StructField, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		// we can't access the value of unexported fields
		if t.Field(i).PkgPath != "" {
			continue
		}

		f[i] = t.Field(i)
	}

	return v, f
}
