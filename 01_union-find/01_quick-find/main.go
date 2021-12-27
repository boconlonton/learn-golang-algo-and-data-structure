package main

import "fmt"

type QuickFindUF struct {
	id []int
}

func main() {
	p := QuickFindUF{id: []int{}}
	p.initialize(10)
	fmt.Println(p.connected(1, 2))
	p.union(1, 2)
	fmt.Println(p.connected(1, 2))

}

// initialize set id of each object to itself.
func (q *QuickFindUF) initialize(N int) {
	for i := 0; i < N; i++ {
		(*q).id = append((*q).id, i)
	}
}

// connected check whether p and q are in the same component.
func (q *QuickFindUF) connected(x, y int) bool {
	return (*q).id[x] == (*q).id[y]
}

func (q *QuickFindUF) union(x, y int) {
	tmp := (*q).id
	xid := tmp[x]
	yid := tmp[y]
	for i := 0; i < len((*q).id); i++ {
		if tmp[i] == xid {
			tmp[i] = yid
		}
	}
}
