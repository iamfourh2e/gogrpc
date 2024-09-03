package main

import (
	"net"
	"seyhadot/gogrpc/pb"
	"seyhadot/gogrpc/repo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {

	grpcServer := grpc.NewServer()

	pb.RegisterEchoServer(grpcServer, repo.NewEchoRepo())
	pb.RegisterProductServiceServer(grpcServer, repo.NewProductRepo())
	//run the server
	PORT := ":50051"
	if l, err := net.Listen("tcp", PORT); err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	} else {
		// the gRPC server
		if err := grpcServer.Serve(l); err != nil {
			grpclog.Fatal("unable to start server", err)
		}
	}

}
