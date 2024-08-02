// Package math provides an interface for mathematical equations
package math

import "golang.org/x/exp/constraints"

// A Number is any signed or unsigned integer or float
type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two Numbers and returns the sum
//
// For more information about [addition]
//
// [addition]: https://mathisfun.com/numbers/addition.html
func Add[T Number](a T, b T) T {
	return a + b
}