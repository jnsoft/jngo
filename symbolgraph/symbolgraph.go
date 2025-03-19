package symbolgraph

import (
	"fmt"
	"strings"

	"github.com/jnsoft/jngo/graph"
	"github.com/jnsoft/jngo/red_black_bst"
	"github.com/jnsoft/jngo/stringhelper"
)

type (
	SymbolGraph struct {
		st   *red_black_bst.RedBlackTree[string, int]
		keys []string
		g    *graph.Graph
	}
)

// builds graph from strings instead of integers for vertices (wrapper for regular graph)
// example input: MovieName1/Actor1/Actor2/Actor3\nMovieName2/Actor4/Actor5/Actor2\n... (delimeter = '/')
func NewSymbolGraph(input, delimeter string, is_directed bool) (*SymbolGraph, error) {
	// First pass builds the symbol table by reading strings to associate distinct strings with an index
	st := red_black_bst.NewRedBlackTree[string, int]()
	lines := stringhelper.ToLines(input)
	for _, line := range lines {
		a := strings.Split(line, delimeter)
		for i := 0; i < len(a); i++ {
			if !st.Contains(a[i]) {
				st.Put(a[i], st.Size())
			}
		}
	}

	// inverted index to get string keys in an aray
	keys := make([]string, st.Size())
	for _, name := range st.Keys() {
		ix, _ := st.Get(name)
		keys[ix] = name
	}

	// second pass builds the graph by connecting first vertex on each line to all others
	g, _ := graph.NewGraph(st.Size(), is_directed)
	for _, line := range lines {
		a := strings.Split(line, delimeter)
		v, _ := st.Get(a[0])
		for i := 1; i < len(a); i++ {
			w, _ := st.Get(a[i])
			g.AddEdge(v, w)
		}
	}
	return &SymbolGraph{
		st:   st,
		keys: keys,
		g:    g,
	}, nil
}

func (sg *SymbolGraph) Contains(str string) bool {
	return sg.st.Contains(str)
}

func (sg *SymbolGraph) Index(str string) int {
	ix, _ := sg.st.Get(str)
	return ix
}

func (sg *SymbolGraph) Name(v int) string {
	return sg.keys[v]
}

// DegreesOfSeparation / 2 = "Kevin Bacon index"
func (sg *SymbolGraph) DegreesOfSeparation(source, sink string) (int, string) {
	var sb strings.Builder
	ret := -1
	if sg.g.IsDirected() {
		panic("can't calculate degrees of separation for directed graph")
	}
	if !sg.Contains(source) {
		sb.WriteString(fmt.Sprintf("Symbol graph does not contain %v", source))
	}
	if !sg.Contains(sink) {
		sb.WriteString(fmt.Sprintf("Symbol graph does not contain %v", sink))
	}

	ret = 0
	paths := sg.g.BreadthFirstPaths(sg.Index(source))
	t := sg.Index(sink)
	if paths.HasPathTo(t) {
		for _, v := range paths.PathTo(t) {
			sb.WriteString(fmt.Sprintf("%v", sg.Name(v)))
			ret++
		}
	} else {
		sb.WriteString("Not connected")
	}

	output := sb.String()
	return ret, output
}
