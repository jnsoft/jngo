package graph

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jnsoft/jngo/misc"
	"github.com/jnsoft/jngo/stack"
	"github.com/jnsoft/jngo/stringhelper"
)

type (
	Graph struct {
		v             int // Number of vertices
		e             int // Number of edges
		adjacencyList [][]int
		isDirected    bool
	}
)

func (g *Graph) Adj(v int) ([]int, error) {
	if v < 0 || v >= g.v {
		return nil, errors.New("vertex out of bounds")
	}
	return g.adjacencyList[v], nil
}

func readEdge(e string) (v, w int, err error) {
	// Replace multiple spaces with a single space
	re := regexp.MustCompile(`\s+`)
	e = re.ReplaceAllString(e, " ")

	e = strings.TrimSpace(e)

	vs := strings.Split(e, " ")
	if len(vs) < 2 {
		return 0, 0, fmt.Errorf("invalid edge format")
	}

	v, err1 := strconv.Atoi(vs[0])
	w, err2 := strconv.Atoi(vs[1])
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("failed to parse vertices: %s", e)
	}

	return v, w, nil
}

func NewGraph(v int, is_directed bool) (*Graph, error) {
	if v < 0 {
		return nil, errors.New("number of vertices must be nonnegative")
	}
	g := &Graph{
		v:             v,
		e:             0,
		adjacencyList: make([][]int, v),
		isDirected:    is_directed,
	}
	for i := 0; i < v; i++ {
		g.adjacencyList[i] = []int{}
	}
	return g, nil
}

func NewGraphFromString(g string, is_directed bool) (*Graph, error) {
	lines := stringhelper.ToLines(g)
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid input: no vertex count specified")
	}
	v, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse vertex count: %v", err)
	}

	e, err := strconv.Atoi(lines[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse edge count: %v", err)
	}

	graph, err := NewGraph(v, is_directed)
	if err != nil {
		return nil, err
	}

	graph.e = e

	for i := 2; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" {
			continue
		}

		v1, w1, err := readEdge(lines[i])
		if err != nil {
			return nil, fmt.Errorf("failed to parse edge: %v", err)
		}
		err = graph.addEdgeRaw(v1, w1)
		e--
		if err != nil {
			return nil, fmt.Errorf("failed to add edge: %v", err)
		}
	}
	if graph.isDirected && e != 0 {
		return nil, fmt.Errorf("wrong number of edges, want %v, got %v", graph.e, graph.e-e)
	}
	if !graph.isDirected && e != -graph.e {
		return nil, fmt.Errorf("wrong number of edges, want %v, got %v", graph.e, graph.e-e/2)
	}
	return graph, nil
}

func (g *Graph) AddEdge(v, w int) error {
	if v < 0 || v >= g.v || w < 0 || w >= g.v {
		return errors.New("vertex out of bounds")
	}
	g.e++
	g.adjacencyList[v] = append(g.adjacencyList[v], w)
	if !g.isDirected {
		g.adjacencyList[w] = append(g.adjacencyList[w], v)
	}
	return nil
}

func (g *Graph) addEdgeRaw(v, w int) error {
	if v < 0 || v >= g.v || w < 0 || w >= g.v {
		return errors.New("vertex out of bounds")
	}
	g.adjacencyList[v] = append(g.adjacencyList[v], w)
	return nil
}

func (g *Graph) RemoveEdge(v, w int) error {
	if g.e == 0 {
		return errors.New("no edges to remove in graph")
	}
	if v < 0 || v >= g.v || w < 0 || w >= g.v {
		return fmt.Errorf("vertex out of bounds: %d or %d", v, w)
	}
	adjV := g.adjacencyList[v]
	found := false
	for i, value := range adjV {
		if value == w {
			g.adjacencyList[v] = append(adjV[:i], adjV[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("no edge in graph between %d and %d", v, w)
	}
	if !g.isDirected {
		for i, value := range g.adjacencyList[w] {
			if value == v {
				g.adjacencyList[w] = append(g.adjacencyList[w][:i], g.adjacencyList[w][i+1:]...)
				break
			}
		}
	}
	g.e--
	return nil
}

func (g *Graph) CopyGraph() *Graph {
	cp, _ := NewGraph(g.v, g.isDirected)
	cp.e = g.e
	for v := 0; v < g.v; v++ {
		reverse := stack.New[int]()
		for _, w := range g.adjacencyList[v] {
			reverse.Push(w)
		}
		for !reverse.IsEmpty() {
			cp.adjacencyList[v] = append(cp.adjacencyList[v], reverse.Pop())
		}
	}
	return cp
}

func (g *Graph) Equals(other *Graph) bool {
	if g.v != other.v || g.e != other.e || g.isDirected != other.isDirected {
		return false
	}

	if len(g.adjacencyList) != len(other.adjacencyList) {
		return false
	}

	for i := range g.adjacencyList {
		if !misc.EqualSlices[int](g.adjacencyList[i], other.adjacencyList[i]) {
			return false
		}
	}

	return true
}

func (g *Graph) ToString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d\n%d\n", g.v, g.e))
	for v := 0; v < g.v; v++ {
		for _, w := range g.adjacencyList[v] {
			sb.WriteString(fmt.Sprintf("%d %d\n", v, w))
		}
	}
	return sb.String()
}
