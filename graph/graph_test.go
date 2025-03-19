package graph

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestGraph(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		s := NewGraph(3)

		// check stack is empty
		AssertTrue(t, s.IsEmpty())