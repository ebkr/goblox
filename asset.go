/*
	Created by: ebkr

	ID: Asset ID
	Payload: Data to send to the server

*/

package goblox

import (
	"errors"
	"net/http"
	"strconv"
)

// Asset : API calls for asset related data
type Asset struct {
	id          int
	payload     map[string]string
	cookie      string
	requestType protocolEnum
}

/*
	Asset Struct Setup
*/

// SetID : Set the ID for the call
func (ref *Asset) SetID(id int) {
	ref.id = id
}

// SetPayload : Set the Payload for the call
func (ref *Asset) SetPayload(payload map[string]string) {
	ref.payload = payload
}

// SetCookie : Set the Cookie for the call
func (ref *Asset) SetCookie(cookie string) {
	ref.cookie = cookie
}

// SetRequestType : Set the RequestType for the call
func (ref *Asset) SetRequestType(req protocolEnum) {
	ref.requestType = req
}

/*
	Asset Method Implementation
*/

func (ref *Asset) GetProductInfo() (map[string]interface{}, error) {
	if ref.requestType == GETREQUEST {
		res, err := http.Get("http://api.roblox.com/marketplace/productinfo?assetId=" + strconv.Itoa(ref.id))
		if err != nil {
			return map[string]interface{}{}, err
		}
		jsonFactory := ResponseAsJSON{}
		jsonFactory.SetResponse(res)
		json := jsonFactory.GetJSON()
		res.Body.Close()
		return json, err
	}
	return map[string]interface{}{}, errors.New("Invalid request type")
}
