//Package undirected-graph provides graph data structure representation
//using adjacency lists. For simple purpose of learning, the vertices are
//assumed to be integers.
package undirected_graph

import ("container/list"
        "fmt")

type Vertex struct{
  id int
  neighbours *list.List
}

type Graph struct{
   vertices []Vertex
   no_of_vertices int
   no_of_edges int
}

func NewGraph(size int) *Graph{
  g := new(Graph)
  g.vertices = make([]Vertex, size)
  g.no_of_vertices = size
  //Initialize the graph
  for v := 0 ; v < size ; v++ {
    g.vertices[v].id = v
    g.vertices[v].neighbours = list.New()
  }
  return g
}

func (g *Graph) AddEdge(u, v int) {

  g.vertices[u].neighbours.PushBack(v)
  g.vertices[v].neighbours.PushBack(u)
  g.no_of_edges += 1
}

func (g *Graph) GetNeighbours(u int) *list.List {
   if (g.no_of_vertices == 0){
     return nil
   } else {
      neighbour_list := g.vertices[u].neighbours 
      return neighbours_list
}
}

//BFS method implements breadth first search on the graph using FIFO queue.
//This is a dummy implementation which just prints all the nodes as they
//are traversed.

func (g *Graph) BFS(u int) {

  //First add the given vertex to a queue
  main_queue := list.New()
  main_queue.PushBack(g.vertices[u])

  for e := main_queue.Front() ; e != nil ; e = e.Next() {

    
}
}
   

  

