package goblox

import (
	"bytes"
	"errors"
	"log"
	"net/http"
)

// Authentication : Struct to control login/logout, and store token
type Authentication struct {
	token    string
	username string
	password string
}

// SetCredentials : Set login details
func (ref *Authentication) SetCredentials(username, password string) {
	ref.username = username
	ref.password = password
}

// Login : Get CSRF, and then attempt login
func (ref *Authentication) Login() (bool, error) {
	// Logout to get CSRF token
	ref.Logout()

	// Attempt login
	jsonData := []byte("{\"ctype\":\"Username\", \"cvalue\": \"" + ref.username + "\", \"password\":\"" + ref.password + "\"}")
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://auth.roblox.com/v2/login", bytes.NewBuffer(jsonData))
	req.Header.Set("X-CSRF-TOKEN", ref.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err == nil {
		jsonFactory := ResponseAsJSON{}
		jsonFactory.SetResponse(res)
		json := jsonFactory.GetJSON()
		if json["errors"] == nil {
			if json["user"] != nil && json["user"].(map[string]interface{})["id"] != nil {
				return true, err
			} else {
				return false, errors.New("Could not find user field, or user > id field")
			}
		}
		res.Body.Close()
	}
	return false, err
}

// Logout : Logout,
func (ref *Authentication) Logout() error {
	jsonData := []byte("")
	res, err := http.Post("https://api.roblox.com/sign-out/v1", "application/json", bytes.NewBuffer(jsonData))
	res.Header.Set("Accept", "application/json")
	res.Header.Set("Content-Type", "application/json")
	if res.Header["X-Csrf-Token"] != nil {
		ref.token = res.Header["X-Csrf-Token"][0]
	}
	if err == nil {
		jsonFactory := ResponseAsJSON{}
		jsonFactory.SetResponse(res)
		json := jsonFactory.GetJSON()
		for k, v := range json {
			log.Println(k)
			switch v.(type) {
			case []interface{}:
				for _, k2 := range v.([]interface{}) {
					log.Println(k2)
				}
				break
			case map[string]interface{}:
				for k2, v2 := range v.(map[string]interface{}) {
					log.Println(k2 + ":" + v2.(string))
				}
			}
		}
		res.Body.Close()
	}
	return err
}

// GetCSRFToken : Return the X-CSRF-TOKEN associated with the login
func (ref *Authentication) GetCSRFToken() string {
	return ref.token
}

// SetCSRFToken : Return the X-CSRF-TOKEN associated with the login
func (ref *Authentication) SetCSRFToken(token string) {
	ref.token = token
}
