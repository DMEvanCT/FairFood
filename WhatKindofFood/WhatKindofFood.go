package WhatKindofFood

import (
	"github.com/DMEvanCT/GoBase/Database"
	"log"
)


func GetFairFood( FoodType string , stream FoodInMyBelly_GetFoodServer) error {
	var FoodName string
	var FoodLocation string
	var FoodDesc string
	var FoodPrice int32
	var Vendor string



	db := Database.DatabaseInitAllHost("/etc/dm", "FileService", "FileService.username", "FileService.password", "")

	// Start the connection the connection to the db
	fetchmesomefood, err := db.Begin( )
	if err != nil {
		log.Fatal("There was an error connecting to the database")
	}

	defer fetchmesomefood.Rollback()

	fmsf, err := fetchmesomefood.Query("SELECT FoodName, FoodLocation, FoodDesc, FoodType, FoodPrice, Vendor FROM fair.fairfood where foodtype = " + FoodType )
	if err != nil {
		log.Println("There was an error with the query ", err )
	}

	defer fmsf.Close()

	for fmsf.Next() {
		err := fmsf.Scan(&FoodName, &FoodLocation, &FoodDesc, &FoodType, &FoodPrice, &Vendor)
		if err != nil {
			log.Println("There was an error with your query during the scan", err)
		}
		err = stream.Send(&FoodTypeResp{
			Foodname:             FoodName,
			Foodprice:            FoodPrice,
			Fooddesc:             FoodDesc,
			Vendor:               Vendor,
			Foodloc:              FoodLocation,
			Foodtype:             FoodType,
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
