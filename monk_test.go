package gomonk_test

import (
	"gomonk"
	"testing"
)

func TestNewPropertyErrorStore(t *testing.T) {
	e := gomonk.NewPropertyErrorStore()

	t.Run("created a nil value", func(t *testing.T) {
		if e != nil {
			return
		}
		t.Error("created a nil value")
	})

	t.Run("create empty store", func(t *testing.T) {
		if exp, got := 0, len(e); exp != got {
			t.Errorf("expected to be of size %d, got %d", exp, got)
		}
	})
}
