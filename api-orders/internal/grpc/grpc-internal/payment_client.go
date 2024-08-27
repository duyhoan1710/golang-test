package grpc_internal

import (
	"context"
	"fmt"
	"log"
	"time"

	grpc_protogen "api-orders/internal/grpc/protogen/golang"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IPaymentGRPCClient interface {
	ProcessPayment(orderId string, userId string) bool
}

type PaymentGRPCClient struct {
	client grpc_protogen.PaymentClient
}

func NewPaymentGRPCClient(grpcServerAddress string) (IPaymentGRPCClient, error) {
	conn, err := grpc.NewClient(grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return &PaymentGRPCClient{
		client: grpc_protogen.NewPaymentClient(conn),
	}, nil
}

func (s *PaymentGRPCClient) ProcessPayment(orderId string, userId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := s.client.PaymentProcess(ctx, &grpc_protogen.PaymentRequest{
		UserId:  userId,
		OrderId: orderId,
	})
	if err != nil {
		log.Fatalf("Failed to process payment: %v", err)
	}

	log.Printf("Payment response: %v", resp.IsSuccess)

	return resp.IsSuccess
}
