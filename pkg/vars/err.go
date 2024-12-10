package vars

import "errors"

var (
	ErrAlreadyExist     = errors.New("already exist")
	ErrDBNotInitialized = errors.New("Database not initialized")
)
