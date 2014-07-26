package structure

import "fmt"

func ExampleToMap() {
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

	m, err := ToMap(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", m["Name"])
	fmt.Printf("%#v\n", m["ID"])
	fmt.Printf("%#v\n", m["Enabled"])
	// Output:
	// "Arslan"
	// 123456
	// true

}

func ExampleToMap_tags() {
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

	m, err := ToMap(s)
	if err != nil {
		panic(err)
	}

	// access them by the custom tags defined above
	fmt.Printf("%#v\n", m["server_name"])
	fmt.Printf("%#v\n", m["server_id"])
	fmt.Printf("%#v\n", m["enabled"])
	// Output:
	// "Zeynep"
	// 789012
	// false

}
