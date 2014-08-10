package structure

import "reflect"

// Field represents a single struct field that encapsulates high level
// functions around the field.
type Field struct {
	value reflect.Value
	field reflect.StructField
}

// Tag returns the value associated with key in the tag string. If there is no
// such key in the tag, Tag returns the empty string.
func (f *Field) Tag(key string) string {
	return f.field.Tag.Get(key)
}

// Value returns the underlying value of of the field.
func (f *Field) Value() interface{} {
	return f.value.Interface()
}

// IsEmbedded returns true if the given field is an anonymous field (embedded)
func (f *Field) IsEmbedded() bool {
	return f.field.Anonymous
}
