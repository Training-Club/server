package tests

import (
	"tc-server/util"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	cases := []struct {
		username string
		valid    bool
	}{
		{"user.name", true},
		{"user_name123", true},
		{"User.Name_123", true},
		{"invalid username", false},
		{"user@name", false},
		{"", false},
		{".username", false},
		{"_username", false},
	}

	for _, c := range cases {
		result := util.ValidateUsername(c.username)
		if result != c.valid {
			t.Errorf("ValidateUsername(%q) == %v, want %v", c.username, result, c.valid)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	cases := []struct {
		email string
		valid bool
	}{
		{"email@example.com", true},
		{"firstname.lastname@example.com", true},
		{"email@subdomain.example.com", true},
		{"email@123.123.123.123", true},
		{"email@[123.123.123.123]", false}, // This pattern is technically valid but not covered by our regex
		{"plainaddress", false},
		{"@no-local-part.com", false},
		{"Outlook Contact <outlook.contact@domain.com>", false},
		{"", false},
	}

	for _, c := range cases {
		result := util.ValidateEmail(c.email)
		if result != c.valid {
			t.Errorf("ValidateEmail(%q) == %v, want %v", c.email, result, c.valid)
		}
	}
}
