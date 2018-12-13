package goblox

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ResponseAsJSON struct {
	response *http.Response
}

func (ref *ResponseAsJSON) SetResponse(response *http.Response) {
	ref.response = response
}

func (ref *ResponseAsJSON) GetJSON() map[string]interface{} {
	data, err := ioutil.ReadAll(ref.response.Body)
	if err != nil {
		return map[string]interface{}{}
	}
	var raw map[string]interface{}
	json.Unmarshal(data, &raw)
	return raw
}
