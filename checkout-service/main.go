package main

import (
	"checkout-service/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCheckoutServiceServer
}

func (s *server) ProcessCheckout(ctx context.Context, in *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	log.Printf("Received: %v", in.GetOrderId())
	return &pb.CheckoutResponse{Confirmation: "Order processed"}, nil
}

func main() {
	//http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("Checkout request received")
	//
	//	// Simulate some business logic, e.g., processing Checkout details
	//	fmt.Println("Processing Checkout...")
	//
	//	// Respond that the Checkout was processed successfully
	//	w.WriteHeader(http.StatusOK)
	//	fmt.Fprintf(w, "Checkout processed successfully")
	//})
	//
	//fmt.Println("Checkout service started on port 9091")
	//if err := http.ListenAndServe(":9091", nil); err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}
	fmt.Println("Checkout now service started on port 9090")
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCheckoutServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
