package strutils

import (
	"unicode"
)

// ToSnakeCase converts a string to snake case.
func ToSnakeCase(s string) string {
	// rs is the string as an array of runes.
	rs := []rune(s)

	// res will be the result of the conversion.
	res := []rune{}

	for _, r := range rs {
		switch {
		case unicode.IsLetter(r):
			if unicode.IsLower(r) {
				res = append(res, r)
			} else {
				// append an underscore if the rune is uppercase
				// only if the previous added rune is not an underscore.
				if len(res) > 0 && res[len(res)-1] != '_' {
					res = append(res, '_')
				}

				// append the lowercase rune
				res = append(res, unicode.ToLower(r))
			}

		// if a digit is found, append an underscore only if
		// the previous rune is not a digit.
		case unicode.IsDigit(r):

			if len(res) > 0 && res[len(res)-1] != '_' && !unicode.IsDigit(res[len(res)-1]) {
				res = append(res, '_')
			}

			res = append(res, r)
		// if an special rune is found, append an underscore only if
		// the previous rune is not an underscore.
		default:
			// if no previous rune was added, to the result
			// we add an underscore and continue iterating.
			if len(res) == 0 {
				res = append(res, '_')
				continue
			}

			// checks if the previous rune is an underscore
			// and if it is not, we append an underscore.
			if len(res) > 0 && res[len(res)-1] != '_' {
				res = append(res, '_')
			}
		}
	}

	return string(res)
}
