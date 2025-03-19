package graph

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jnsoft/jngo/stack"
	"github.com/jnsoft/jngo/stringhelper"
)

type (
	Graph struct {
		v             int // Number of vertices
		e             int // Number of edges
		adjacencyList [][]int
	}
)

func (g *Graph) adj(v int) ([]int, error) {
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

func NewGraph(v int) (*Graph, error) {
	if v < 0 {
		return nil, errors.New("number of vertices must be nonnegative")
	}
	g := &Graph{
		v:             v,
		e:             0,
		adjacencyList: make([][]int, v),
	}
	for i := 0; i < v; i++ {
		g.adjacencyList[i] = []int{}
	}
	return g, nil
}

func NewGraphFromString(g string) (*Graph, error) {
	lines := stringhelper.ToLines(g)
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid input: no vertex count specified")
	}
	v, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse vertex count: %v", err)
	}

	graph, err := NewGraph(v)
	if err != nil {
		return nil, err
	}

	for i := 2; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" {
			continue
		}

		v1, w1, err := readEdge(lines[i])
		if err != nil {
			return nil, fmt.Errorf("failed to parse edge: %v", err)
		}
		err = graph.AddEdge(v1, w1)
		if err != nil {
			return nil, fmt.Errorf("failed to add edge: %v", err)
		}
	}
	return graph, nil
}

func (g *Graph) CopyGraph() *Graph {
	cp, _ := NewGraph(g.v)
	cp.e = g.e
	for v := 0; v < g.v; v++ {
		reverse := stack.New[int]()
		for w := range g.adjacencyList[v] {
			reverse.Push(w)
		}
		for !reverse.IsEmpty() {
			cp.adjacencyList[v] = append(cp.adjacencyList[v], reverse.Pop())
		}
	}
	return cp
}

func (g *Graph) AddEdge(v, w int, directed bool) error {
	if v < 0 || v >= g.v || w < 0 || w >= g.v {
		return errors.New("vertex out of bounds")
	}
	g.e++
	g.adjacencyList[v] = append(g.adjacencyList[v], w)
	if !directed {
		g.adjacencyList[w] = append(g.adjacencyList[w], v)
	}
	return nil
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
