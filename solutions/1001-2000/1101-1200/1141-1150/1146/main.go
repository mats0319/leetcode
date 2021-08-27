package mario

type SnapshotArray struct {
	snapID  int
	changes map[[2]int]int // [snap_id, index] - new value，go语言，能使用“==”的，都能用作map的key
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		changes: make(map[[2]int]int),
	}
}

func (sa *SnapshotArray) Set(index int, val int) {
	sa.changes[[2]int{sa.snapID, index}] = val
}

func (sa *SnapshotArray) Snap() int {
	id := sa.snapID
	sa.snapID++

	return id
}

func (sa *SnapshotArray) Get(index int, snap_id int) int {
	res := 0
	for i := snap_id; i >= 0; i-- {
		v, ok := sa.changes[[2]int{i, index}]
		if ok {
			res = v
			break
		}
	}

	return res
}
