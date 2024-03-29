package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"github.com/docktermj/go-xyzzy-grpc/g2diagnosticserver"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var (
	port = flag.Int("port", 50051, "The server port")
)

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	// Configure the "log" standard library.

	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)
	logger.SetLevel(logger.LevelTrace)

	// Parse input options.

	flag.Parse()

	// Set up socket listener.

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}

	// Create server.

	grpcServer := grpc.NewServer()
	pb.RegisterG2DiagnosticServer(grpcServer, &g2diagnosticserver.G2DiagnosticServer{})
	logger.Infof("Server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}
