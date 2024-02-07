// dir graph, undir graph, weighted, unweighted
package graph

import (
	"encoding/json"
	"fmt"

	"github.com/lorenzotinfena/goji/collections"
	"github.com/lorenzotinfena/goji/math"
	"github.com/lorenzotinfena/goji/utils/constraints"
)

// Using code from: github.com/hediet/vscode-debug-visualizer/blob/master/demos/golang/demo.go

type nodeGraphDataDebug struct {
	ID    string `json:"id"`
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
	Shape string `json:"shape,omitempty"`
}

type edgeGraphDataDebug struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Label  string `json:"label,omitempty"`
	ID     string `json:"id"`
	Color  string `json:"color,omitempty"`
	Dashes bool   `json:"dashes,omitempty"`
}

type graphDebug struct {
	Kind  map[string]bool      `json:"kind"`
	Nodes []nodeGraphDataDebug `json:"nodes"`
	Edges []edgeGraphDataDebug `json:"edges"`
}

func UnitString[V comparable](
	vertices []V,
	adjacents func(V) []V,
	colors map[V]int,
) string {
	colorsbase := []string{
		"pink",
		"red",
		"yellow",
		"green",
		"cyan",
		"blue",
		"saddlebrown",
		"Slateblue",
		"darkorange",
		"white",
	}
	getColor := func(v V) string {
		if colors == nil {
			return "white"
		} else {
			return colorsbase[math.Min(len(colorsbase)-1, colors[v])]
		}
	}

	graph1 := &graphDebug{
		Kind:  map[string]bool{"graph": true},
		Nodes: []nodeGraphDataDebug{},
		Edges: []edgeGraphDataDebug{},
	}
	for _, v := range vertices {
		s := fmt.Sprint(v)
		graph1.Nodes = append(graph1.Nodes, nodeGraphDataDebug{ID: s, Label: s, Color: getColor(v)})
	}

	for _, v := range vertices {
		for _, adj := range adjacents(v) {
			s1 := fmt.Sprint(v)
			s2 := fmt.Sprint(adj)
			tmp := edgeGraphDataDebug{From: s1, To: s2}
			graph1.Edges = append(graph1.Edges, tmp)
		}
	}
	s, _ := json.Marshal(graph1)
	return string(s)
}

func WeightedString[V comparable, W constraints.Integer | constraints.Float](
	vertices []V,
	adjacents func(V) []collections.Pair[V, W],
	colors map[V]int,
) string {
	colorsbase := []string{
		"pink",
		"red",
		"yellow",
		"green",
		"cyan",
		"blue",
		"saddlebrown",
		"Slateblue",
		"darkorange",
		"white",
	}
	getColor := func(v V) string {
		if colors == nil {
			return "white"
		} else {
			return colorsbase[math.Min(len(colorsbase)-1, colors[v])]
		}
	}

	graph1 := &graphDebug{
		Kind:  map[string]bool{"graph": true},
		Nodes: []nodeGraphDataDebug{},
		Edges: []edgeGraphDataDebug{},
	}
	for _, v := range vertices {
		s := fmt.Sprint(v)
		graph1.Nodes = append(graph1.Nodes, nodeGraphDataDebug{ID: s, Label: s, Color: getColor(v)})
	}

	for _, v := range vertices {
		for _, adj := range adjacents(v) {
			s1 := fmt.Sprint(v)
			s2 := fmt.Sprint(adj.First)
			tmp := edgeGraphDataDebug{From: s1, To: s2}
			tmp.Label = fmt.Sprint(adj.Second)
			graph1.Edges = append(graph1.Edges, tmp)
		}
	}
	s, _ := json.Marshal(graph1)
	return string(s)
}
