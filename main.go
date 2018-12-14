package main

import (
	"log"

	"github.com/ebkr/goblox/goblox"
)

func main() {
	mapData := map[string]string{"userId": "9706143", "assetId": "1192464705"}
	log.SetPrefix("> ")
	log.Println("Marketplace API: ")
	market := goblox.Marketplace{}
	market.SetRequestType(goblox.GETREQUEST)
	market.SetData(mapData)
	log.Println(market.GetProductInfo())

	log.Println("------------------------------")
	log.Println("------------------------------")
	log.Println("------------------------------")

	log.Println("Ownership API:")
	asset := goblox.Ownership{}
	asset.SetRequestType(goblox.GETREQUEST)
	asset.SetData(mapData)
	log.Println(asset.HasAsset())

	log.Println("------------------------------")
	log.Println("------------------------------")
	log.Println("------------------------------")
}
