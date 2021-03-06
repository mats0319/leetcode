package mario

// regard it start with '/'
func simplifyPath(path string) string {
	stack := make([]string, 0)
	i := 0
	for i < len(path) {
		nextSlash := i + 1
		for nextSlash < len(path) && path[nextSlash] != '/' {
			nextSlash++
		}

		str := path[i+1 : nextSlash]
		switch str {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, str)
		}

		i = nextSlash
	}

	res := ""
	for _, s := range stack {
		res += "/" + s
	}

	if len(res) < 1 {
		res = "/"
	}

	return res
}
