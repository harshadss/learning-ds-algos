package undirected_graph

import ("testing"
        "fmt")

func TestGraphCreation(t *testing.T){
  tmp := NewGraph(1)

  fmt.Println(tmp)
}
