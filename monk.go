// Package gomonk provides validation utilities and error handling for property-based validation.
package gomonk

// PropertyError represents a map of property names to error messages.
type PropertyError map[string]string

// ValidationError represents a map of field names to PropertyError objects.
type ValidationError map[string]PropertyError

// ValidatorFn is a function type that performs validation on a value of type T and adds any validation
// errors to the provided PropertyError.
type ValidatorFn[T any] func(v *T, e PropertyError)

// NewPropertyErrorStore creates a new empty PropertyError object.
func NewPropertyErrorStore() PropertyError {
	return make(PropertyError)
}
