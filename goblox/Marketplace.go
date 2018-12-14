/*
	Created by: ebkr

	ID: Asset ID
	Payload: Data to send to the server

*/

package goblox

// Marketplace : API calls for Marketplace related calls
type Marketplace struct {
	APIRequest
}

/*
	Asset Method Implementation
*/

// GetProductInfo : API call to Marketplace/productinfo
func (ref *Marketplace) GetProductInfo() (map[string]interface{}, error) {
	return ref.makeRequest("http://api.roblox.com/marketplace/productinfo?assetId=" + ref.data["assetId"])
}

// GetProductInfo : API call to Marketplace/game-pass-product-info
func (ref *Marketplace) GetGamePassProductInfo() (map[string]interface{}, error) {
	return ref.makeRequest("http://api.roblox.com/marketplace/game-pass-product-info?gamePassId=" + ref.data["assetId"])
}
