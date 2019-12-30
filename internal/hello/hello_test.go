package hello

import "testing"

func Test_hello(t *testing.T) {
	actual := Hello()
	if actual != "Hello world \n" {
		t.Fatalf("expected: %s, but was: %s", "Hello", actual)
	}
}

