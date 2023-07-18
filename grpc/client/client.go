package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "com.ai.bff-purchase-order-inquiry/proto"
)

func GetOrder(ctx context.Context, id string) (*pb.GetOrderResponse, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)
	req := &pb.GetOrderRequest{
		Id: id,
	}

	resp, err := client.GetOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
