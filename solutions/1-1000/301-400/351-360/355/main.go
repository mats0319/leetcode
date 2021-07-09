package mario

type Twitter struct {
	tweets     *listNode     // pre node
	subscribes map[int][]int // follower - subscribe ids
}

type listNode struct {
	tweetID int
	userID  int
	next    *listNode
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		tweets:     &listNode{},
		subscribes: make(map[int][]int, 0),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets.addToHead(tweetId, userId)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    relations := this.subscribes[userId]
    relations = append(relations, userId)

    return this.tweets.getRelatedNode(relations, 10)
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
    subscribeIDs := this.subscribes[followerId]
    subscribeIDs = append(subscribeIDs, followeeId)

    this.subscribes[followerId] = subscribeIDs
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
    subscribeIDs := this.subscribes[followerId]

    for i := range subscribeIDs {
        if subscribeIDs[i] == followeeId {
            subscribeIDs = append(subscribeIDs[:i], subscribeIDs[i+1:]...)
            break
        }
    }

    this.subscribes[followerId] = subscribeIDs
}

func (ln *listNode) addToHead(tweetID, userID int) {
	ln.next = &listNode{
		tweetID: tweetID,
		userID:  userID,
		next:    ln.next,
	}
}

func (ln *listNode) getRelatedNode(relations []int, max int) []int {
	res := make([]int, 0, max)
	for p := ln.next; p != nil; p = p.next {
		if !contains(relations, p.userID) {
			continue
		}

		res = append(res, p.tweetID)

		if len(res) >= max {
			break
		}
	}

	return res
}

func contains(arr []int, item int) bool {
	isExist := false
	for i := range arr {
		if arr[i] == item {
			isExist = true
			break
		}
	}

	return isExist
}
