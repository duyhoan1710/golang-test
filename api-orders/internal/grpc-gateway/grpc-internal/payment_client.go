package grpc_internal

import (
	"context"
	"log"
	"time"

	"api-orders/config"
	grpc_protogen "api-orders/internal/grpc-gateway/protogen/golang"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ProcessPayment(env *config.Env, orderId string, userId string) bool {
	conn, err := grpc.NewClient(env.GRPC_PAYMENT_SERVER_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := grpc_protogen.NewPaymentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.PaymentProcess(ctx, &grpc_protogen.PaymentRequest{
		UserId:  userId,
		OrderId: orderId,
	})
	if err != nil {
		log.Fatalf("Failed to process payment: %v", err)
	}

	log.Printf("Payment response: %v", resp.IsSuccess)

	return resp.IsSuccess
}
