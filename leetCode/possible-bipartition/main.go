package main

type UnionFind struct {
	p []int
}

func (u *UnionFind) find(x int) int {
	if u.p[x] != x {
		u.p[x] = u.find(u.p[x])
	}
	return u.p[x]
}

func (u *UnionFind) union(a, b int) {
	u.p[u.find(a)] = u.p[u.find(b)]
}

func (u *UnionFind) query(a, b int) bool {
	return u.find(a) == u.find(b)
}

// 并查集
func possibleBipartition(n int, dislikes [][]int) bool {
	u := UnionFind{
		p: make([]int, 4010),
	}
	for i := 1; i < 2*n+1; i++ {
		u.p[i] = i
	}
	for _, dislike := range dislikes {
		a, b := dislike[0], dislike[1]
		if u.query(a, b) {
			return false
		}
		u.union(a, b+n)
		u.union(b, a+n)

		/*
			 a -> b+n
			 b -> a+n

			 a -> c+n
			 c -> a+n

			// 此时就判断 b, c是联通的
		*/

	}
	return true
}
