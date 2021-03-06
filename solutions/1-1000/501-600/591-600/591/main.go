package mario

import (
	"strings"
)

func isValid(code string) bool {
	tag, content := splitTag(code)
	if len(tag) < 1 && len(content) < 1 {
		return false
	}

	if !isValidContent(content) {
		return false
	}

	return true
}

// splitTag simulate str is valid and distinguish str into 'tag name' and 'tag content',
// if 'tag name' is invalid or it can not distinguish successfully, it will return two nil string
func splitTag(str string) (string, string) {
	closeIndex := strings.Index(str, ">")

	// can not parse 'tag name start'
	if !(len(str) > 0 && str[0] == '<' && closeIndex >= 0) {
		return "", ""
	}

	// can not parse 'tag name end'
	length := closeIndex - 1
	if len(str) < 2*length+5 || str[len(str)-1-length-2:len(str)-1-length] != "</" || str[len(str)-1] != '>' {
		return "", ""
	}

	tagNameStart := str[1:closeIndex]
	tagNameEnd := str[len(str)-1-length : len(str)-1]

	if !isValidTagName(tagNameStart, tagNameEnd) {
		return "", ""
	}

	return tagNameStart, str[length+3 : len(str)-1-length-2]
}

// isValidTagName judge if two tag is equal and valid
func isValidTagName(start, end string) bool {
	if len(start) != len(end) {
		return false
	}

	validFlag := true
	for i, char := range start {
		if start[i] != end[i] || !('A' <= char && char <= 'Z') {
			validFlag = false
			break
		}
	}

	return validFlag
}

// isValidContent return if input string is a valid content
func isValidContent(content string) bool {
	return false
}
