package main

import (
	"context"
	"fmt"
	"log"
	"myproject/protos/student"

	"net"

	"google.golang.org/grpc"
)

type server struct {
	student.UnimplementedStudentServiceServer
}

func (s *server) GetStudentInfo(ctx context.Context, in *student.StudentRequest) (*student.StudentResponse, error) {
	return &student.StudentResponse{Id: in.Id, Name: in.Name, Age: 20}, nil
}

func main() {
	fmt.Println("server is listening on port: 50051")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	student.RegisterStudentServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
