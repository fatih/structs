# Structure [![GoDoc](https://godoc.org/github.com/fatih/structure?status.svg)](http://godoc.org/github.com/fatih/structure) [![Build Status](https://travis-ci.org/fatih/structure.svg)](https://travis-ci.org/fatih/structure)

Structure contains various utilitis to work with Go (Golang) structs.

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
	Name:    "Arslan",
	ID:      123456,
	Enabled: true,
}
```

#### Map()

Convert a struct to a `map[string]interface{}`

```go
m := structure.Map(s)

// prints: map[string]interface {}{"Name":"Arslan", "ID":123456, "Enabled":true}
fmt.Printf("%#v", m)
```

#### ToSlice()

Convert the values of a struct to a `[]interface{}`. Slice values are
**sorted** by default according to the field names.

```go
m := structure.ToSlice(s)

// prints: []interface {}{true, 123456, "Arslan"}
fmt.Printf("%#v", m)
```

#### Fields()

Convert the fields of a struct to a `[]string`. Slice values are **sorted** by
default according to the field names.

```go
m := structure.Fields(s)

// prints: []string{"Enabled", "ID", "Name"}
fmt.Printf("%#v", m)
```

#### IsStruct()

Check if it's a struct or a pointer to struct

```go
if structure.IsStruct(s) {
    fmt.Println("s is a struct") 
}
```
	
