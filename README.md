# Structure [![GoDoc](https://godoc.org/github.com/fatih/structure?status.svg)](http://godoc.org/github.com/fatih/structure) [![Build Status](https://travis-ci.org/fatih/structure.svg)](https://travis-ci.org/fatih/structure)

Structure contains various utilities to work with Go (Golang) structs.

**WIP, Use with care until it's finished and announced**

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
// => [true, 123456, "gopher"]
v := structure.Values(s)

// Convert the fields of a struct to a []string. 
// => ["Enabled", "ID", "Name"]
f := structure.Fields(s)

// Return the struct name
// => "Server"
n := structure.Name(s)

// Check if the fields of a struct is initialized or not.
if structure.IsValid(s) {
    fmt.Println("s is initialized")
}

// Check if it's a struct or a pointer to struct
if structure.IsStruct(s) {
    fmt.Println("s is a struct")
}

```

