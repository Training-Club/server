package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"regexp"
)

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

// ValidatePassword parses a string input and
// returns true if the provided string is a vlid
// password format.
func ValidatePassword(s string) bool {
	if len(s) < 6 || len(s) > 32 {
		return false
	}

	return true
}

// ValidateToken parses an encoded token (assumed to be a JWT signed by this service)
// and will return it as a converted jwt token object.
func ValidateToken(encoded string, pubkey string) (*jwt.Token, error) {
	return jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token format %v", token.Header["alg"])
		}
		return []byte(pubkey), nil
	})
}
