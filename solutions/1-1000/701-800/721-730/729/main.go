package mario

type MyCalendar struct {
	interval []item
}

type item struct {
	s int
	e int
}

func Constructor() MyCalendar {
	return MyCalendar{
		interval: make([]item, 0),
	}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	index := 0
	for index < len(mc.interval) && mc.interval[index].s <= start {
		index++
	}

	exit := false
	isValid := false

	if index == 0 {
		exit = true

		if len(mc.interval) < 1 || end < mc.interval[0].s {
			isValid = true
			mc.interval = append([]item{{s: start, e: end}}, mc.interval...)
		} else if end == mc.interval[0].s {
			isValid = true
			mc.interval[0].s = start
		} else {
			isValid = false
		}
	} else if index == len(mc.interval) {
		exit = true

		if start < mc.interval[index-1].e {
			isValid = false
		} else if start == mc.interval[index-1].e {
			isValid = true
			mc.interval[index-1].e = end
		} else {
			isValid = true
			mc.interval = append(mc.interval, item{s: start, e: end})
		}
	} else if mc.interval[index-1].e > start || end > mc.interval[index].s {
		exit = true
		isValid = false
	}

	if exit {
		return isValid
	}

	if mc.interval[index-1].e == start && end == mc.interval[index].s {
		newInterval := make([]item, 0, len(mc.interval)-1)
		newInterval = append(newInterval, mc.interval[:index-1]...)
		newInterval = append(newInterval, item{s: mc.interval[index-1].s, e: mc.interval[index].e})
		newInterval = append(newInterval, mc.interval[index+1:]...)

		mc.interval = newInterval
	} else if mc.interval[index-1].e == start {
		mc.interval[index-1].e = end
	} else if end == mc.interval[index].s {
		mc.interval[index].s = start
	} else {
		newInterval := make([]item, 0, len(mc.interval)+1)
		newInterval = append(newInterval, mc.interval[:index]...)
		newInterval = append(newInterval, item{s: start, e: end})
		newInterval = append(newInterval, mc.interval[index:]...)

		mc.interval = newInterval
	}

	return true
}
