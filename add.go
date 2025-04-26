// publicmodule is a package containing some utilities functions
package publicmodule

import "golang.org/x/exp/constraints"

// Number is an interface which defines type contraints
type Number interface {
	constraints.Float | constraints.Integer
}

// Add is generic function which adds two floats or integers
func Add[T Number](a, b T) T {
	return a + b
}
