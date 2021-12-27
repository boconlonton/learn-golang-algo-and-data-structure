package main

import "fmt"

type QuickUnionUF struct {
	id []int
	sz []int // count number of objects in the tree rooted at i.
}

func main() {
	p := QuickUnionUF{id: []int{}, sz: []int{}}
	p.initialize(10)
	fmt.Println(p.connected(1, 2))
	p.union(1, 2)
	fmt.Println(p.connected(1, 2))

}

// initialize set id of each object to itself.
func (q *QuickUnionUF) initialize(N int) {
	for i := 0; i < N; i++ {
		(*q).id = append((*q).id, i)
		(*q).sz = append((*q).sz, 1)
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

// union change root of p to point to root of q.
// modified to link root of smaller tree to root of larger tree.
// update the sz[] array.
func (q *QuickUnionUF) union(x, y int) {
	tmpArr := (*q).id
	tmpSize := (*q).sz
	i := q.root(x)
	j := q.root(y)
	if i == j {
		return
	}
	if tmpSize[x] < tmpSize[y] {
		tmpArr[i] = j
		tmpSize[j] += tmpSize[i]
	} else {
		tmpArr[j] = i
		tmpSize[i] += tmpSize[j]
	}
}
