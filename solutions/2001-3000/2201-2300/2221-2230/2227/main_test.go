package mario

import "testing"

func TestConstructor(t *testing.T) {
	ins := Constructor(
		[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
		[]string{"aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa", "aa"},
		[]string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"})

	expectedEncryptRes := []string{"aaaaaaaa"}

	encryptRes := make([]string, 0, 1)
	encryptRes = append(encryptRes, ins.Encrypt("abcd"))

	expectedDecryptRes := []int{8}

	decryptRes := make([]int, 0, 1)
	decryptRes = append(decryptRes, ins.Decrypt("aaaaaaaa"))

	if len(expectedEncryptRes) != len(encryptRes) || len(expectedDecryptRes) != len(decryptRes) {
		t.Log("unmatched amount")
		t.Fail()
	}

	for i := range expectedEncryptRes {
		if expectedEncryptRes[i] != encryptRes[i] {
			t.Logf("encrypt test failed on index: %d, want: %s, get: %s\n", i, expectedEncryptRes[i], encryptRes[i])
			t.Fail()
		}
	}

	for i := range expectedDecryptRes {
		if expectedDecryptRes[i] != decryptRes[i] {
			t.Logf("decrypt test failed on index: %d, want: %d, get: %d\n", i, expectedDecryptRes[i], decryptRes[i])
			t.Fail()
		}
	}
}

//["Encrypter","encrypt","decrypt"]
//[[["b"],["ca"],["aaa","cacbc","bbaba","bb"]],["bbb"],["cacaca"]]
func TestConstructor_Case2(t *testing.T) {
	ins := Constructor(
		[]byte{'b'},
		[]string{"ca"},
		[]string{"aaa", "cacbc", "bbaba", "bb"})

	expectedEncryptRes := []string{"cacaca"}

	encryptRes := make([]string, 0, 1)
	encryptRes = append(encryptRes, ins.Encrypt("bbb"))

	expectedDecryptRes := []int{0}

	decryptRes := make([]int, 0, 1)
	decryptRes = append(decryptRes, ins.Decrypt("cacaca"))

	if len(expectedEncryptRes) != len(encryptRes) || len(expectedDecryptRes) != len(decryptRes) {
		t.Log("unmatched amount")
		t.Fail()
	}

	for i := range expectedEncryptRes {
		if expectedEncryptRes[i] != encryptRes[i] {
			t.Logf("encrypt test failed on index: %d, want: %s, get: %s\n", i, expectedEncryptRes[i], encryptRes[i])
			t.Fail()
		}
	}

	for i := range expectedDecryptRes {
		if expectedDecryptRes[i] != decryptRes[i] {
			t.Logf("decrypt test failed on index: %d, want: %d, get: %d\n", i, expectedDecryptRes[i], decryptRes[i])
			t.Fail()
		}
	}
}
