package recaps

import "fmt"

// Node struct represents a node in the singly linked list.
type Node struct {
	Info string
	Next *Node
}

// List struct represents a singly linked list data structure.
type List struct {
	pHead *Node
	pTail *Node
}

// CreateNode method initialize a node based on the data info.
func (n *Node) CreateNode(d string) *Node {
	return &Node{Info: d, Next: nil}
}

// Init method initialize an empty list.
func (l *List) Init() {
	l.pHead = nil
	l.pTail = nil
}

// AddHead method add a node at the beginning of a linked list.
func (l *List) AddHead(p *Node) {
	if l.pHead != nil {
		l.pHead, l.pTail = p, l.pHead
	} else {
		l.pHead, p.Next = p, l.pTail
	}
}

// AddTail method add a node at the end of a linked list.
func (l *List) AddTail(p *Node) {
	if l.pHead != nil {
		l.pHead, l.pTail = p, l.pHead
	} else {
		l.pTail.Next = p
		l.pTail = p
	}
}

// AddAfterQ method add a node at a custom position which after Q node.
func (l *List) AddAfterQ(p *Node, q *Node) {
	if q != nil {
		q.Next, p.Next = p, q.Next
		if l.pTail == q {
			// q is the last node
			l.pTail = p
		}
	} else {
		// Add p to the head
		l.AddHead(p)
	}
}

// Scan method travels through all the items in the list.
func (l *List) Scan() {
	p := l.pHead
	for p != nil {
		fmt.Printf("Node(info=%v)", p.Info)
		p = p.Next
	}
}

// Find method find & return the node which has the required info.
func (l *List) Find(s string) *Node {
	p := l.pHead
	for p != nil {
		if p.Info == s {
			return p
		}
		p = p.Next
	}
	return nil
}

// RemoveHead method remove the first node in the list.
func (l *List) RemoveHead(x *string) bool {
	if l.pHead != nil {
		*x = l.pHead.Info
		l.pHead = l.pHead.Next
		return true // Success
	}
	return false // Empty list
}

// RemoveAfterQ method remove node at a specific position.
func (l *List) RemoveAfterQ(q *Node, x *string) bool {
	if q != nil {
		t := q.Next // targeted node.
		if t != nil {
			// q isn't the last node.
			if t == l.pTail {
				// targeted node is the last node => change the tail node to q (nearest)
				l.pTail = q
			}
			q.Next = t.Next
			*x = q.Next.Info
		}
		return true
	}
	return false
}

// RemoveTail method remove node at the end of the list.
func (l *List) RemoveTail(x *string) bool {
	t := l.pHead
	for t.Next.Next != nil {
		t = t.Next
	}
	return l.RemoveAfterQ(t, x)
}
