package goblox

type User struct {
	APIRequest
}

// Follow : Follow a user
func (ref *User) Follow() (map[string]interface{}, error) {
	return ref.makeRequest("http://api.roblox.com/user/follow")
}

func (ref *User) GetFriendshipCount() (map[string]interface{}, error) {
	return ref.makeRequest("http://api.roblox.com/user/get-friendship-count?userId=" + ref.data["userId"])
}
