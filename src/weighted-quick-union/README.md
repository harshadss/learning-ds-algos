Weighted quick union is the enhanced implementation of union-find data
structure as covered by Prof. Robert Sedgewick. This implementation is
much better than the simple quick-find implementation which has fast find
but slow union operation. 

On a intel core i3 1.7GHz machine, the simple implmentation quick-find has find
operation on a 1000 element problem running in 3 ns, however the union 
operations worst case takes 2520 ns. The improved implementation of weighted
union find which sacrifices find speed for better union implementation has
find and union both running in approx 23 ns. These numbers are calculated
using benchmarking facilities provided by go testing library.
