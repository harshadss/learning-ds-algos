package quick_find

import (
	"testing"
)

func TestFindPositive(t *testing.T) {
	uf := NewUf(10)
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 5)
	uf.Union(5, 4)
	if uf.Find(1, 4) && uf.Find(2, 5) {
		t.Log("positive find test passes")
	} else {
		t.Fail()
	}
}

func TestFindNegative(t *testing.T) {
	uf := NewUf(10)
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 5)
	uf.Union(4, 7)
	uf.Union(7, 8)
	uf.Union(8, 9)
	if !(uf.Find(1, 9)) && !(uf.Find(5, 4)) {
		t.Log("negative find test passes")
	} else {
		t.Fail()
	}
}

//BenchmarkFind is used to check the speed of find operation. Absolute number
//is not important, but the relative value with other implementations of
//union-find (e.g. quick-union and weighted-quick-union) is to be checked.
func BenchmarkFind(b *testing.B) {
	uf := NewUf(1000)
	//Setup
	for i := 0; i < 999; i++ {
		uf.Union(i, i+1)
	}
	b.ResetTimer() //Reset timer as we want to benchmark only find
	for i := 0; i < b.N; i++ {
		_ = uf.Find(1, 999)
	}
}

//BenchmarkUnion benchmarks the worst case where a large number of elements
//need to be re-assigned due to union operation.
func BenchmarkUnion(b *testing.B) {
	uf := NewUf(1000)
	for i := 0; i < 998; i++ {
		uf.Union(i+1, i+2)
	}
	b.ResetTimer()
	// Invoking the worst case
	for j := 0; j < b.N; j++ {
		uf.Union(0, 1)
	}
}
