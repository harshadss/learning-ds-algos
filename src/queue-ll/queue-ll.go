//Package queue-ll provides implementation of FIFO queue using linked list.
package queue_ll

//import "fmt"

type node struct{
    val interface{}
    previous *node
    next *node}

type Queue struct{
    head *node
    tail *node
    len int }

func (s *Queue) Len() int {return s.len }

func (s *Queue) IsEmpty() bool {return s.head == nil }

//Function Push pushes the given value at the end of queue. If the queue is
//empty, create a new node, set head and tail pointers to that node.
//If the queue is not empty, create a new node. Set the new node's next pointer
//to the previous tail. Set the previous pointer of current tail node to the
//new node. Set the queue's tail to point to this new node.
//Attempting to insert nil value panics.
func (s *Queue) Push(val interface{}) {
    if s.IsEmpty() {
        tmp := new(node)
        tmp.val = val
        s.head = tmp
        s.tail = tmp
        s.len += 1
    } else {
        current_tail := s.tail
        tmp := new(node)
        tmp.val = val
	current_tail.previous = tmp
        tmp.next = current_tail
        s.tail = tmp
        s.len += 1
    }
}

//Function Pop pops one item from the front of queue. If the queue is empty
//return nil (assuming that clients do not use nil as legitimate value).
//If the queue is not empty, get the value from the node at the head. Set
//the head to point to the previous node of current head
func (s *Queue) Pop() interface{} {
                if s.IsEmpty(){
                   panic("blah")
                } else {
                   current_head := s.head
                   val := current_head.val
                   s.head = current_head.previous
                   current_head.val = nil 
                   return val }
}
