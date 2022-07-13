package g2diagnosticclient

import (
	"context"
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

var (
	grpcConnection     *grpc.ClientConn
	g2diagnosticClient pb.G2diagnosticClient
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) (G2diagnostic, error) {
	var err error = nil
	g2diagnostic := G2diagnosticImpl{}

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		return &g2diagnostic, jsonErr
	}

	err = g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	return &g2diagnostic, err
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
	g2engine, err := getTestObject(ctx)
	testError(test, ctx, g2engine, err)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestInit(test *testing.T) {
	ctx := context.TODO()
	g2diagnostic := &G2diagnosticClientImpl{}
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2diagnostic, jsonErr)
	err := g2diagnostic.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2diagnostic, err)
}
