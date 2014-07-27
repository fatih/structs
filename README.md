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

Convert it to a map[string]interface{}

```go
m, err := structure.ToMap(s)
if err != nil {
	panic(err)
}

// prints: map[string]interface {}{"Name":"Arslan", "ID":123456, "Enabled":true}
fmt.Printf("%#v", m)
```

Convert to a `[]interface{}`. Slice values are **sorted** by default according
to the field names.

```go
m, err := structure.ToSlice(s)
if err != nil {
	panic(err)
}

// prints: []interface {}{true, 123456, "Arslan"}
fmt.Printf("%#v", m)
```

Check if it's a struct or a pointer to struct

```go
if structure.IsStruct(s) {
    fmt.Println("s is a struct") 
}
```
	
