package mario

type Status uint8
type CharType uint8

const (
    Status_Init Status = iota
    Status_Prefix_Space
    Status_Number_Sign
    Status_Integer
    Status_Point_With_Number_Before
    Status_Point_Without_Number_Before
    Status_Decimal
    Status_E
    Status_E_Sign
    Status_E_Integer
    Status_Suffix_Space
)

const (
    Type_Space CharType = iota
    Type_Sign
    Type_Number
    Type_Point
    Type_E
)

func isNumber(s string) bool {
    fsm := makeFSM()

    var isNum, ok bool
    status := Status_Init
    for i := range s {
        status, ok = getFSMNextStatus(fsm, status, getCharType(s[i]))
        if !ok {
            isNum = false
            break
        }
        if status == Status_Integer || status == Status_Decimal {
            isNum = true
        }
    }

    return isNum
}

func getFSMNextStatus(fsm map[Status]map[CharType]Status, statusNow Status, nextChar CharType) (status Status, ok bool) {
    var t map[CharType]Status
    t, ok = fsm[statusNow]
    if !ok {
        ok = false
        return
    }

    status, ok = t[nextChar]
    if !ok {
        ok = false
        return
    }

    ok = true

    return
}

func makeFSM() map[Status]map[CharType]Status {
    fsm := make(map[Status]map[CharType]Status)
    fsm[Status_Init] = map[CharType]Status{
        Type_Space: Status_Prefix_Space,
        Type_Sign: Status_Number_Sign,
        Type_Number: Status_Integer,
        Type_Point: Status_Point_Without_Number_Before,
    }
    fsm[Status_Prefix_Space] = map[CharType]Status{
        Type_Space: Status_Prefix_Space,
        Type_Sign: Status_Number_Sign,
        Type_Number: Status_Integer,
        Type_Point: Status_Point_Without_Number_Before,
    }
    fsm[Status_Number_Sign] = map[CharType]Status{
        Type_Number: Status_Integer,
        Type_Point: Status_Point_Without_Number_Before,
    }
    fsm[Status_Integer] = map[CharType]Status{
        Type_Number: Status_Integer,
        Type_Point: Status_Point_With_Number_Before,
        Type_E: Status_E,
    }
    fsm[Status_Point_With_Number_Before] = map[CharType]Status{
        Type_Number: Status_Decimal,
        Type_E: Status_E,
    }
    fsm[Status_Point_Without_Number_Before] = map[CharType]Status{
        Type_Number: Status_Decimal,
    }
    fsm[Status_Decimal] = map[CharType]Status{
        Type_Number: Status_Decimal,
        Type_E: Status_E,
    }
    fsm[Status_E] = map[CharType]Status{
        Type_Sign: Status_E_Sign,
        Type_Number: Status_E_Integer,
    }
    fsm[Status_E_Sign] = map[CharType]Status{
        Type_Number: Status_E_Integer,
    }
    fsm[Status_E_Integer] = map[CharType]Status{
        Type_Number: Status_E_Integer,
    }
    fsm[Status_Suffix_Space] = map[CharType]Status{
        Type_Space: Status_Suffix_Space,
    }

    return fsm
}

func getCharType(char byte) (typ CharType) {
    switch char {
    case ' ': typ = Type_Space
    case '+', '-': typ = Type_Sign
    case '0','1','2','3','4','5','6','7','8','9': typ = Type_Number
    case '.': typ = Type_Point
    case 'e', 'E': typ = Type_E
    }

    return
}
