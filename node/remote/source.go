package remote

import (
	"context"

	"google.golang.org/grpc"

	"github.com/chain4energy/juno/v4/node"
)

var (
	_ node.Source = &Source{}
)

// Source implements the keeper.Source interface relying on a GRPC connection
type Source struct {
	Ctx          context.Context
	GrpcConn     *grpc.ClientConn
	Restendpoint string
}

// NewSource returns a new Source instance
func NewSource(config *GRPCConfig, restEndpoint string) (*Source, error) {
	return &Source{
		Ctx:          context.Background(),
		GrpcConn:     MustCreateGrpcConnection(config),
		Restendpoint: restEndpoint,
	}, nil
}

// Type implements keeper.Type
func (k Source) Type() string {
	return node.RemoteKeeper
}
