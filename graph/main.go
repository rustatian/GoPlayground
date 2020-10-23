package main

import (
	"bytes"
	"log"

	"github.com/goccy/go-graphviz"
)

func main() {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}

	n, err := graph.CreateNode("n")
	if err != nil {
		log.Fatal(err)
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		log.Fatal(err)
	}
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		log.Fatal(err)
	}

	e.SetLabel("e")
	// 1. write encoded PNG data to buffer
	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}

	// 2. get as image.Image instance
	image, err := g.RenderImage(graph)
	if err != nil {
		log.Fatal(err)
	}

	_ = image

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "./graph.png"); err != nil {
		log.Fatal(err)
	}
}
