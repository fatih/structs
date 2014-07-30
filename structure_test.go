package structure

import (
	"reflect"
	"testing"
)

func TestMapNonStruct(t *testing.T) {
	foo := []string{"foo"}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Passing a non struct into Map should panic")
		}
	}()

	// this should panic. We are going to recover and and test it
	_ = Map(foo)
}

func TestMap(t *testing.T) {
	var T = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	a := Map(T)

	if typ := reflect.TypeOf(a).Kind(); typ != reflect.Map {
		t.Errorf("Map should return a map type, got: %v", typ)
	}

	// we have three fields
	if len(a) != 3 {
		t.Errorf("Map should return a map of len 3, got: %d", len(a))
	}

	inMap := func(val interface{}) bool {
		for _, v := range a {
			if reflect.DeepEqual(v, val) {
				return true
			}
		}

		return false
	}

	for _, val := range []interface{}{"a-value", 2, true} {
		if !inMap(val) {
			t.Errorf("Map should have the value %v", val)
		}
	}

}

func TestMap_Tag(t *testing.T) {
	var T = struct {
		A string `structure:"x"`
		B int    `structure:"y"`
		C bool   `structure:"z"`
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	a := Map(T)

	inMap := func(key interface{}) bool {
		for k := range a {
			if reflect.DeepEqual(k, key) {
				return true
			}
		}
		return false
	}

	for _, key := range []string{"x", "y", "z"} {
		if !inMap(key) {
			t.Errorf("Map should have the key %v", key)
		}
	}

}

func TestStruct(t *testing.T) {
	var T = struct{}{}

	if !IsStruct(T) {
		t.Errorf("T should be a struct, got: %T", T)
	}

	if !IsStruct(&T) {
		t.Errorf("T should be a struct, got: %T", T)
	}

}

func TestValues(t *testing.T) {
	var T = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	s := Values(T)

	if typ := reflect.TypeOf(s).Kind(); typ != reflect.Slice {
		t.Errorf("Values should return a slice type, got: %v", typ)
	}

	inSlice := func(val interface{}) bool {
		for _, v := range s {
			if reflect.DeepEqual(v, val) {
				return true
			}
		}
		return false
	}

	for _, val := range []interface{}{"a-value", 2, true} {
		if !inSlice(val) {
			t.Errorf("Values should have the value %v", val)
		}
	}
}

func TestFields(t *testing.T) {
	var T = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	s := Fields(T)

	if len(s) != 3 {
		t.Errorf("Fields should return a slice of len 3, got: %d", len(s))
	}

	inSlice := func(val string) bool {
		for _, v := range s {
			if reflect.DeepEqual(v, val) {
				return true
			}
		}
		return false
	}

	for _, val := range []string{"A", "B", "C"} {
		if !inSlice(val) {
			t.Errorf("Fields should have the value %v", val)
		}
	}
}

func TestIsValid(t *testing.T) {
	var T = struct {
		A string
		B int
		C bool `structure:"-"`
		D []string
	}{
		A: "a-value",
		B: 2,
	}

	ok := IsValid(T)
	if ok {
		t.Error("IsValid should return false because D is not initialized")
	}

	var X = struct {
		A string
		F *bool
	}{
		A: "a-value",
	}

	ok = IsValid(X)
	if ok {
		t.Error("IsValid should return false because F is not initialized")
	}
}

func TestName(t *testing.T) {
	type Foo struct {
		A string
		B bool
	}
	f := &Foo{}

	n := Name(f)
	if n != "Foo" {
		t.Error("Name should return Foo, got: %s", n)
	}
}
