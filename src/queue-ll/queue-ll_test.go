package queue_ll

import ("testing")

func TestIsEmptyOnEmptyQueue(t *testing.T) {
    tmp := new(Queue)
    if tmp.IsEmpty() {
        t.Log("IsEmpty on empty queue works as expected") 
    } else {
        t.Fail() 
    }
}

func TestIsEmptyOnNonEmptyQueue(t *testing.T) {
    tmp := new(Queue)
    tmp.Push(1)
    if !tmp.IsEmpty() {
        t.Log("IsEmpty on non empty queue works as expected") 
    } else {
        t.Fail() 
    }
}

func TestQueueFuctionality(t *testing.T) {
    tmp := new(Queue)
    tmp.Push(1)
    tmp.Push(2)
    tmp.Push(3)
    tmp.Push(4)
    first := tmp.Pop()
    second := tmp.Pop()
    third := tmp.Pop()
    fourth := tmp.Pop()
    //fmt.Println(first, second, third, fourth)
    if (first == 1) && (second == 2) && (third == 3) && (fourth == 4) {
        t.Log("queue push and pop works as expected")
    } else {
        t.Fail()
   }
}

func TestPopOnEmptyQueuePanics(t *testing.T){
    defer func(){
        if r := recover() ; r != nil {
            t.Log("popping from empty queue panics, as expected")
          }
      }()
    tmp := new(Queue)
    tmp.Pop()
}

func TestLenOnEmptyQueue(t *testing.T){
    tmp := new(Queue)
    if tmp.Len() == 0 {
        t.Log("len works on empty queue as expected")
    } else {
        t.Fail()
    }
}

func TestLenOnNonEmptyQueue(t *testing.T){
    tmp := new(Queue)
    tmp.Push(1)
    tmp.Push(2)
    if tmp.Len() == 2 {
       t.Log("len works on non empty queue as expected")
    } else {
       t.Fail() 
   }
}
