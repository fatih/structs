# Structure [![GoDoc](https://godoc.org/github.com/fatih/structure?status.svg)](http://godoc.org/github.com/fatih/structure) [![Build Status](https://travis-ci.org/fatih/structure.svg)](https://travis-ci.org/fatih/structure) [![Coverage Status](https://img.shields.io/coveralls/fatih/structure.svg)](https://coveralls.io/r/fatih/structure)

Structure contains various utilities to work with Go (Golang) structs. It was
initially used by me to convert a struct into a `map[string]interface{}`. With
time I've added other utilities for structs.  It's basically a high level
package based on primitives from the reflect package. Feel free to add new
functions or improve the existing code.

## Install

```bash
go get github.com/fatih/structure
```

## Usage and Examples

Lets define and declare a struct

```go
type Server struct {
	Name        string `json:"name,omitempty"`
	ID          int
	Enabled     bool
	users       []string // not exported
	http.Server          // embedded
}

server := &Server{
	Name:    "gopher",
	ID:      123456,
	Enabled: true,
}
```

### Struct methods

Let's create a new `Struct` type. 

```go
// Create a new struct type:
s := structure.New(server)

// Convert a struct to a map[string]interface{}
// => {"Name":"gopher", "ID":123456, "Enabled":true}
m := s.Map()

// Convert the values of a struct to a []interface{}
// => ["gopher", 123456, true]
v := s.Values()

// Convert the values of a struct to a []*Field
// (see "Field methods" for more info about fields)
f := s.Fields()

// Check if any field of a struct is initialized or not.
if s.HasZero() {
    fmt.Println("s has a zero value field")
}

// Check if all fields of a struct is initialized or not.
if s.IsZero() {
    fmt.Println("all fields of s is zero value")
}

// Return the struct name
// => "Server"
n := s.Name()
```

Most of the struct methods are available as global functions without the need
for a `New()` constructor:

```go
m := structure.Map(s)      // Get a map[string]interface{}
v := structure.Values(s)   // Get a []interface{}
f := structure.Fields(s)   // Get a []*Field
n := structure.Name(s)     // Get the struct name
h := structure.HasZero(s)  // Check if any field is initialized
z := structure.IsZero(s)   // Check if all fields are initialized
i := structure.IsStruct(s) // Check if s is a struct or a pointer to struct
```

### Field methods

We can easily examine a single Field for more detail. Below you can see how we
get and interact with various field methods:


```go
s := structure.New(server)

// Get the Field struct for the "Name" field
name := s.Field("Name")

// Get the underlying value,  value => "gopher"
value := name.Value().(string)

// Check if the field is exported or not
if name.IsExported() {
	fmt.Println("Name field is exported")
}

// Check if the value is a zero value, such as "" for string, 0 for int
if !name.IsZero() {
	fmt.Println("Name is initialized")
}

// Check if the field is an anonymous (embedded) field
if !name.IsEmbedded() {
	fmt.Println("Name is not an embedded field")
}

// Get the Field's tag value for tag name "json", tag value => "name,omitempty"
tagValue := name.Tag("json")
```

Nested structs are supported too:

```go
addrField := s.Field("Server").Field("Addr")

// Get the value for addr
a := addrField.Value().(string)
```

We can also get a slice of Fields from the Struct type to iterate over all
fields. This is handy if you whish to examine all fields:

```go
// Convert the fields of a struct to a []*Field
fields := s.Fields()

for _, f := range fields {
	fmt.Printf("field name: %+v\n", f.Name())

	if f.IsExported() {
		fmt.Printf("value   : %+v\n", f.Value())
		fmt.Printf("is zero : %+v\n", f.IsZero())
	}
}
```

## Credits

 * [Fatih Arslan](https://github.com/fatih)
 * [Cihangir Savas](https://github.com/cihangir)

## License

The MIT License (MIT) - see LICENSE.md for more details


