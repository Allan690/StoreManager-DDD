package entity

import "errors"

var ErrNotFound = errors.New("not found")
var ErrInvalidEntity = errors.New("invalid entity")
var ErrInvalidUserEmailOrPassword = errors.New("invalid email and/or password")
var ErrInvalidAccessToken = errors.New("invalid access token")
