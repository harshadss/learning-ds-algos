package quick_union

type Uf struct{
	arr []int } //arr is a slice of integers

func NewUf(size int) *Uf {
    uf := new(Uf)
    uf.arr = make([]int, size) // arr field is already known so no :=
    // iterate and initialize each
    // This is important : else empty one will have everything connected
    for ind := range(uf.arr) {
        uf.arr[ind] = ind
}
    return uf } // new gives a pointer to struct

func (u *Uf) findRoot(component int) int{
    for component != u.arr[component] {
        component = u.arr[component]
    }
    return component }

func (u *Uf) Find(i int, j int) bool {
    // checks whether i and j are connected to each other
    // simply check if i and j have the same value in it
    return (u.findRoot(i) == u.findRoot(j) ) }

func (uf *Uf) Union(i int, j int) {

  // If the components are same, just return 
  // Else, set the root of j equal to root of i
  if i == j {
     return 
  } else {
       root_i := uf.findRoot(i)
       root_j := uf.findRoot(j)
       uf.arr[root_j] = uf.arr[root_i]
}
}
