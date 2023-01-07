package dmens

import "errors"

var (
	ErrUserNotRegistered    = errors.New("the user not registered")
	ErrGetNullConfiguration = errors.New("get remote configuration with null")
)
