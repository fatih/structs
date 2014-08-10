package structure

import "testing"

// A test struct that defines all cases
type Foo struct {
	A    string
	B    int    `structure:"y"`
	C    bool   `json:"c"`
	d    string // not exported
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

	f := &Foo{
		A: "gopher",
		B: 1,
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
