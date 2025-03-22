package symbolgraph

import (
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestSymbolGraph(t *testing.T) {

	t.Run("New SymbolGraph", func(t *testing.T) {
		input := `MovieName1/Actor1/Actor2/Actor3
MovieName2/Actor4/Actor5/Actor2
MovieName3/Actor1/Actor2/Actor3
MovieName4/Actor4/Actor5/Actor6
`
		sg, err := NewSymbolGraph(input, "/", false)
		AssertNil(t, err)

		res, _ := sg.DegreesOfSeparation("Actor1", "Actor6")

		AssertEqual(t, res, 7)
		AssertEqual(t, res/2, 3)

	})
}
