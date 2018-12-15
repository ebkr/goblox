package marketplace

import (
	"encoding/json"
	"strconv"
)

func (ref *Marketplace) GetProductInfo(assetId int) (map[string]interface{}, error) {
	read, err := ref.SendRequest("http://api.roblox.com/marketplace/productinfo?assetId="+strconv.Itoa(assetId), map[string]interface{}{})
	if err != nil {
		return map[string]interface{}{}, err
	}
	var dt map[string]interface{}
	jErr := json.Unmarshal([]byte(read), &dt)
	if jErr != nil {
		return map[string]interface{}{}, jErr
	}
	return dt, err
}
