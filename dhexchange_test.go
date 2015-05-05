package dhexchange

import (
	"testing"
)

func TestExchange(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := Randomkey()
		b := Randomkey()
		A := Exchange(a)
		B := Exchange(b)
		secretA := Secret(a, B)
		secretB := Secret(b, A)
		if secretA.Cmp(secretB) != 0 {
			t.Errorf("test[%d] failed\n", i)
		}
	}
}
