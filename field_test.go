package structure

import "testing"

// A test struct that defines all cases
type Foo struct {
	A    string
	B    int    `structure:"y"`
	C    bool   `json:"c"`
	d    string // not exported
	x    string `xml:"x"` // not exported, with tag
	*Bar        // embedded
}

type Bar struct {
	E string
	F int
	g []string
}

func newStruct() *Struct {
	b := &Bar{
		E: "example",
		F: 2,
		g: []string{"zeynep", "fatih"},
	}

	// B and x is not initialized for testing
	f := &Foo{
		A: "gopher",
		C: true,
		d: "small",
	}
	f.Bar = b

	return New(f)
}

func TestField(t *testing.T) {
	s := newStruct()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Retrieveing a non existing field from the struct should panic")
		}
	}()

	_ = s.Field("no-field")
}

func TestField_Tag(t *testing.T) {
	s := newStruct()

	v := s.Field("B").Tag("json")
	if v != "" {
		t.Errorf("Field's tag value of a non existing tag should return empty, got: %s", v)
	}

	v = s.Field("C").Tag("json")
	if v != "c" {
		t.Errorf("Field's tag value of the existing field C should return 'c', got: %s", v)
	}

	v = s.Field("d").Tag("json")
	if v != "" {
		t.Errorf("Field's tag value of a non exported field should return empty, got: %s", v)
	}

	v = s.Field("x").Tag("xml")
	if v != "x" {
		t.Errorf("Field's tag value of a non exported field with a tag should return 'x', got: %s", v)
	}

	v = s.Field("A").Tag("json")
	if v != "" {
		t.Errorf("Field's tag value of a existing field without a tag should return empty, got: %s", v)
	}
}

func TestField_Value(t *testing.T) {
	s := newStruct()

	v := s.Field("A").Value()
	val, ok := v.(string)
	if !ok {
		t.Errorf("Field's value of a A should be string")
	}

	if val != "gopher" {
		t.Errorf("Field's value of a existing tag should return 'gopher', got: %s", val)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Value of a non exported field from the field should panic")
		}
	}()

	// should panic
	_ = s.Field("d").Value()
}

func TestField_IsEmbedded(t *testing.T) {
	s := newStruct()

	if !s.Field("Bar").IsEmbedded() {
		t.Errorf("Fields 'Bar' field is an embedded field")
	}

	if s.Field("d").IsEmbedded() {
		t.Errorf("Fields 'd' field is not an embedded field")
	}
}

func TestField_IsExported(t *testing.T) {
	s := newStruct()

	if !s.Field("Bar").IsExported() {
		t.Errorf("Fields 'Bar' field is an exported field")
	}

	if !s.Field("A").IsExported() {
		t.Errorf("Fields 'A' field is an exported field")
	}

	if s.Field("d").IsExported() {
		t.Errorf("Fields 'd' field is not an exported field")
	}
}

func TestField_IsZero(t *testing.T) {
	s := newStruct()

	if s.Field("A").IsZero() {
		t.Errorf("Fields 'A' field is an initialized field")
	}

	if !s.Field("B").IsZero() {
		t.Errorf("Fields 'B' field is not an initialized field")
	}
}

func TestField_Name(t *testing.T) {
	s := newStruct()

	if s.Field("A").Name() != "A" {
		t.Errorf("Fields 'A' field should have the name 'A'")
	}
}
