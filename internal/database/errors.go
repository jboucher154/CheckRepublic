package database

import "errors"

var (
    ErrNotFound = errors.New("not found")
    ErrConflict = errors.New("conflict") //e.g. for duplicate entries etc
    ErrIncorrectRequest = errors.New("incorrect request")
)
