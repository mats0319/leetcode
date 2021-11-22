package mario

import "testing"

type resData struct {
	isInt     bool
	intValue  int
	isBool    bool
	boolValue bool
}

func TestCase1(t *testing.T) {
	res := make([]*resData, 0)

	mcd := Constructor(3)
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertLast(1)})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertLast(2)})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertFront(3)})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertFront(4)})
	res = append(res, &resData{isInt: true, intValue: mcd.GetRear()})
	res = append(res, &resData{isBool: true, boolValue: mcd.IsFull()})
	res = append(res, &resData{isBool: true, boolValue: mcd.DeleteLast()})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertFront(4)})
	res = append(res, &resData{isInt: true, intValue: mcd.GetFront()})

	if !assertResDataSlice(res, []*resData{
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: false},
		{isInt: true, intValue: 2},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isInt: true, intValue: 4},
	}) {
		t.Errorf("641 test case 1 failed.")
	}
}

func TestCase2(t *testing.T) {
	res := make([]*resData, 0)

	mcd := Constructor(8)
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertFront(5)})
	res = append(res, &resData{isInt: true, intValue: mcd.GetFront()})
	res = append(res, &resData{isBool: true, boolValue: mcd.IsEmpty()})
	res = append(res, &resData{isBool: true, boolValue: mcd.DeleteFront()})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertLast(3)})
	res = append(res, &resData{isInt: true, intValue: mcd.GetRear()})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertLast(7)})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertFront(7)})
	res = append(res, &resData{isBool: true, boolValue: mcd.DeleteLast()})
	res = append(res, &resData{isBool: true, boolValue: mcd.InsertLast(4)})
	res = append(res, &resData{isBool: true, boolValue: mcd.IsEmpty()})

	if !assertResDataSlice(res, []*resData{
		{isBool: true, boolValue: true},
		{isInt: true, intValue: 5},
		{isBool: true, boolValue: false},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isInt: true, intValue: 3},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: true},
		{isBool: true, boolValue: false},
	}) {
		t.Errorf("641 test case 2 failed.")
	}
}

func assertResDataSlice(a, b []*resData) bool {
	if len(a) != len(b) {
		return false
	}

	isEqual := true
	for i := 0; i < len(a); i++ {
        if (a[i].isInt && b[i].isInt && a[i].intValue != b[i].intValue) ||
			(a[i].isBool && b[i].isBool && a[i].boolValue != b[i].boolValue) {
            isEqual = false
            break
        }
	}

	return isEqual
}
