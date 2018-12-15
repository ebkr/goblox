package network

import (
	"log"
	"testing"
)

// Test_main :
func Test_main(test *testing.T) {
	req := NetworkRequest{}
	req.New()

	req.SetContentType(APPJSON)
	req.SetRequestType(GET)
	log.Println(req.SendRequest("http://api.roblox.com/user/get-username-by-id?username=ebkrrbx", map[string]interface{}{
		"userId": 0,
	}))
}
