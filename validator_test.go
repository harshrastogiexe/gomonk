package gomonk

import (
	"fmt"
	"testing"
)

func TestNewMonkValidator(t *testing.T) {
	v := NewMonkValidator[string](make(map[string][]ValidatorFn[string]))

	t.Run("created a nil value", func(t *testing.T) {
		if v != nil {
			return
		}
		t.Error("created a nil value")
	})
	t.Run("create empty property store", func(t *testing.T) {
		if exp, got := 0, len(v.pe); exp != got {
			t.Errorf("expected to be of size %d, got %d", exp, got)
		}
	})
}

func TestValidate(t *testing.T) {
	type User struct {
		Name  string
		Age   int
		Email string
	}

	rb := NewRuleBook[User]()
	rb.AddRuleFor("age", func(v *User, e PropertyError) {
		min := 18
		if v.Age < min {
			e["value"] = fmt.Sprintf("age should be more than %d", min)
		}
	})

	v := rb.Build()

	t.Run("should return validation error", func(t *testing.T) {
		tu := &User{
			Name:  "Harsh Rastogi",
			Age:   18,
			Email: "harsh@email.com",
		}

		e := v.Validate(tu)

		if e == nil {
			t.Error("got nil error")
		}

		if _, ok := e["age"]["value"]; !ok && tu.Age < 18 {
			t.Error("expected age error to be present")
		}

		if _, ok := e["age"]["value"]; ok && tu.Age >= 18 {
			t.Error("expected age error not to be present")
		}
	})
}
