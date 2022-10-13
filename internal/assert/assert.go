package assert

import (
	"fmt"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		fmt.Errorf("got %v; want %v", actual, expected)
	}
}
