package mario

type UndergroundSystem struct {
	checkInInfoMap map[int]*checkInInfo               // id - check in info
	runTimeInfoMap map[string]map[string]*runTimeInfo // start station - end station - run time info list
}

type checkInInfo struct {
	stationName string
	time        int
}

type runTimeInfo struct {
	timeSum      int
	sampleAmount int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		checkInInfoMap: make(map[int]*checkInInfo),
		runTimeInfoMap: make(map[string]map[string]*runTimeInfo),
	}
}

func (s *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	s.checkInInfoMap[id] = &checkInInfo{
		stationName: stationName,
		time:        t,
	}
}

func (s *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkInData, _ := s.checkInInfoMap[id]
	delete(s.checkInInfoMap, id)

	endStationMap, ok := s.runTimeInfoMap[checkInData.stationName]
	if !ok {
		endStationMap = make(map[string]*runTimeInfo)
		endStationMap[stationName] = &runTimeInfo{
			timeSum:      t - checkInData.time,
			sampleAmount: 1,
		}
	} else {
		runTimeData, ok := endStationMap[stationName]
		if !ok {
			endStationMap[stationName] = &runTimeInfo{
				timeSum:      t - checkInData.time,
				sampleAmount: 1,
			}
		} else {
			runTimeData.timeSum += t - checkInData.time
			runTimeData.sampleAmount++

			endStationMap[stationName] = runTimeData
		}
	}

	s.runTimeInfoMap[checkInData.stationName] = endStationMap
}

func (s *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	endStationMap, ok := s.runTimeInfoMap[startStation]
	if !ok {
		return -1
	}

	runTimeData, ok := endStationMap[endStation]
	if !ok {
		return -1
	}

	return float64(runTimeData.timeSum) / float64(runTimeData.sampleAmount)
}
