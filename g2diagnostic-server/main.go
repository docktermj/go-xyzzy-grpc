package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"

	"github.com/docktermj/g2-sdk-go/g2diagnostic"
	//	"github.com/docktermj/g2-sdk-go/g2engine"
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	//	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"github.com/docktermj/go-xyzzy-helpers/logmessage"
	"google.golang.org/grpc"
)

const MessageIdFormat = "senzing-6023%04d"

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedG2DiagnosticServer
}

func traceEnter(messageNumber int, request interface{}) {
	if logger.IsTrace() {

		// Get calling function name.

		pc, _, _, _ := runtime.Caller(1)
		functionName := runtime.FuncForPC(pc).Name()

		// Simplify request for logging.

		var jsonString map[string]string
		requestAsJson, _ := json.Marshal(request)
		json.Unmarshal([]byte(requestAsJson), &jsonString)

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

		pc, _, _, _ := runtime.Caller(1)
		functionName := runtime.FuncForPC(pc).Name()

		// Simplify response for logging.

		var jsonString map[string]string
		responseAsJson, _ := json.Marshal(response)
		json.Unmarshal([]byte(responseAsJson), &jsonString)

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

func (s *server) Init(ctx context.Context, request *pb.InitRequest) (*pb.Empty, error) {
	traceEnter(5055, request)
	g2diagnostic := g2diagnostic.G2diagnosticImpl{}
	err := g2diagnostic.Init(ctx, request.ModuleName, request.IniParams, int(request.VerboseLogging))
	response := pb.Empty{}
	traceExit(5056, response)
	return &response, err
}

func main() {

	// Configure the "log" standard library.

	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)
	logger.SetLevel(logger.LevelTrace)

	// Parse input options.

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create server.

	s := grpc.NewServer()
	pb.RegisterG2DiagnosticServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
