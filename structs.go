package structure

// Struct encapsulates a struct type to provide several high level functions
// around the structure.
type Struct struct {
	raw interface{}
}

// New returns a new *Struct with the struct s.
func New(s interface{}) *Struct {
	return &Struct{
		raw: s,
	}
}

// Map converts the given struct to a map[string]interface{}. For more info
// refer to Map() function.
func (s *Struct) Map() map[string]interface{} {
	return Map(s.raw)
}

// Values converts the given struct to a []interface{}. For more info refer to
// Values() function.
func (s *Struct) Values() []interface{} {
	return Values(s.raw)
}

// Fields returns a slice of field names For more info refer to Fields()
// function.
func (s *Struct) Fields() []string {
	return Fields(s.raw)
}

// Field returns a new Field struct that provides several high level functions
// around a single struct field entitiy. It panics if the field is not found or
// is unexported.
func (s *Struct) Field(name string) *Field {
	f, ok := s.FieldOk(name)
	if !ok {
		panic("field not found")
	}

	return f
}

// Field returns a new Field struct that provides several high level functions
// around a single struct field entitiy and a boolean indicating if the field
// was found. It panics if the or is unexported.
func (s *Struct) FieldOk(name string) (*Field, bool) {
	v := strctVal(s.raw)
	t := v.Type()

	field, ok := t.FieldByName(name)
	if !ok {
		return nil, false
	}

	if field.PkgPath != "" {
		panic("unexported field access is not allowed")
	}

	return &Field{
		field: field,
		value: v.FieldByName(name),
	}, true
}

// IsZero returns true if all fields is equal to a zero value. For more info
// refer to IsZero() function.
func (s *Struct) IsZero() bool {
	return IsZero(s.raw)
}

// HasZero returns true if any field is equal to a zero value. For more info
// refer to HasZero() function.
func (s *Struct) HasZero() bool {
	return HasZero(s.raw)
}

// Name returns the structs's type name within its package. For more info refer
// to Name() function.
func (s *Struct) Name() string {
	return Name(s)
}

// IsStruct returns true if its a struct or a pointer to struct.
func (s *Struct) IsStruct() bool {
	return IsStruct(s.raw)
}
