package mario

// rooms.length >= 1
func canVisitAllRooms(rooms [][]int) bool {
    hasVisited := make([]bool, len(rooms))
    hasVisited[0] = true

    queue := make([]int, 0, len(rooms))
    queue = append(queue, 0)

    for i := 0; i < len(queue); i++ {
        room := queue[i]
        for j := 0; j < len(rooms[room]); j++ {
            key := rooms[room][j]
            if hasVisited[key] {
                continue
            }

            hasVisited[key] = true
            queue = append(queue, key)
        }
    }

    canVisit := true
    for i := range hasVisited {
        if !hasVisited[i] {
            canVisit = false
            break
        }
    }

    return canVisit
}
