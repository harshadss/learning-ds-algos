# This module implements graph datastructure using python lists
# The idea is to get a high level understanding of graph data
# structure and the associated operations

from collections import deque

class Vertex:
	def __init__(self, nodeid):
		self.nodeid = nodeid
		self.neighbours = list()

	def GetNeighbours(self):
		return self.neighbours

class Graph:
	def __init__(self, size):
		self.no_of_vertices = size
		self.no_of_edges = 0
		self.vertices = list()
		for i in xrange(size):
			tmp_vertex = Vertex(i)
			self.vertices.append(tmp_vertex)

	def AddEdge(self, u, v):
		self.vertices[u].neighbours.append(self.vertices[v])
		self.no_of_edges += 1

	def BreadthFirstSearch(self, u):
		all_vertices = self.vertices
		given_vertex = all_vertices[u]
		exploration_queue = deque([given_vertex])
		explored_nodes = []
		while exploration_queue:
			vertex = exploration_queue.pop()
			if vertex in explored_nodes:
				continue
			else:
				explored_nodes.append(vertex)
			for neighbour in vertex.GetNeighbours():
				exploration_queue.appendleft(neighbour)	
		return explored_nodes

	def DepthFirstSearch(self, u):
		all_vertices = self.vertices
		given_vertex = all_vertices[u]
		exploration_queue = deque([given_vertex])
		explored_nodes = []
		hops = 0
		while exploration_queue:
			vertex = exploration_queue.popleft()
			print ("now at node: %d", vertex.nodeid)
			if vertex in explored_nodes:
				print ("node %d is already visited. Backtracking.", vertex.nodeid)
				continue
			else:
				explored_nodes.append(vertex)
			for neighbour in vertex.GetNeighbours():
				hops += 1
				exploration_queue.appendleft(neighbour)	
				print [ver.nodeid for ver in exploration_queue]
		print hops
		return explored_nodes
