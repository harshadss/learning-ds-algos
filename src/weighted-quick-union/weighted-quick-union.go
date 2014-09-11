//Package weighted_quick_union offers data structures and algorithms for the
//union-find problem. The union operation is implemented such that the
//the depth of the tree is always limited to lg N. The proof is on the
//following lines : In a union operation of i and j, we add tree containing
// i to tree containing j only if size i.e. number of elements of tree
//containing j is more than the size of tree containing i. Suppose that the
//size of tree containing i is x. Then the union addition will have resulting
//tree with size atleast 2 *x as tree containing j is atleast of size x, for
//merge condition to meet. Thus, for any element i, the size of the tree can
// double at the most lg N times (as starting with 1, we can double at the most
// lg N time before exhausting all the elements available)
package weighted_quick_union

//The data structure has additional array s for maintaining the size for each
// tree rooted at the index.
type Uf struct {
	arr []int
	s   []int
}

//NewUf returns a pointer to the union-find data type. The constructor also
//initializes all the elements correctly.
func NewUf(size int) *Uf {
	uf := new(Uf)
	uf.arr = make([]int, size)
	uf.s = make([]int, size)
	for ind := range uf.arr {
		uf.arr[ind] = ind
		uf.s[ind] = 1 // all trees start with single element
	}
	return uf
}

//findRoot is a private method which traverses the path till root and
//returns the root index value
func (u *Uf) findRoot(component int) int {
	for component != u.arr[component] {
		component = u.arr[component]
	}
	return component
}

func (u *Uf) findSize(component int) int {
	return u.s[component]
}

//Find operation returns a flag indicating whether elements i and j are
//connected. It simply checks if the two elements belong to tree rooted
//at the same element.
func (u *Uf) Find(i int, j int) bool {
	return (u.findRoot(i) == u.findRoot(j))
}

//Union operation implements the weighted union logic. For creating union
//of elements i and j, it first finds out root element for each. Then it
//checks the size of the tree to which each element i and j belongs. The
//union is implemented by changing the root of smaller tree to root of bigger
//tree always. This ensures that the data structure never ends up with really
//long trees ala linked list. See package documentation for proof sketch.
func (uf *Uf) Union(i int, j int) {

	root_i := uf.findRoot(i)
	size_i := uf.findSize(root_i)
	root_j := uf.findRoot(j)
	size_j := uf.findSize(root_j)

	if i == j {
		return
	} else if size_i < size_j {
		uf.arr[root_i] = uf.arr[root_j]
		uf.s[root_j] += size_i
	} else {
		uf.arr[root_j] = uf.arr[root_i]
		uf.s[root_i] += size_j
	}
}
