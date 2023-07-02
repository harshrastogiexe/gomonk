package gomonk

// MonkValidator represents a validator that performs property-based validation.
type MonkValidator[T any] struct {
	validators map[string][]ValidatorFn[T]
	pe         map[string]PropertyError
}

// NewMonkValidator creates a new instance of MonkValidator with the provided validators.
func NewMonkValidator[T any](validators map[string][]ValidatorFn[T]) *MonkValidator[T] {
	return &MonkValidator[T]{
		validators: validators,
		pe:         make(map[string]PropertyError),
	}
}

// Validate performs validation on the given value using the registered validators.
// It returns a ValidationError object containing any validation errors found.
func (m *MonkValidator[T]) Validate(value *T) ValidationError {
	ve := make(ValidationError)
	for key, vFns := range m.validators {
		m.pe[key] = make(PropertyError)
		for _, fn := range vFns {
			fn(value, m.pe[key])
		}
		ve[key] = m.pe[key]
	}
	return ve
}
