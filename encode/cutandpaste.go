package encode

import (
	"errors"
	"fmt"
	"strings"
)

type CookieMap map[string]string

// GetProfile attempts to produce a profile
// from the CookieMap, returning an error if any
// required fields are missing from the cookie.
func (c CookieMap) GetProfile() (*Profile, error) {
	email, ok := c["email"]
	if !ok {
		return nil, errors.New("Could not find email")
	}
	uid, ok := c["uid"]
	if !ok {
		return nil, errors.New("Could not find uid")
	}
	role, ok := c["role"]
	if !ok {
		return nil, errors.New("Could not find role")
	}
	return &Profile{email, uid, role}, nil
}

// ParseCookie creates a CookieMap from a cookie string.
//
// Empty cookies produce empty maps. Keys with no assignment
// are treated as empty, as are keys with empty assignments.
// A second "=" in a value is accepted,
// so "x=y=z" becomes {"x": "y=z"}.
func ParseCookie(cookie string) *CookieMap {
	m := make(CookieMap)
	if len(cookie) == 0 {
		return &m
	}
	pairs := strings.Split(cookie, "&")
	for _, pair_string := range pairs {
		pair := strings.SplitN(pair_string, "=", 2)
		key := pair[0]
		if len(pair) == 1 {
			m[key] = ""
		} else if len(pair) == 2 {
			m[key] = pair[1]
		}
	}
	return &m
}

type Profile struct {
	Email string
	// Could be int but doesn't really matter
	UID  string
	Role string
}

// Cookie encodes the profile's data into a cookie.
func (p *Profile) Cookie() string {
	t := "email=%s&uid=%s&role=%s"
	return fmt.Sprintf(t, p.Email, p.UID, p.Role)
}

// ProfileFor provides suuuuuper secure email sanitization
// to ensure that emails don't have extra settings that will
// be injected into the cookie.
func ProfileFor(email string) *Profile {
	noAmp := strings.ReplaceAll(email, "&", "")
	sanitized := strings.ReplaceAll(noAmp, "=", "")
	return &Profile{sanitized, "10", "user"}
}
