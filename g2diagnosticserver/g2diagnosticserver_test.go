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

func testError(test *testing.T, ctx context.Context, err error) {
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
	_, err := getTestObject(ctx)
	testError(test, ctx, err)
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
	testError(test, ctx, err)
	test.Log("Actual:", actual)
}

func TestClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)
	request := pb.ClearLastExceptionRequest{}
	actual, err := g2diagnosticServer.ClearLastException(ctx, &request)
	testError(test, ctx, err)
	test.Log("Actual:", actual)
}

func TestDestroy(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)
	request := pb.DestroyRequest{}
	actual, err := g2diagnosticServer.Destroy(ctx, &request)
	testError(test, ctx, err)
	test.Log("Actual:", actual)
}

func TestEntityListBySize(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)

	entitySize := int32(10)
	request := pb.GetEntityListBySizeRequest{
		EntitySize: entitySize,
	}
	actual, err := g2diagnosticServer.GetEntityListBySize(ctx, &request)
	testError(test, ctx, err)
	test.Log("Actual:", actual)

	entityListBySizeHandle := actual.Result
	request2 := pb.FetchNextEntityBySizeRequest{
		EntityListBySizeHandle: entityListBySizeHandle,
	}
	actual2, err2 := g2diagnosticServer.FetchNextEntityBySize(ctx, &request2)
	testError(test, ctx, err2)
	test.Log("Actual:", actual2)

	request3 := pb.CloseEntityListBySizeRequest{
		EntityListBySizeHandle: entityListBySizeHandle,
	}
	actual3, err3 := g2diagnosticServer.CloseEntityListBySize(ctx, &request3)
	testError(test, ctx, err3)
	test.Log("Actual:", actual3)
}

func TestFindEntitiesByFeatureIDs(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)
	features := "{\"ENTITY_ID\":1,\"LIB_FEAT_IDS\":[1,3,4]}"
	request := pb.FindEntitiesByFeatureIDsRequest{
		Features: features,
	}
	actual, err := g2diagnosticServer.FindEntitiesByFeatureIDs(ctx, &request)
	testError(test, ctx, err)
	test.Log("Actual:", actual)
}

func TestGetAvailableMemoryy(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)
	request := pb.GetAvailableMemoryRequest{}
	actual, err := g2diagnosticServer.GetAvailableMemory(ctx, &request)
	testError(test, ctx, err)
	test.Log("Actual:", actual)
}

func TestGetDataSourceCounts(test *testing.T) {
	ctx := context.TODO()
	g2diagnosticServer, _ := getTestObject(ctx)
	request := pb.GetDataSourceCountsRequest{}
	actual, err := g2diagnosticServer.GetDataSourceCounts(ctx, &request)
	testError(test, ctx, err)
	test.Log("Actual:", actual)
}
