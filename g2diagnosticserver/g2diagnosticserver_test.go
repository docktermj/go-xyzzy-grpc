package g2diagnosticserver

import (
	"context"
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"testing"

	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context) (G2DiagnosticServer, error) {
	var err error = nil
	g2diagnosticServer := G2DiagnosticServer{}

	moduleName := "Test module name"
	verboseLogging := 0 // 0 for no Senzing logging; 1 for logging
	iniParams, jsonErr := g2configuration.BuildSimpleSystemConfigurationJson("")
	if jsonErr != nil {
		return g2diagnosticServer, jsonErr
	}

	request := pb.InitRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		VerboseLogging: int32(verboseLogging),
	}
	_, err = g2diagnosticServer.Init(ctx, &request)
	return g2diagnosticServer, err
}

func testError(test *testing.T, ctx context.Context, g2diagnosticServer G2DiagnosticServer, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestGetObject(test *testing.T) {
	ctx := context.TODO()
	g2engine, err := getTestObject(ctx)
	testError(test, ctx, g2engine, err)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestCheckDBPerf(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)
	request := pb.CheckDBPerfRequest{
		SecondsToRun: int32(1),
	}
	actual, err := g2diagnosticServer.CheckDBPerf(ctx, &request)
	testError(test, ctx, g2diagnosticServer, err)
	test.Log("Actual:", actual)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)
	request := pb.Empty{}
	actual, err := g2diagnosticServer.ClearLastException(ctx, &request)
	testError(test, ctx, g2diagnosticServer, err)
	test.Log("Actual:", actual)
}
