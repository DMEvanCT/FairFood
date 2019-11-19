package main

import (
	"github.com/DMEvanCT/FairFood/WhatKindofFood"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
port = ":50051"
)

type server struct {}

func (s server) GetFood(req *WhatKindofFood.FoodTypeRequest, stream WhatKindofFood.FoodInMyBelly_GetFoodServer) error {
	foodtype := req.Foodtype
	err := WhatKindofFood.GetFairFood(foodtype,stream)
	if err != nil {
		log.Println("There seems to be an error running GetFairFood", err )
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err )
	}
	s := grpc.NewServer()
	WhatKindofFood.RegisterFoodInMyBellyServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve %v", err)
	}

}
