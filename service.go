package shh

import (
	"fmt"
	"log"
	"net"

	"github.com/bregydoc/shh/proto"
	"google.golang.org/grpc"
)

func (w *Wizard) runRPCService() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	// TODO: Improve the grpc security

	grpcServer := grpc.NewServer()
	proto.RegisterSHHServer(grpcServer, w)
	return grpcServer.Serve(lis)
}
