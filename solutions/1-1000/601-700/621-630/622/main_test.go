package mario

import (
    "fmt"
    "testing"
)

func TestMyCircularQueueCase1(t *testing.T) {
    ins := Constructor(2)
    fmt.Printf("%v ", ins.EnQueue(8))
    fmt.Printf("%v ", ins.EnQueue(8))
    fmt.Printf("%v ", ins.Front())
    fmt.Printf("%v ", ins.EnQueue(4))
    fmt.Printf("%v ", ins.DeQueue())
    fmt.Printf("%v ", ins.EnQueue(1))
    fmt.Printf("%v ", ins.EnQueue(1))
    fmt.Printf("%v ", ins.Rear())
    fmt.Printf("%v ", ins.IsEmpty())
    fmt.Printf("%v ", ins.Front())
    fmt.Printf("%v ", ins.DeQueue())
    fmt.Println()
    fmt.Println("true true 8 false true true false 1 false 8 true")
}

func TestMyCircularQueueCase2(t *testing.T) {
    ins := Constructor(81)
    fmt.Printf("%v ", ins.EnQueue(69))
    fmt.Printf("%v ", ins.DeQueue())
    fmt.Printf("%v ", ins.EnQueue(92))
    fmt.Printf("%v ", ins.EnQueue(12))
    fmt.Printf("%v ", ins.DeQueue())
    fmt.Printf("%v ", ins.IsFull())
    fmt.Printf("%v ", ins.IsFull())
    fmt.Printf("%v ", ins.Front())
    fmt.Println()
    fmt.Println("true true true true true false false 12")
}
