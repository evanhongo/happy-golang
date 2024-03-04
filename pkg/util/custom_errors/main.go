package custom_errors

import "errors"

var (
	ErrServerProblem    = errors.New("internal server error")
	ErrDuplicatedEmail  = errors.New("duplicated email not allowed")
	ErrEmailNotFound    = errors.New("email not found")
	ErrWrongPassword    = errors.New("wrong password")
	ErrNotAuthenticated = errors.New("not authenticated")
)
