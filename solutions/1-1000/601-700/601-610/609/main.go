package mario

import "strings"

func findDuplicate(paths []string) [][]string {
	fileNameMap := make(map[string][]string)

	for i := range paths {
		desp := strings.Split(paths[i], " ") // 0: root, >0: file name + content, e.g. "root 1.txt(abcd) 2.txt(abcd)"
		if len(desp) < 2 {
			continue // invalid input
		}

		for j := 1; j < len(desp); j++ {
			fileDesp := strings.Split(desp[j], "(") // 0: file name, 1: file content(with ')')
			if len(fileDesp) < 2 {
				continue // invalid input
			}

			newFileName := strings.Join([]string{desp[0], fileDesp[0]}, "/")
			fileNameList := fileNameMap[fileDesp[1]]
			if !contains(fileNameList, newFileName) {
				fileNameList = append(fileNameList, newFileName)
				fileNameMap[fileDesp[1]] = fileNameList
			}
		}
	}

	res := make([][]string, 0)
	for _, group := range fileNameMap {
		if len(group) < 2 {
			continue
		}

		res = append(res, group)
	}

	return res
}

func contains(slice []string, item string) bool {
	isContained := false
	for i := range slice {
		if slice[i] == item {
			isContained = true
			break
		}
	}

	return isContained
}
