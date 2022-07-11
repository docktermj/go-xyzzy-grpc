/*
 *
 */

// Package main implements a client for the service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"github.com/docktermj/go-xyzzy-grpc/g2diagnosticclient"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

var (
	grpcAddress = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {

	// Configure the "log" standard library.

	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)
	logger.SetLevel(logger.LevelInfo)

	// Quick-and-dirty command line parameters. (Replace with Viper)

	flag.Parse()

	// Set up a connection and client to the gRPC server.

	grpcConnection, err := grpc.Dial(*grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer grpcConnection.Close()
	g2diagnosticGrpcClient := pb.NewG2DiagnosticClient(grpcConnection)

	g2diagnosticClient := g2diagnosticclient.G2diagnosticClientImpl{
		G2DiagnosticGrpcClient: g2diagnosticGrpcClient,
	}

	// Create request parameters.

	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		log.Fatalf("Could not build Configuration JSON: %v", jsonErr)
	}

	// Create a context.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call.

	err = g2diagnosticClient.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		logger.Fatalf("could not Init: %v", err)
	}
}
