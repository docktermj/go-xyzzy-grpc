package g2diagnosticserver

import (
	pb "github.com/docktermj/go-xyzzy-grpc/g2diagnostic"
)

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const MessageIdFormat = "senzing-6023%04d"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// server is used to implement helloworld.GreeterServer.
type G2DiagnosticServer struct {
	pb.UnimplementedG2DiagnosticServer
}
