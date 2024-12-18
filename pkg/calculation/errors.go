package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionVyZero = errors.New("division by zero")
)