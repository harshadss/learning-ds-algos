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
	def __init__(self, size, is_directed = False):
		self.no_of_vertices = size
		self.no_of_edges = 0
		self.vertices = list()
		self.is_directed = is_directed
		for i in xrange(size):
			self.vertices.append(Vertex(i))

	def AddEdge(self, src, dest):
		self.vertices[src].neighbours.append(self.vertices[dest])
		if not self.is_directed:
			self.vertices[dest].neighbours.append(self.vertices[src])
		self.no_of_edges += 1

	def BreadthFirstSearch(self, u):
		explored_nodes = set([self.vertices[u]])
		candidate_nodes = deque([self.vertices[u]])
		while candidate_nodes:
			candidate = candidate_nodes.pop() # FIFO
			for neighbour in candidate.GetNeighbours():
				if neighbour not in explored_nodes:
					print ("explored %d" % neighbour.nodeid)
					explored_nodes.add(neighbour)
					candidate_nodes.appendleft(neighbour) 
		return explored_nodes

	def DepthFirstSearch(self, u):
		explored_nodes = set([self.vertices[u]])
		candidate_nodes = deque([self.vertices[u]])
		while candidate_nodes:
			candidate = candidate_nodes.popleft() # LIFO
			for neighbour in candidate.GetNeighbours():
				if neighbour not in explored_nodes:
					print ("explored %d" % neighbour.nodeid)
					explored_nodes.add(neighbour)
					candidate_nodes.appendleft(neighbour) 
		return explored_nodes
