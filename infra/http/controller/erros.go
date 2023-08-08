package controller

import (
	"errors"
)

var ErrWithoutEmail = errors.New("email is required")
var ErrWithoutName = errors.New("name is required")
