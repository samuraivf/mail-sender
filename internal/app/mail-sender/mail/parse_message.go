package mail

import (
	"errors"
	"strings"
)

var (
	errInvalidEmail = errors.New("error invalid email")
)

type message struct {
	email string
	link  string
}

func parseMessage(msg string) (message, error) {
	parts := strings.Split(msg, ":")
	email := parts[0]
	link := parts[1]

	if ok := validateEmail(email); !ok {
		return message{}, errInvalidEmail
	}

	return message{email: email, link: link}, nil
}
