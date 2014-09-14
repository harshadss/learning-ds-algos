//Package stack_ll provides implementation of stack using linked list,
//done for learning purpose.
package stack_ll

import "errors"

type item interface{}

type node struct{
    val item
    next *node}

type Stack struct{
    top *node
    size int }

func (s *Stack) Len() int {return s.size}
func (s *Stack) Push(val item) {
                              b := new(node)
                              b.val = val
                              b.next = s.top
                              s.top = b
                              s.size += 1 }
func (s *Stack) IsEmpty() bool { return s.top == nil }

func (s *Stack) Pop() (item, error) {
                    if s.IsEmpty(){
                       return 0, errors.New("pop on empty stack")
                     } else {
                        temp := s.top.val
                        s.top = s.top.next
                        s.size -= 1
                        return temp, nil } }
