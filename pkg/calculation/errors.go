package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero = errors.New("division by zero")
	ErrInvalidSymbol = errors.New("invalid symbols")
	ErrEmptyExpression = errors.New("empty expression")
)