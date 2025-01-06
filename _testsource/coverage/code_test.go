package example_test

import (
	"testing"

	"github.com/adombeck/example"
)

func TestHello(t *testing.T) {
	if example.Hello() != "Hello world!" {
		t.Fail()
	}
}
