package structure

import "strings"

// tagOptions contains a slice of tag options
type tagOptions []string

// Has returns true if the given opt is available inside the slice.
func (t tagOptions) Has(opt string) bool {
	for _, tagOpt := range t {
		if tagOpt == opt {
			return true
		}
	}

	return false
}

// parseTag splits a struct field's tag into its name and a list of options
// which comes after a name
func parseTag(tag string) (string, tagOptions) {
	res := strings.Split(tag, ",")

	// tag = ""
	if len(res) == 0 {
		return tag, res
	}

	// tag = "name"
	if len(res) == 1 {
		return tag, res[1:]
	}

	// tag is one of followings:
	// "name,opt"
	// "name,opt,opt2"
	// ",opt"
	return res[0], res[1:]
}
