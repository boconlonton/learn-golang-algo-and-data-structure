package main

import (
	"testing"
)

func TestQuicKUnion(t *testing.T) {
	p := QuickUnionUF{id: []int{}}
	p.initialize(10)
	if p.connected(1, 2) {
		t.Error("expected False got True")
	}
	p.union(1, 2)
	if !(p.connected(1, 2)) {
		t.Error("expected True got False")
	}
}
