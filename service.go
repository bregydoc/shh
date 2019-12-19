package shh

import (

	"log"
	"net"
	"strings"

	"github.com/bregydoc/shh/proto"
	"google.golang.org/grpc"
)

func (w *Wizard) runRPCService() error {
	port := w.rpcPort
	if !strings.HasPrefix(w.rpcPort, ":") {
		port = ":" + port
	}
	
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	// TODO: Improve the grpc security
	grpcServer := grpc.NewServer()
	proto.RegisterSHHServer(grpcServer, w)
	return grpcServer.Serve(lis)
}
