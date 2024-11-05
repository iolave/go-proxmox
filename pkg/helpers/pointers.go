package helpers

// NewInt returns an int pointer of the passed
// value.
func NewInt(v int) *int {
	return &v
}

// NewStr returns a string pointer of the passed
// value.
func NewStr(v string) *string {
	return &v
}

// NewBool returns bool pointer of the passed
// value.
func NewBool(v bool) *bool {
	return &v
}
