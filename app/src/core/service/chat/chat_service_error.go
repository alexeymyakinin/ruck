package chat

import "errors"

var (
	ErrIDNotFound = errors.New("cannot find chat with specified id")
)
