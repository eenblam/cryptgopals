package encode

import (
	"reflect"
	"testing"
)

func TestParseCookie(t *testing.T) {
	cases := []struct {
		Name     string
		Cookie   string
		Expected *CookieMap
	}{
		{
			"Empty cookie is empty map",
			"",
			&CookieMap{},
		},
		{
			"Simple happy path",
			"email=foo@bar.com&uid=10&role=user&nokey&emptyvalue=",
			&CookieMap{
				"email":      "foo@bar.com",
				"uid":        "10",
				"role":       "user",
				"nokey":      "",
				"emptyvalue": "",
			},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := ParseCookie(test.Cookie)
			if !reflect.DeepEqual(*got, *test.Expected) {
				t.Error("Didn't get the map we expected")
			}
		})
	}
}

func TestGetProfile(t *testing.T) {
	cases := []struct {
		Name     string
		Cookie   string
		Expected *Profile
		Error    bool
	}{
		{
			"A good cookie produces a good Profile",
			"email=foo@bar.com&uid=10&role=user",
			&Profile{"foo@bar.com", "10", "user"},
			false,
		},
		{
			"No email, no profile",
			"emailll=foo@bar.com&uid=10&role=user",
			nil,
			true,
		},
		{
			"No uid, no profile",
			"email=foo@bar.com&gid=10&role=user",
			nil,
			true,
		},
		{
			"No role, no profile",
			"email=foo@bar.com&uid=10&roll=user",
			nil,
			true,
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			m := ParseCookie(test.Cookie)
			profile, err := m.GetProfile()
			if err == nil && test.Error {
				t.Error("Expected error, got none")
			}
			if err != nil && !test.Error {
				t.Errorf("Unexpected error: %s", err)
			}
			if !test.Error && !reflect.DeepEqual(*profile, *test.Expected) {
				t.Error("Generated profiles do not match")
			}
		})
	}
}

func TestProfileFor(t *testing.T) {
	cases := []struct {
		Name     string
		Email    string
		Expected *Profile
	}{
		{
			"A good email provides a good profile",
			"foo@bar.com",
			&Profile{"foo@bar.com", "10", "user"},
		},
		{
			"A bad email is properly escaped",
			"foo@bar.com&role=admin",
			&Profile{"foo@bar.comroleadmin", "10", "user"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := ProfileFor(test.Email)
			if !reflect.DeepEqual(*got, *test.Expected) {
				t.Error("Didn't get the Profile we expected")
			}
		})
	}
}

func TestCookie(t *testing.T) {
	// Coveraaaaage
	p := Profile{"foo@bar.com", "10", "user"}
	got := p.Cookie()
	expected := "email=foo@bar.com&uid=10&role=user"
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
