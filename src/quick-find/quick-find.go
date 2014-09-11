//Package quick_find provides the data structures and algorithms for union-find
//problem. The implementation uses array of N elements and maintains the
//invariant that the value of indices of connected elements is the same.
//This makes find operation very fast as it does two array lookups and one
//comparison. The union operation ends up scanning over all elements of the
//grid, making it slow.
package quick_find

type Uf struct {
	arr []int
}

//NewUf returns a pointer to the union-find data type. The constructor also
//initializes all the elements correctly.
func NewUf(size int) *Uf {
	uf := new(Uf)
	uf.arr = make([]int, size)
	for ind := range uf.arr {
		uf.arr[ind] = ind
	}
	return uf
}

// Find returns a flag indicating whether element i is connected to element j
// or not. Simply checks if corresponding array indices have the same value.
func (u *Uf) Find(i int, j int) bool {
	return (u.arr[i] == u.arr[j])
}

// Union operation connects element i to element j, by simply assigning the
// value of elements in cluster of i to all elements in cluster of j. This
// scans over the entire array for a single union operation and hence will
// be slow for large sizes.
func (uf *Uf) Union(i int, j int) {
	to_be_changed := -1
	set_to := -1
	if i == j {
		return
	} else {
		//Arbitrarily selects the value at index of first element and
		// assigns the same for the connected components of other element.
		to_be_changed = uf.arr[j]
		set_to = uf.arr[i]
	}

	for ind := range uf.arr {
		if uf.arr[ind] == to_be_changed {
			uf.arr[ind] = set_to
		}
	}
	return
}
