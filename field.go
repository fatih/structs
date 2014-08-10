package structure

import "reflect"

// Field represents a single struct field that encapsulates many high level
// function around a singel struct field
type Field struct {
	value reflect.Value
	field reflect.StructField
}

// Tag returns the value associated with key in the tag string. If there is no
// such key in the tag, Tag returns the empty string
func (f *Field) Tag(key string) string {
	return f.field.Tag.Get(key)
}

func (f *Field) Value() interface{} {
	return f.value.Interface()
}
