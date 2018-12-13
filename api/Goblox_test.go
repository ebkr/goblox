package goblox

import (
	"log"
	"reflect"
	"testing"
)

// Test_SetEnum : Set protocolEnums, and check their values
func Test_SetEnum(test *testing.T) {
	log.Println("> Running Test: SetEnum")
	e := GETREQUEST
	if e != 0 {
		test.Errorf("Invalid enum: GETREQUEST")
	}
	e = POSTREQUEST
	if e != 1 {
		test.Errorf("Invalid enum: POSTREQUEST")
	}
}

// Test_GetProductInfo : Test API calls to retrieve product information
func Test_GetProductInfo(test *testing.T) {
	log.Println("> Running Test: GetProductInfo")
	asset := Asset{}
	asset.SetID(1584277735)
	asset.requestType = GETREQUEST
	response, err := asset.GetProductInfo()
	if err != nil {
		test.Errorf(err.Error())
	}
	if reflect.TypeOf(response["AssetId"]) != reflect.TypeOf(float64(0)) {
		test.Errorf("Wrong format")
	}
}
