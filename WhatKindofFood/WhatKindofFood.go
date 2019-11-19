package WhatKindofFood

import (
	"log"
	"github.com/DMEvanCT/GoBase/Database"
)


func GetFairFood( ftype string, stream FoodInMyBelly_GetFoodServer) error {
	var FoodName string
	var FoodLocation string
	var FoodDesc string
	var FoodPrice float32
	var VendorName string


	db := Database.DBInitSSLHostConfv2("/etc/dm", "FoodService", "FoodService.username", "FoodService.password", "FoodService.host", "FoodService.port")

	// Start the connection the connection to the db
	fetchmesomefood, err := db.Begin( )
	if err != nil {
		log.Fatal("There was an error connecting to the database", err)
	}

	defer fetchmesomefood.Rollback()

	fmsf, err := fetchmesomefood.Query("SELECT FoodName, FoodLocation, FoodDesc, FoodPrice, VendorName FROM fair.fairfood WHERE FoodType = " +"'" + ftype + "'")
	if err != nil {
		log.Println("There was an error with the query ", err )
	}

	defer fmsf.Close()

	for fmsf.Next() {
		err := fmsf.Scan(&FoodName, &FoodLocation, &FoodDesc, &FoodPrice, &VendorName)
		if err != nil {
			log.Println("There was an error with your query during the scan", err)
		}
		err = stream.Send(&FoodTypeResp{
			Foodname:             FoodName,
			Foodprice:            FoodPrice,
			Fooddesc:             FoodDesc,
			Vendor:               VendorName,
			Foodloc:              FoodLocation,
			Foodtype:             ftype,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		})
		if err != nil {
			log.Printf("There was an error sending the stream: %v", err)
		}

	}
	return nil
}
