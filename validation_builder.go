package gomonk

// MonkValidatorBuilder is a builder for constructing a MonkValidator.
type MonkValidatorBuilder[T any] struct {
	validators map[string][]ValidatorFn[T]
}

// NewRuleBook creates a new MonkValidatorBuilder instance.
func NewRuleBook[T any]() *MonkValidatorBuilder[T] {
	return &MonkValidatorBuilder[T]{
		validators: make(map[string][]ValidatorFn[T]),
	}
}

// AddRuleFor adds a validator function for the specified field.
func (m *MonkValidatorBuilder[T]) AddRuleFor(field string, validator ValidatorFn[T]) {
	m.validators[field] = append(m.validators[field], validator)
}

// Build constructs a MonkValidator with the registered validators.
func (m *MonkValidatorBuilder[T]) Build() *MonkValidator[T] {
	return NewMonkValidator(m.validators)
}
