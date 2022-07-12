package g2diagnosticserver

import (
	"context"
	"encoding/json"
	"runtime"
	"strconv"

	"github.com/docktermj/g2-sdk-go/g2diagnostic"
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/docktermj/go-xyzzy-helpers/logmessage"
)

// ----------------------------------------------------------------------------
// Internal methods - names begin with lower case
// ----------------------------------------------------------------------------

func getG2diagnostic() *g2diagnostic.G2diagnosticImpl {
	result := g2diagnostic.G2diagnosticImpl{}
	return &result
}

func traceEnter(messageNumber int, request interface{}) {
	if logger.IsTrace() {

		// Get calling function name.

		programCounter, file, line, _ := runtime.Caller(1)
		functionName := runtime.FuncForPC(programCounter).Name()

		// Assemble values to be logged.

		var jsonString map[string]string
		requestAsJson, _ := json.Marshal(request)
		json.Unmarshal([]byte(requestAsJson), &jsonString)
		jsonString["runtimeFile"] = file
		jsonString["runtimeFunction"] = functionName
		jsonString["runtimeLine"] = strconv.Itoa(line)

		// Construct message.

		message := logmessage.BuildMessageUsingMap(
			logger.BuildMessageId(MessageIdFormat, messageNumber),
			"TRACE",
			"Enter "+functionName,
			jsonString,
		)

		// Log message.

		logger.Trace(message)
	}
}

func traceExit(messageNumber int, response interface{}) {
	if logger.IsTrace() {

		// Get calling function name.

		programCounter, file, line, _ := runtime.Caller(1)
		functionName := runtime.FuncForPC(programCounter).Name()

		// Assemble values to be logged.

		var jsonString map[string]string
		responseAsJson, _ := json.Marshal(response)
		json.Unmarshal([]byte(responseAsJson), &jsonString)
		jsonString["runtimeFile"] = file
		jsonString["runtimeFunction"] = functionName
		jsonString["runtimeLine"] = strconv.Itoa(line)

		// Construct message.

		message := logmessage.BuildMessageUsingMap(
			logger.BuildMessageId(MessageIdFormat, messageNumber),
			"TRACE",
			"Exit "+functionName,
			jsonString,
		)

		// Log message.

		logger.Trace(message)
	}
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (server *G2DiagnosticServer) CheckDBPerf(ctx context.Context, request *pb.CheckDBPerfRequest) (*pb.CheckDBPerfResponse, error) {
	traceEnter(5100, request)
	g2diagnostic := getG2diagnostic()
	result, err := g2diagnostic.CheckDBPerf(ctx, int(request.SecondsToRun))
	response := pb.CheckDBPerfResponse{
		Result: result,
	}
	traceExit(5101, response)
	return &response, err
}

func (server *G2DiagnosticServer) ClearLastException(ctx context.Context, request *pb.ClearLastExceptionRequest) (*pb.ClearLastExceptionResponse, error) {
	traceEnter(5102, request)
	g2diagnostic := getG2diagnostic()
	err := g2diagnostic.ClearLastException(ctx)
	response := pb.ClearLastExceptionResponse{}
	traceExit(5103, response)
	return &response, err
}

func (server *G2DiagnosticServer) CloseEntityListBySize(ctx context.Context, request *pb.CloseEntityListBySizeRequest) (*pb.CloseEntityListBySizeResponse, error) {
	traceEnter(5198, request)
	g2diagnostic := getG2diagnostic()
	err := g2diagnostic.CloseEntityListBySize(ctx, request.EntityListBySizeHandle)
	response := pb.CloseEntityListBySizeResponse{}
	traceExit(5199, response)
	return &response, err
}

func (server *G2DiagnosticServer) Destroy(ctx context.Context, request *pb.DestroyRequest) (*pb.DestroyResponse, error) {
	traceEnter(5198, request)
	g2diagnostic := getG2diagnostic()
	err := g2diagnostic.Destroy(ctx)
	response := pb.DestroyResponse{}
	traceExit(5199, response)
	return &response, err
}

func (server *G2DiagnosticServer) FetchNextEntityBySize(ctx context.Context, request *pb.FetchNextEntityBySizeRequest) (*pb.FetchNextEntityBySizeResponse, error) {
	traceEnter(5198, request)
	g2diagnostic := getG2diagnostic()
	result, err := g2diagnostic.FetchNextEntityBySize(ctx, request.EntityListBySizeHandle)
	response := pb.FetchNextEntityBySizeResponse{
		Result: result,
	}
	traceExit(5199, response)
	return &response, err
}

func (server *G2DiagnosticServer) FindEntitiesByFeatureIDs(ctx context.Context, request *pb.FindEntitiesByFeatureIDsRequest) (*pb.FindEntitiesByFeatureIDsResponse, error) {
	traceEnter(5198, request)
	g2diagnostic := getG2diagnostic()
	result, err := g2diagnostic.FindEntitiesByFeatureIDs(ctx, request.Features)
	response := pb.FindEntitiesByFeatureIDsResponse{
		Result: result,
	}
	traceExit(5199, response)
	return &response, err
}

func (server *G2DiagnosticServer) GetAvailableMemory(ctx context.Context, request *pb.GetAvailableMemoryRequest) (*pb.GetAvailableMemoryResponse, error) {
	traceEnter(5198, request)
	g2diagnostic := getG2diagnostic()
	result, err := g2diagnostic.GetAvailableMemory(ctx)
	response := pb.GetAvailableMemoryResponse{
		Result: result,
	}
	traceExit(5199, response)
	return &response, err
}

func (server *G2DiagnosticServer) GetDataSourceCounts(ctx context.Context, request *pb.GetDataSourceCountsRequest) (*pb.GetDataSourceCountsResponse, error) {
	traceEnter(5198, request)
	g2diagnostic := getG2diagnostic()
	result, err := g2diagnostic.GetDataSourceCounts(ctx)
	response := pb.GetDataSourceCountsResponse{
		Result: result,
	}
	traceExit(5199, response)
	return &response, err
}

func (server *G2DiagnosticServer) Init(ctx context.Context, request *pb.InitRequest) (*pb.InitResponse, error) {
	traceEnter(5198, request)
	g2diagnostic := getG2diagnostic()
	err := g2diagnostic.Init(ctx, request.ModuleName, request.IniParams, int(request.VerboseLogging))
	response := pb.InitResponse{}
	traceExit(5156, response)
	return &response, err
}
