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

	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	t := v.Type()
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
	v := reflect.ValueOf(s)

	// if pointer get the underlying element≤
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// we can't access the value of unexported fields
		if field.PkgPath != "" {
			continue
		}

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
