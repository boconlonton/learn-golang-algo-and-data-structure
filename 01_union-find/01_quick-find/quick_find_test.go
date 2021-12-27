package main

import (
	"testing"
)

func TestQuicKFind(t *testing.T) {
	p := QuickFindUF{id: []int{}}
	p.initialize(10)
	if p.connected(1, 2) {
		t.Error("expected False got True")
	}
	p.union(1, 2)
	if !(p.connected(1, 2)) {
		t.Error("expected True got False")
	}
}
