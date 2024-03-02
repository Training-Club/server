package util

import "regexp"

// ValidateUsername parses a string input and
// returns true if the provided string is a valid
// username format.
func ValidateUsername(s string) bool {
	if len(s) == 0 {
		return false
	}

	rexp, err := regexp.Compile(`^[a-zA-Z0-9._]+$`)
	if err != nil {
		return false
	}

	first := string(s[0])
	if first == "." || first == "_" {
		return false
	}

	return rexp.MatchString(s)
}

// ValidateEmail parses a string input and
// returns true if the provided string is a valid
// email format.
func ValidateEmail(s string) bool {
	rexp, err := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if err != nil {
		return false
	}

	return rexp.MatchString(s)
}
