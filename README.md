# Structure [![GoDoc](https://godoc.org/github.com/fatih/structure?status.png)](http://godoc.org/github.com/fatih/structure) [![Build Status](https://travis-ci.org/fatih/structure.png)](https://travis-ci.org/fatih/structure)

Structure contains various utilitis to work with Go structs.

## Install

```bash
go get github.com/fatih/structure
```

## Examples

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

m, err := structure.ToMap(s)
if err != nil {
	panic(err)
}

fmt.Printf("%#v", m)
// Output: map[string]interface {}{"Name":"Arslan", "ID":123456, "Enabled":true}
```
	
