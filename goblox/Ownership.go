package goblox

import "log"

type Ownership struct {
	APIRequest
}

func (ref *Ownership) HasAsset() (map[string]interface{}, error) {
	log.Println("URL: " + "http://api.roblox.com/ownership/hasasset?userId=" + ref.data["userId"] + "&assetId=" + ref.data["assetId"])
	return ref.makeRequest("http://api.roblox.com/ownership/hasasset?userId=" + ref.data["userId"] + "&assetId=" + ref.data["assetId"])
}
