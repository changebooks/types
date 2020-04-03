package strs

import "strings"

// Uppercase the first character of the word in a native string
func UpperFirst(s string, otherLower bool) string {
	for _, c := range s {
		first := string(c)
		other := s[len(first):]
		if otherLower {
			other = strings.ToLower(other)
		}

		return strings.ToUpper(first) + other
	}

	return ""
}

// Lowercase the first character of the word in a native string
func LowerFirst(s string, otherUpper bool) string {
	for _, c := range s {
		first := string(c)
		other := s[len(first):]
		if otherUpper {
			other = strings.ToUpper(other)
		}

		return strings.ToLower(first) + other
	}

	return ""
}
