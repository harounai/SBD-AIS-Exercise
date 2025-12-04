package server

import (
	"context"
	"exc8/pb"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
	drinks []*pb.Drink
	orders map[int32]int32
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions

// GetDrinks returns the list of drinks
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.DrinkList, error) {
	return &pb.DrinkList{Drinks: s.drinks}, nil
}

// OrderDrink stores the ordered drinks in memory
func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*emptypb.Empty, error) {
	for _, item := range req.Items {
		s.orders[item.DrinkId] += item.Quantity
	}
	return &emptypb.Empty{}, nil
}

// GetOrders returns the total ordered drinks
func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.OrderTotals, error) {
	var totals []*pb.OrderItem
	for id, qty := range s.orders {
		totals = append(totals, &pb.OrderItem{
			DrinkId:  id,
			Quantity: qty,
		})
	}
	return &pb.OrderTotals{Totals: totals}, nil
}
