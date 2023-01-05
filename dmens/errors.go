package dmens

import "errors"

var (
	ErrUserNotRegistered    = errors.New("The user not registered")
	ErrGetNullConfiguration = errors.New("get remote configuration with null")
)
