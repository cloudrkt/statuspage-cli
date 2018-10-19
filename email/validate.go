package email

import (
	"errors"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateFormat for email address
func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("invalid emailadress: " + email)
	}
	return nil
}
