package mario

type Encrypter struct {
	encryptMap map[byte]string // origin byte - encrypted string
	decryptMap map[string]int  // encrypted string - times
}

// Constructor
//    1 <= keys.length == values.length <= 26
//    values[i].length == 2
//    1 <= dictionary.length <= 100
//    1 <= dictionary[i].length <= 100
//    所有 keys[i] 和 dictionary[i] 互不相同
//    所有 word1[i] 都出现在 keys 中
//    word2.length 是偶数
func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	encryptMap := make(map[byte]string, len(keys))
	for i := range keys {
		encryptMap[keys[i]] = values[i]
	}

	decryptMap := make(map[string]int, len(dictionary))
	for i := range dictionary {
		byteSlice := make([]byte, 0, len(dictionary[i])*2)
		for j := 0; j < len(dictionary[i]) && contains(keys, dictionary[i][j]); j++ {
			byteSlice = append(byteSlice, encryptMap[dictionary[i][j]]...)
		}

		if len(byteSlice) == cap(byteSlice) {
			decryptMap[string(byteSlice)]++
		}
	}

	return Encrypter{
		encryptMap: encryptMap,
		decryptMap: decryptMap,
	}
}

func (e *Encrypter) Encrypt(word1 string) string {
	byteSlice := make([]byte, 0, len(word1)*2)
	for i := 0; i < len(word1); i++ {
		byteSlice = append(byteSlice, e.encryptMap[word1[i]]...)
	}

	return string(byteSlice)
}

func (e *Encrypter) Decrypt(word2 string) int {
	return e.decryptMap[word2]
}

func contains(arr []byte, v byte) bool {
	containsFlag := false
	for i := range arr {
		if arr[i] == v {
			containsFlag = true
			break
		}
	}

	return containsFlag
}
