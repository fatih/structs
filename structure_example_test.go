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

func ExampleMap_nested() {
	// By default field with struct types are processed too. We can stop
	// processing them via "omitnested" tag option.
	type Server struct {
		Name string    `structure:"server_name"`
		ID   int32     `structure:"server_id"`
		Time time.Time `structure:"time,omitnested"` // do not convert to map[string]interface{}
	}

	const shortForm = "2006-Jan-02"
	t, _ := time.Parse("2006-Jan-02", "2013-Feb-03")

	s := &Server{
		Name: "Zeynep",
		ID:   789012,
		Time: t,
	}

	m := Map(s)

	// access them by the custom tags defined above
	fmt.Printf("%v\n", m["server_name"])
	fmt.Printf("%v\n", m["server_id"])
	fmt.Printf("%v\n", m["time"].(time.Time))
	// Output:
	// Zeynep
	// 789012
	// 2013-02-03 00:00:00 +0000 UTC
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

	fmt.Printf("Values: %+v\n", m)
	// Output:
	// Values: [Fatih 135790 false]
}

func ExampleValues_tags() {
	type Location struct {
		City    string
		Country string
	}

	type Server struct {
		Name     string
		ID       int32
		Enabled  bool
		Location Location `structure:"-"` // values from location are not included anymore
	}

	s := &Server{
		Name:     "Fatih",
		ID:       135790,
		Enabled:  false,
		Location: Location{City: "Ankara", Country: "Turkey"},
	}

	// Let get all values from the struct s. Note that we don't include values
	// from the Location field
	m := Values(s)

	fmt.Printf("Values: %+v\n", m)
	// Output:
	// Values: [Fatih 135790 false]
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

	fmt.Printf("Fields: %+v\n", m)
	// Output:
	// Fields: [Name LastAccessed Number]
}

func ExampleFields_nested() {
	type Person struct {
		Name   string
		Number int
	}

	type Access struct {
		Person        Person `structure:",omitnested"`
		HasPermission bool
		LastAccessed  time.Time
	}

	s := &Access{
		Person:        Person{Name: "fatih", Number: 1234567},
		LastAccessed:  time.Now(),
		HasPermission: true,
	}

	// Let's get all fields from the struct s. Note that we don't include the
	// fields from the Person field anymore due to "omitnested" tag option.
	m := Fields(s)

	fmt.Printf("Fields: %+v\n", m)
	// Output:
	// Fields: [Person HasPermission LastAccessed]
}

func ExampleIsZero() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	// Nothing is initalized
	a := &Server{}
	isZeroA := IsZero(a)

	// Name and Enabled is initialized, but not ID
	b := &Server{
		Name:    "Golang",
		Enabled: true,
	}
	isZeroB := IsZero(b)

	fmt.Printf("%#v\n", isZeroA)
	fmt.Printf("%#v\n", isZeroB)
	// Output:
	// true
	// false
}

func ExampleHasZero() {
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
	hasZeroA := HasZero(a)

	// Name and Number is initialized.
	b := &Access{
		Name:         "Fatih",
		LastAccessed: time.Now(),
		Number:       12345,
	}
	hasZeroB := HasZero(b)

	fmt.Printf("%#v\n", hasZeroA)
	fmt.Printf("%#v\n", hasZeroB)
	// Output:
	// true
	// false
}

func ExampleHas() {
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

	found := Has(s, "LastAccessed")

	fmt.Printf("Has: %+v\n", found)
	// Output:
	// Has: true
}
