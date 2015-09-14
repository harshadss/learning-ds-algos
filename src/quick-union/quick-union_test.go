package quick_union

import ("testing")

func TestRootEmptyQF(t *testing.T) {
   uf := NewUf(10)
   flag := true
   for i := 0 ; i < 10 ; i++ {
	if uf.findRoot(i) != i {
            flag = false 
         }
   }
   if flag {
       t.Log("Root of empty strucutres are correct")
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
   if !(uf.Find(1, 9)) {
        t.Log("Negative find test passes")
    } else {
        t.Fail()
   }
}

func TestFindPositive(t *testing.T) {
   uf := NewUf(10)
   uf.Union(1, 2)
   uf.Union(2, 3)
   uf.Union(3, 5)
   uf.Union(5, 7)
   uf.Union(7, 8)
   uf.Union(8, 9)
   if (uf.Find(1, 9)) {
        t.Log("Positive find test passes")
    } else {
        t.Fail()
   }
}

func BenchmarkFind(b *testing.B) {
   uf := NewUf(1000)
   for i := 0 ; i < 999 ; i++ {
       uf.Union(i, i + 1)
   }
   b.ResetTimer()
   for i := 0 ; i < b.N ; i++ {
	   _ = uf.Find(1, 999)
  }
}

func BenchmarkUnion(b *testing.B) {
   uf := NewUf(1000)
   for i := 0 ; i < 998 ; i ++ {
     // setting all to be connected
        uf.Union(i + 1 , i + 2)
   }
   b.ResetTimer()
  // Invoking the worst case
  for j := 0 ; j < b.N ; j++ {
	uf.Union(0, 1)
 }
}
