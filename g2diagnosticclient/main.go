/*
 *
 */

// Package main implements a client for the service.
package g2diagnosticclient

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6025%04d"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2diagnosticClientImpl struct {
	G2DiagnosticClient G2DiagnosticClient
}

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type G2diagnosticClient interface {
	//	CheckDBPerf(ctx context.Context, secondsToRun int) (string, error)
	//	ClearLastException(ctx context.Context) error
	//	CloseEntityListBySize(ctx context.Context, entityListBySizeHandle interface{}) error
	//	Destroy(ctx context.Context) error
	//	FetchNextEntityBySize(ctx context.Context, entityListBySizeHandle interface{}) (string, error)
	//	FindEntitiesByFeatureIDs(ctx context.Context, features string) (string, error)
	//	GetAvailableMemory(ctx context.Context) (int64, error)
	//	GetDataSourceCounts(ctx context.Context) (string, error)
	//	GetDBInfo(ctx context.Context) (string, error)
	//	GetEntityDetails(ctx context.Context, entityID int64, includeInternalFeatures int) (string, error)
	//	GetEntityListBySize(ctx context.Context, entitySize int) (interface{}, error)
	//	GetEntityResume(ctx context.Context, entityID int64) (string, error)
	//	GetEntitySizeBreakdown(ctx context.Context, minimumEntitySize int, includeInternalFeatures int) (string, error)
	//	GetFeature(ctx context.Context, libFeatID int64) (string, error)
	//	GetGenericFeatures(ctx context.Context, featureType string, maximumEstimatedCount int) (string, error)
	//	GetLastException(ctx context.Context) (string, error)
	//	GetLastExceptionCode(ctx context.Context) (string, error)
	//	GetLogicalCores(ctx context.Context) (int, error)
	//	GetMappingStatistics(ctx context.Context, includeInternalFeatures int) (string, error)
	//	GetPhysicalCores(ctx context.Context) (int, error)
	//	GetRelationshipDetails(ctx context.Context, relationshipID int64, includeInternalFeatures int) (string, error)
	//	GetResolutionStatistics(ctx context.Context) (string, error)
	//	GetTotalSystemMemory(ctx context.Context) (int64, error)
	Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error
	//	InitWithConfigID(ctx context.Context, moduleName string, iniParams string, initConfigID int64, verboseLogging int) error
	//	Reinit(ctx context.Context, initConfigID int64) error
}

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
	g2diagnosticClient := pb.NewG2DiagnosticClient(grpcConnection)

	// Create a context.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create request parameters.

	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		log.Fatalf("Could not build Configuration JSON: %v", jsonErr)
	}

	initRequest := pb.InitRequest{
		ModuleName:     "Test module name",
		IniParams:      iniParams,
		VerboseLogging: int32(0), // 0 for no Senzing logging; 1 for logging
	}

	// Contact the server and print out its response.

	result, err := g2diagnosticClient.Init(ctx, &initRequest)
	if err != nil {
		logger.Fatalf("could not Init: %v", err)
	}
	logger.Infof("Result: %v", result)
}
