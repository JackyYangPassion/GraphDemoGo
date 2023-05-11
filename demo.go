package main

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"os"
)

/**
 *目标：
 * 1. 图基本操作 CRUD
 * 2. 算法 基本操作
 * 3. 可视化
 */
func main() {
	fmt.Println("demo graph!")
	//初始化图
	/**
	*1. 增加点
	*2. 增加边
	*3. 图基本操作
	 */
	g := graph.New(graph.IntHash, graph.Directed(), graph.PreventCycles())

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

	/**
	 * 图算法 demo
	 */
	// DFS 实现
	_ = graph.DFS(g, 1, func(value int) bool {
		fmt.Println(value)
		return false
	})

	//最大联通子图
	scc, _ := graph.StronglyConnectedComponents(g)
	fmt.Println("最大联通子图： ", scc)

	//排序
	order, _ := graph.TopologicalSort(g)
	fmt.Println("节点排序： ", order)

	//最短路径
	path, _ := graph.ShortestPath(g, 1, 5)
	fmt.Println("两点最短路径： ", path)

	/**
	 * 图可视化
	 * dot -Tsvg -Kneato -O mygraph.gv
	 */
	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)

	//支持自定义后端：KV等
	//g := graph.NewWithStore(graph.IntHash, myStore)

	//TODO: test demo

}
