package friends

func (ref *Friends) Follow(id int) (string, error) {
	return ref.SendRequest("http://api.roblox.com/user/follow", map[string]interface{}{
		"followedUserId": id,
	})
}
