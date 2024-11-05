package helpers

// BoolToInt converts a bool to an int.
//
// It returns "1" if b == true and "0" otherwise.
func BoolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
