package backend

import "github.com/johanbrandhorst/grpcweb-boilerplate/proto/server"

// Backend should be used to implement the server interface
// exposed by the generated server proto.
type Backend struct {
}

// Ensure struct implements interface
var _ server.BackendServer = (*Backend)(nil)
