package g2diagnosticclient

import (
	"context"
	"fmt"
	g2d "github.com/docktermj/g2-sdk-go/g2diagnostic"
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

var (
	grpcAddress    = "localhost:50051"
	grpcConnection *grpc.ClientConn
	g2diagnostic   g2d.G2diagnostic
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getGrpcConnection() *grpc.ClientConn {
	var err error
	if grpcConnection == nil {
		fmt.Println(">>>>>>>>>>>>>>> Getting connection")
		grpcConnection, err = grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Fatalf("Did not connect: %v", err)
		}
		//		defer grpcConnection.Close()
	}
	return grpcConnection
}

//func getG2diagnosticClient() G2diagnosticClient {
//	if g2diagnosticClient == (G2diagnosticClient{}) {
//		fmt.Println(">>>>>>>>>>>>>>> Getting G2diagnosticClient")
//
//
//	}
//	return g2diagnosticClient
//}

func getTestObject(ctx context.Context) g2d.G2diagnostic {
	if g2diagnostic == nil {

		fmt.Println(">>>>>>>>>>>>>>> Getting G2diagnosticClient")

		grpcConnection := getGrpcConnection()
		g2diagnostic = &G2diagnosticClient{
			G2DiagnosticGrpcClient: pb.NewG2DiagnosticClient(grpcConnection),
		}

		moduleName := "Test module name"
		verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
		iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			logger.Fatalf("Cannot construct system configuration: %v", jsonErr)
		}

		initErr := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			logger.Fatalf("Cannot Init: %v", initErr)
		}
	}

	return g2diagnostic
}

func testError(test *testing.T, ctx context.Context, g2diagnostic G2diagnostic, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		lastException, _ := g2diagnostic.GetLastException(ctx)
		assert.FailNow(test, lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

// Reference: https://medium.com/nerd-for-tech/setup-and-teardown-unit-test-in-go-bd6fa1b785cd
func setupSuite(test testing.TB) func(test testing.TB) {
	test.Log("setup suite")

	// Return a function to teardown the test
	return func(test testing.TB) {
		test.Log("teardown suite")
	}
}

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	getTestObject(ctx)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestCheckDBPerf(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	secondsToRun := 1
	actual, err := g2diagnostic.CheckDBPerf(ctx, secondsToRun)
	testError(test, ctx, g2diagnostic, err)
	test.Log("Actual:", actual)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	g2diagnostic.ClearLastException(ctx)
}

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := getTestObject(ctx)
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2diagnostic, jsonErr)
	err := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}
