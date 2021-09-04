package mario

// border is valid
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	list := [][]int{{sr, sc}}
	oldColor := image[sr][sc]

	for len(list) > 0 {
		point := list[0]
		list = list[1:]

		image[point[0]][point[1]] = -1
		if point[0]-1 >= 0 && image[point[0]-1][point[1]] == oldColor {
			list = append(list, []int{point[0] - 1, point[1]})
		}
		if point[0]+1 < len(image) && image[point[0]+1][point[1]] == oldColor {
			list = append(list, []int{point[0] + 1, point[1]})
		}
		if point[1]-1 >= 0 && image[point[0]][point[1]-1] == oldColor {
			list = append(list, []int{point[0], point[1] - 1})
		}
		if point[1]+1 < len(image[0]) && image[point[0]][point[1]+1] == oldColor {
			list = append(list, []int{point[0], point[1] + 1})
		}
	}

	for i := range image {
		for j := range image[i] {
			if image[i][j] == -1 {
				image[i][j] = newColor
			}
		}
	}

	return image
}
