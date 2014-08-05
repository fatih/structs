# Structure [![GoDoc](https://godoc.org/github.com/fatih/structure?status.svg)](http://godoc.org/github.com/fatih/structure) [![Build Status](https://travis-ci.org/fatih/structure.svg)](https://travis-ci.org/fatih/structure)

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
	Name    string
	ID      int32
	Enabled bool
}

s := &Server{
	Name:    "gopher",
	ID:      123456,
	Enabled: true,
}
```

```go
// Convert a struct to a map[string]interface{}
// => {"Name":"gopher", "ID":123456, "Enabled":true}
m := structure.Map(s)

// Convert the values of a struct to a []interface{}
// => [123456, "gopher", true]
v := structure.Values(s)

// Convert the fields of a struct to a []string. 
// => ["Name", "ID", "Enabled"]
f := structure.Fields(s)

// Return the struct name
// => "Server"
n := structure.Name(s)

// Check if field name exists
// => true
h := structure.Has(s, "Enabled")

// Check if a field of a struct is initialized or not.
if structure.HasZero(s) {
    fmt.Println("s has a zero value field")
}

// Check if all field of a struct is initialized or not.
if structure.IsZero(s) {
    fmt.Println("all fields of s is zero value")
}

// Check if it's a struct or a pointer to struct
if structure.IsStruct(s) {
    fmt.Println("s is a struct")
}

```

## Credits

 * [Fatih Arslan](https://github.com/fatih)
 * [Cihangir Savas](https://github.com/cihangir)

## License

The MIT License (MIT) - see LICENSE.md for more details


