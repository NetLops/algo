package matrixGraph

import "fmt"

type VertexType rune // 顶点类型由用户定义
type EdgeType int32  // 边上的权值类型应由用户定义
const (
	MAXVEX   = 100   // 最大顶点数，应由用户定义
	INFINITY = 65535 // 用65535代表♾️
)

type MGraph struct {
	vexs                  [MAXVEX]VertexType       // 顶点表
	arc                   [MAXVEX][MAXVEX]EdgeType // 邻接矩阵，可看作边表
	numVertexes, numEdges int32                    // 图中当前的顶点数和边数
}

// 建立无向图的邻接矩阵表示
func CreateMGraph(g *MGraph) {
	var i, j, k, w int32
	fmt.Println("输入顶点数和边数:")
	//fmt.Scanln((*g).numVertexes, (*g).numEdges)
	//a, b := 0, 0
	//_, err := fmt.Scanf("%d,%d", &a, &b)
	_, err := fmt.Scanf("%v,%v\n", &((*g).numVertexes), &((*g).numEdges))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i = 0; i < (*g).numVertexes; i++ {
		_, err2 := fmt.Scanln(&((*g).vexs[i]))
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
	for i = 0; i < (*g).numVertexes; i++ {
		for j = 0; j < (*g).numVertexes; j++ {
			(*g).arc[i][j] = INFINITY // 邻接矩阵初始化
		}
	}
	for k = 0; k < (*g).numEdges; k++ { // 读入numEdges条边，建立邻接矩阵
		fmt.Println("输入边(vi, vj)上的下标i，下标j和权w:")
		fmt.Scanf("%d,%d,%d\n", &i, &j, &w) // 插入边(vi, vj)上的权w
		(*g).arc[i][j] = EdgeType(w)
		(*g).arc[j][i] = (*g).arc[i][j]
	}
	fmt.Println(*g)
}
