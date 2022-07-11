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

func (server *G2DiagnosticServer) Init(ctx context.Context, request *pb.InitRequest) (*pb.Empty, error) {
	traceEnter(5055, request)
	g2diagnostic := g2diagnostic.G2diagnosticImpl{}
	err := g2diagnostic.Init(ctx, request.ModuleName, request.IniParams, int(request.VerboseLogging))
	response := pb.Empty{}
	traceExit(5056, response)
	return &response, err
}
