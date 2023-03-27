package main

import (
	"fmt"
	"github.com/dominikbraun/graph"
)

func main() {
	fmt.Println("Hello, World!")
	g := graph.New(graph.IntHash)

	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)
	_ = g.AddVertex(5)

	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(1, 4)
	_ = g.AddEdge(2, 3)
	_ = g.AddEdge(2, 4)
	_ = g.AddEdge(2, 5)
	_ = g.AddEdge(3, 5)

	// DFS 实现
	_ = graph.DFS(g, 1, func(value int) bool {
		fmt.Println(value)
		return false
	})

	//最大联通子图
	scc, _ := graph.StronglyConnectedComponents(g)
	fmt.Println(scc)

	//排序
	order, _ := graph.TopologicalSort(g)
	fmt.Println(order)

}
