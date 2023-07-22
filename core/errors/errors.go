package errors

import "errors"

var (
	ErrNoRowsAffected          = errors.New("no rows affected")
	ErrNoRowsRetrieved         = errors.New("no rows retrieved")
	ErrUnexpectedNextResultSet = errors.New("unexpected next result set")
)
