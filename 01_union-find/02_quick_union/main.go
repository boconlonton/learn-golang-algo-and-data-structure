package main

import "fmt"

type QuickUnionUF struct {
	id []int
}

func main() {
	p := QuickUnionUF{id: []int{}}
	p.initialize(10)
	fmt.Println(p.connected(1, 2))
	p.union(1, 2)
	fmt.Println(p.connected(1, 2))

}

// initialize set id of each object to itself.
func (q *QuickUnionUF) initialize(N int) {
	for i := 0; i < N; i++ {
		(*q).id = append((*q).id, i)
	}
}

// root chase parent pointers until reach root
func (q *QuickUnionUF) root(x int) int {
	tmp := (*q).id
	for tmp[x] != x {
		x = tmp[x]
	}
	return x
}

// connected check whether p and q are in the same component.
func (q *QuickUnionUF) connected(x, y int) bool {
	return q.root(x) == q.root(y)
}

func (q *QuickUnionUF) union(x, y int) {
	tmp := (*q).id
	tmp[q.root(x)] = q.root(y)
}
