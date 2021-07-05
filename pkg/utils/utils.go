package utils

import (
	"math"
)

// Number checker
type Number struct {
	epsilon float64
}

// NewNumber ...
func NewNumber(
	epsilon float64,
) *Number {
	return &Number{
		epsilon: epsilon,
	}
}

// IsFloatInt returns true if value is a whole number.
func (ts *Number) IsFloatInt(value float64) bool {
	_, frac := math.Modf(math.Abs(value))
	return frac < ts.epsilon || frac > 1.0-ts.epsilon
}
