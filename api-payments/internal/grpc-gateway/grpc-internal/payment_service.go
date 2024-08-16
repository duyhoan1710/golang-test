package grpc_internal

import (
	grpc_protogen "api-payments/internal/grpc-gateway/protogen/golang"
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	config "api-payments/config"
)

type IPaymentService interface {
	PaymentProcess(c context.Context, req *grpc_protogen.PaymentRequest) (*grpc_protogen.PaymentResponse, error)
}

type PaymentService struct {
	grpc_protogen.UnimplementedPaymentServer
}

func NewServer(grpcServer *grpc.Server) {
	paymentGrpc := &PaymentService{}
	grpc_protogen.RegisterPaymentServer(grpcServer, paymentGrpc)
}

func (paymentService *PaymentService) PaymentProcess(ctx context.Context, req *grpc_protogen.PaymentRequest) (*grpc_protogen.PaymentResponse, error) {
	// Implement your payment processing logic here
	log.Printf("Processing payment for order %s with state %s", req.OrderId, req.UserId)

	rand.Seed(time.Now().UnixNano())

	return &grpc_protogen.PaymentResponse{
		IsSuccess: rand.Intn(2) == 1,
	}, nil
}

func StartGRPCServer(env *config.Env) {
	lis, err := net.Listen("tcp", env.GRPCServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc_protogen.RegisterPaymentServer(grpcServer, &PaymentService{})

	log.Printf("gRPC server is running on port %s", env.GRPCServerAddress)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
