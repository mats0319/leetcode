package mario

type Robot struct {
	w int
	h int

	side1 int // w
	side2 int // w+h
	side3 int // w+h+w
	side4 int // w+h+w+h, also 'a' of road

	steps int  // total steps % side4
	init  bool // init mod, face 'east'

	// init mod: (0, 0) east,
	// attention: below, use w = width-1, h = height-1
	// a circle steps: a = (w+h)*2, steps %= side4
	// case:
	// 1. steps = 0: (0,0), if init, face 'east', or face 'south'
	// 2. 1 <= steps <= side1: (steps, 0), face 'east'
	// 3. side1 < steps <= side2: (side1, steps-side1), face 'north'
	// 4. side2 < steps <= side3: ( w-(steps-side2) , h), face 'west'
	// 5. side3 < steps < side4: (0, h-(steps-side3) ), face 'south'
}

const (
	north = "North"
	south = "South"
	west  = "West"
	east  = "East"
)

func Constructor(width int, height int) Robot {
	w := width - 1
	h := height - 1

	return Robot{
		w: w,
		h: h,

		side1: w,
		side2: w + h,
		side3: w*2 + h,
		side4: (w + h) * 2,

		init: true,
	}
}

func (r *Robot) Step(num int) {
	r.steps = (r.steps + num) % r.side4
	r.init = false
}

func (r *Robot) GetPos() []int {
	res := []int{0, 0}

	switch {
	case r.steps == 0:
		// do nothing
	case 1 <= r.steps && r.steps <= r.side1:
		res[0] = r.steps
	case r.side1 < r.steps && r.steps <= r.side2:
		res[0] = r.side1
		res[1] = r.steps - r.side1
	case r.side2 < r.steps && r.steps <= r.side3:
		res[0] = r.side3 - r.steps
		res[1] = r.h
	case r.side3 < r.steps && r.steps < r.side4:
		res[1] = r.side4 - r.steps
	default:
		// testdata case
	}

	return res
}

func (r *Robot) GetDir() string {
	res := ""

	switch {
	case r.steps == 0:
		if r.init {
			res = east
		} else {
			res = south
		}
	case 1 <= r.steps && r.steps <= r.side1:
		res = east
	case r.side1 < r.steps && r.steps <= r.side2:
		res = north
	case r.side2 < r.steps && r.steps <= r.side3:
		res = west
	case r.side3 < r.steps && r.steps < r.side4:
		res = south
	default:
		// testdata case
	}

	return res
}
