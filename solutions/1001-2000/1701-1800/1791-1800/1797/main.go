package mario

type AuthenticationManager struct {
	tokenMap   map[string]int // token - expire time
	expiration int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		tokenMap:   make(map[string]int),
		expiration: timeToLive,
	}
}

// Generate all 'token id' are unique
func (am *AuthenticationManager) Generate(tokenId string, currentTime int) {
	am.tokenMap[tokenId] = currentTime + am.expiration
}

func (am *AuthenticationManager) Renew(tokenId string, currentTime int) {
	expireTime, ok := am.tokenMap[tokenId]
	if !ok || expireTime <= currentTime {
		return
	}

	am.tokenMap[tokenId] = currentTime + am.expiration
}

func (am *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for _, expireTime := range am.tokenMap {
		if expireTime > currentTime {
			count++
		}
	}

	return count
}
