package graph

import (
	"math/rand"
	"strings"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestGraph(t *testing.T) {

	t.Run("New Graph", func(t *testing.T) {
		vertices := 5
		directed_g, err := NewGraph(vertices, true)
		AssertNil(t, err)
		AssertEqual(t, directed_g.e, 0)
		AssertEqual(t, directed_g.v, 5)

		g, err := NewGraph(vertices, false)
		AssertNil(t, err)
		AssertEqual(t, g.e, 0)
		AssertEqual(t, g.v, 5)
	})

	t.Run("Add Edges", func(t *testing.T) {
		vertices := 5
		g_dir, _ := NewGraph(vertices, true)
		g_undir, _ := NewGraph(vertices, false)

		err := g_dir.AddEdge(0, 1)
		AssertNil(t, err)

		err = g_undir.AddEdge(0, 1)
		AssertNil(t, err)

		adj, _ := g_dir.Adj(0)
		adj2, _ := g_dir.Adj(1)
		AssertEqual(t, len(adj), 1)
		AssertEqual(t, adj[0], 1)
		AssertEqual(t, len(adj2), 0)
		AssertEqual(t, g_dir.e, 1)

		adj, _ = g_undir.Adj(0)
		adj2, _ = g_undir.Adj(1)
		AssertEqual(t, len(adj), 1)
		AssertEqual(t, adj[0], 1)
		AssertEqual(t, len(adj2), 1)
		AssertEqual(t, adj2[0], 0)
		AssertEqual(t, g_dir.e, 1)

	})

	t.Run("Remove Edges", func(t *testing.T) {
		vertices := 5
		g_dir, _ := NewGraph(vertices, true)
		g_undir, _ := NewGraph(vertices, false)

		g_dir.AddEdge(0, 1)
		err := g_dir.RemoveEdge(0, 1)
		AssertNil(t, err)

		g_undir.AddEdge(0, 1)
		err = g_undir.RemoveEdge(0, 1)
		AssertNil(t, err)

		adj, _ := g_dir.Adj(0)
		adj2, _ := g_dir.Adj(1)
		AssertEqual(t, len(adj), 0)
		AssertEqual(t, len(adj2), 0)
		AssertEqual(t, g_dir.e, 0)

		adj, _ = g_undir.Adj(0)
		adj2, _ = g_undir.Adj(1)
		AssertEqual(t, len(adj), 0)
		AssertEqual(t, len(adj2), 0)
		AssertEqual(t, g_dir.e, 0)
	})

	t.Run("ToString", func(t *testing.T) {
		g_undir, _ := NewGraph(4, false)
		_ = g_undir.AddEdge(0, 1)
		_ = g_undir.AddEdge(1, 2)
		_ = g_undir.AddEdge(2, 3)

		g_dir, _ := NewGraph(4, true)
		_ = g_dir.AddEdge(0, 1)
		_ = g_dir.AddEdge(1, 2)
		_ = g_dir.AddEdge(2, 3)

		graphString := g_undir.String()
		graphString2 := g_dir.String()

		expectedString := `4
3
0 1
1 0
1 2
2 1
2 3
3 2
`
		expectedString2 := `4
3
0 1
1 2
2 3
`
		AssertEqual(t, graphString, expectedString)
		AssertEqual(t, graphString2, expectedString2)
		if strings.TrimSpace(graphString) != strings.TrimSpace(expectedString) {
			t.Errorf("Expected graph string:\n%s\ngot:\n%s", expectedString, graphString)
		}
	})

	t.Run("NewGraphFromString", func(t *testing.T) {
		size := 2000
		g_undir := getGrapgh(size, false)
		g_dir := getGrapgh(size, true)

		graphString := g_undir.String()
		graphString2 := g_dir.String()

		g_undir2, _ := NewGraphFromString(graphString, false)
		g_dir2, _ := NewGraphFromString(graphString2, true)

		AssertTrue(t, g_dir.Equals(g_dir2))
		AssertTrue(t, g_undir2.Equals(g_undir2))
	})

	t.Run("Graph Copy", func(t *testing.T) {
		size := 2000
		g_undir := getGrapgh(size, false)
		g_dir := getGrapgh(size, true)

		g_undir2 := g_undir.CopyGraph()
		g_dir2 := g_dir.CopyGraph()

		AssertTrue(t, g_dir.Equals(g_dir2))
		AssertTrue(t, g_undir2.Equals(g_undir2))

	})

}

func getGrapgh(size int, directed bool) Graph {
	g, _ := NewGraph(size, directed)
	for i := 0; i < size; i++ {
		r := rand.Intn(size)
		if r == i {
			if i == 0 {
				i = 1
			} else {
				r = 0
			}
		}
		_ = g.AddEdge(i, r)
	}
	return *g
}
