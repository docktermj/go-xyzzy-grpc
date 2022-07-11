/*
 *
 */

// Package main implements a client for the service.
package g2diagnosticclient

import (
	"context"

	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
	"github.com/docktermj/go-xyzzy-helpers/g2configuration"
	"github.com/docktermj/go-xyzzy-helpers/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Init initializes the Senzing G2diagnosis.
func (g2diagnosticclient *G2diagnosticclientImpl) Init(ctx context.Context, moduleName string, iniParams string, verboseLogging int) error {
	initRequest := pb.InitRequest{
		ModuleName:     moduleName,
		IniParams:      iniParams,
		VerboseLogging: int32(verboseLogging),
	}
	_, err := g2diagnosticClient.Init(ctx, &initRequest)
	return err
}
