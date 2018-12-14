package goblox

import (
	"bytes"
	"net/http"
)

type IRequest interface {
	SetData(map[string]string)
	SetPayload(map[string]string)
	SetCookie(string)
	SetRequestType(protocolEnum)
	SetHeader(map[string]string)
	makeRequest(url string) (map[string]interface{}, error)
}

type APIRequest struct {
	data        map[string]string
	payload     map[string]string
	header      map[string]string
	cookie      string
	requestType protocolEnum
}

// SetID : Set the ID for the call
func (ref *APIRequest) SetData(data map[string]string) {
	ref.data = data
}

// SetPayload : Set the Payload for the call
func (ref *APIRequest) SetPayload(payload map[string]string) {
	ref.payload = payload
}

// SetCookie : Set the Cookie for the call
func (ref *APIRequest) SetCookie(cookie string) {
	ref.cookie = cookie
}

// SetRequestType : Set the RequestType for the call
func (ref *APIRequest) SetRequestType(req protocolEnum) {
	ref.requestType = req
}

// SetHeader : Set the RequestType for the call
func (ref *APIRequest) SetHeader(header map[string]string) {
	ref.header = header
}

// makeRequest : Make a HTTP request
func (ref *APIRequest) makeRequest(url string) (map[string]interface{}, error) {
	client := http.Client{}
	var req *http.Request
	var err error
	var str = "{"
	for k, v := range ref.data {
		str += "\"" + k + "\":\"" + v + "\","
	}
	if len(str) > 1 {
		str = str[:len(str)-1]
	}
	str += "}"
	if ref.requestType == GETREQUEST {
		_req, _err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(str)))
		req = _req
		err = _err
	} else if ref.requestType == POSTREQUEST {
		_req, _err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(str)))
		req = _req
		err = _err
	}
	for k, v := range ref.header {
		req.Header.Set(k, v)
	}
	req.Cookie(ref.cookie)
	res, err := client.Do(req)
	if err == nil {
		jsonFactory := ResponseAsJSON{}
		jsonFactory.SetResponse(res)
		json := jsonFactory.GetJSON()
		res.Body.Close()
		return json, err
	}
	return map[string]interface{}{}, err
}
