package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionVyZero = errors.New("division by zero")
	ErrInvalidSymbol = errors.New("invalid symbols")
	ErrUnknownError = errors.New("unknown error")
	ErrEmptyExpression = errors.New("empty expression")
)