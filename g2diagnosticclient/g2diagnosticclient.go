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
func (client *G2diagnosticClientImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	initRequest := pb.InitRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		VerboseLogging: int32(verboseLogging),
	}
	_, err := client.G2DiagnosticGrpcClient.Init(ctx, &initRequest)
	return err
}
