# Structure [![GoDoc](https://godoc.org/github.com/fatih/structure?status.svg)](http://godoc.org/github.com/fatih/structure) [![Build Status](https://travis-ci.org/fatih/structure.svg)](https://travis-ci.org/fatih/structure)

Structure contains various utilitis to work with Go (Golang) structs.

## Install

```bash
go get github.com/fatih/structure
```

## Examples

```go
// Lets define and declare a struct
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

```go
// convert it to a map[string]interface{}
m, err := structure.ToMap(s)
if err != nil {
	panic(err)
}

fmt.Printf("%#v", m)
// prints: map[string]interface {}{"Name":"Arslan", "ID":123456, "Enabled":true}
```

```go
// check if it's a struct or a pointer to struct
if structure.IsStruct(s) {
    fmt.Println("s is a struct") 
}
```
	
