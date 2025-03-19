package graph

import (
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

}
