package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"payment-service/pb"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) ProcessPayment(ctx context.Context, in *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Received payment for transaction %v amount %v", in.GetTransactionId(), in.GetAmount())
	return &pb.PaymentResponse{Success: true, Message: "Payment successful"}, nil
}

func main() {
	fmt.Println("Payment service started on port 9091")
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//http.HandleFunc("/pay", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("Payment request received")
	//
	//	// Simulate some business logic, e.g., processing payment details
	//	fmt.Println("Processing payment...")
	//
	//	// Respond that the payment was processed successfully
	//	w.WriteHeader(http.StatusOK)
	//	fmt.Fprintf(w, "Payment processed successfully")
	//})
	//
	//fmt.Println("Payment service started on port 9092")
	//if err := http.ListenAndServe(":9092", nil); err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}
}
