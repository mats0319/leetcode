package mario

import "testing"

func TestCase1(t *testing.T) {
    arr := Constructor(1)
    arr.Push(0, 1)
    arr.Push(0, 2)

    values := make([]int, 0)
    values = append(values, arr.Pop(0))
    values = append(values, arr.Pop(0))
    values = append(values, arr.Pop(0))

    tag := arr.IsEmpty(0)

    if len(values) != 3 || values[0] != 1 || values[1] != -1 || values[2] != -1 || !tag {
        t.Errorf("test case 1 runs wrong.")
    }
}
