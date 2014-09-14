package stack_ll

import ("testing")

func TestIsEmptyPositive(t *testing.T) {
    tmp := new(Stack)
    if tmp.IsEmpty() {
      t.Log("positive IsEmpty test cleared")
    } else {
      t.Fail() }
}

func TestIsEmptyNegative(t *testing.T) {
    tmp := new(Stack)
    tmp.Push(3)
    if tmp.IsEmpty() {
        t.Fail()
    } else {
      t.Log("negative IsEmpty test cleared") }
}

func TestPopOnEmptyStack(t *testing.T) {
    tmp := new(Stack)
    _, err := tmp.Pop()
    if err != nil {
        t.Log("pop on empty stack test cleared")
    } else {
        t.Fail()
 }
}

func TestPushPop(t *testing.T){
    tmp := new(Stack)
    tmp.Push(1)
    tmp.Push(2)
    val, err := tmp.Pop()
    if (val == 2) && (err == nil) {
        t.Log("stack push-pop functionality test cleared")
    } else {
        t.Fail()
   }
} 
