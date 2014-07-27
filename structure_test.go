package structure

import (
	"reflect"
	"testing"
)

func TestToMapNonStruct(t *testing.T) {
	foo := []string{"foo"}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("Passing a non struct into ToMap should panic")
		}
	}()

	// this should panic. We are going to recover and and test it
	_ = ToMap(foo)
}

func TestToMap(t *testing.T) {
	var T = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	a := ToMap(T)

	if typ := reflect.TypeOf(a).Kind(); typ != reflect.Map {
		t.Errorf("ToMap should return a map type, got: %v", typ)
	}

	// we have three fields
	if len(a) != 3 {
		t.Errorf("ToMap should return a map of len 3, got: %d", len(a))
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
			t.Errorf("ToMap should have the value %v", val)
		}
	}

}

func TestToMap_Tag(t *testing.T) {
	var T = struct {
		A string `structure:"x"`
		B int    `structure:"y"`
		C bool   `structure:"z"`
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	a := ToMap(T)

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
			t.Errorf("ToMap should have the key %v", key)
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

func TestToSlice(t *testing.T) {
	var T = struct {
		A string
		B int
		C bool
	}{
		A: "a-value",
		B: 2,
		C: true,
	}

	s := ToSlice(T)

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
			t.Errorf("ToSlice should have the value %v", val)
		}
	}
}
