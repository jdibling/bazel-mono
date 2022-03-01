package greetings_test

import (
	"testing"

	"github.com/go-playground/assert"
	"github.com/jdibling/bazel-mono/src/greetings"
)

func TestGreeting(t *testing.T) {
	assert.NotEqual(t, greetings.Greeting(), "")
}
