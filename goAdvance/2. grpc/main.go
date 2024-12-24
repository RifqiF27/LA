package main

import (
	"context"
	pb "project/controllers/proto"
	"project/models"

	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type eventServiceServer struct {
	pb.UnimplementedEventServiceServer
}

func (s *eventServiceServer) GetEvents(ctx context.Context, _ *emptypb.Empty) (*pb.EventListResponse, error) {
	var events []*pb.Event
	for _, event := range models.Events {
		events = append(events, &pb.Event{
			Id:          int32(event.ID),
			Title:       event.Title,
			Description: event.Description,
			Date:        event.Date,
			Location:    event.Location,
		})
	}
	return &pb.EventListResponse{Events: events}, nil
}

func (s *eventServiceServer) GetEventDetail(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	for _, event := range models.Events {
		if int32(event.ID) == req.Id {
			return &pb.EventResponse{Event: &pb.Event{
				Id:          int32(event.ID),
				Title:       event.Title,
				Description: event.Description,
				Date:        event.Date,
				Location:    event.Location,
			}}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "event not found")
}

func main() {
	listener, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEventServiceServer(grpcServer, &eventServiceServer{})
	reflection.Register(grpcServer) // Enable reflection for testing with grpc_cli

	log.Println("gRPC server is running on :8086")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
