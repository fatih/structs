package structure

import (
	"fmt"
	"time"
)

func ExampleMap() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	s := &Server{
		Name:    "Arslan",
		ID:      123456,
		Enabled: true,
	}

	m := Map(s)

	fmt.Printf("%#v\n", m["Name"])
	fmt.Printf("%#v\n", m["ID"])
	fmt.Printf("%#v\n", m["Enabled"])
	// Output:
	// "Arslan"
	// 123456
	// true

}

func ExampleMap_tags() {
	// Custom tags can change the map keys instead of using the fields name
	type Server struct {
		Name    string `structure:"server_name"`
		ID      int32  `structure:"server_id"`
		Enabled bool   `structure:"enabled"`
	}

	s := &Server{
		Name: "Zeynep",
		ID:   789012,
	}

	m := Map(s)

	// access them by the custom tags defined above
	fmt.Printf("%#v\n", m["server_name"])
	fmt.Printf("%#v\n", m["server_id"])
	fmt.Printf("%#v\n", m["enabled"])
	// Output:
	// "Zeynep"
	// 789012
	// false

}

func ExampleValues() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	s := &Server{
		Name:    "Fatih",
		ID:      135790,
		Enabled: false,
	}

	m := Values(s)

	// note that the output is sorted according to the field names
	fmt.Printf("Values: %+v\n", m)
	// Output:
	// Values: [false 135790 Fatih]
}

func ExampleFields() {
	type Access struct {
		Name         string
		LastAccessed time.Time
		Number       int
	}

	s := &Access{
		Name:         "Fatih",
		LastAccessed: time.Now(),
		Number:       1234567,
	}

	m := Fields(s)

	// note that the output is sorted according to the field names
	fmt.Printf("Fields: %+v\n", m)
	// Output:
	// Fields: [LastAccessed Name Number]
}

func ExampleIsValid() {
	// Let's define an Access struct. Note that the "Enabled" field is not
	// going to be checked because we added the "structure" tag to the field.
	type Access struct {
		Name         string
		LastAccessed time.Time
		Number       int
		Enabled      bool `structure:"-"`
	}

	// Name and Number is not initialized.
	a := &Access{
		LastAccessed: time.Now(),
	}
	validA := IsValid(a)

	// Name and Number is initialized.
	b := &Access{
		Name:         "Fatih",
		LastAccessed: time.Now(),
		Number:       12345,
	}
	validB := IsValid(b)

	fmt.Printf("%#v\n", validA)
	fmt.Printf("%#v\n", validB)
	// Output:
	// false
	// true
}
