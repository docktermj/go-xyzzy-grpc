/*
 *
 */

// Package main implements a client for the service.
package g2diagnosticclient

import (
	"context"
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
)

// Init initializes the Senzing G2diagnosis.
func (client *G2diagnosticClient) CheckDBPerf(ctx context.Context, secondsToRun int) (string, error) {
	request := pb.CheckDBPerfRequest{
		SecondsToRun: int32(secondsToRun),
	}
	response, err := client.G2DiagnosticGrpcClient.CheckDBPerf(ctx, &request)
	result := response.GetResult()
	return result, err
}

// Init initializes the Senzing G2diagnosis.
func (client *G2diagnosticClient) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	request := pb.InitRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		VerboseLogging: int32(verboseLogging),
	}
	_, err := client.G2DiagnosticGrpcClient.Init(ctx, &request)
	return err
}
