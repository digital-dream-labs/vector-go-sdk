package vector

import (
	"fmt"

	"github.com/digital-dream-labs/hugh/grpc/client"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
	"google.golang.org/grpc"
)

// Vector is the struct containing info about Vector
type Vector struct {
	Conn vectorpb.ExternalInterfaceClient
}

// New returns either a vector struct, or an error on failure
func New(opts ...Option) (*Vector, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.target == "" || cfg.token == "" {
		return nil, fmt.Errorf("configuration options missing")
	}

	c, err := client.New(
		client.WithTarget(cfg.target),
		client.WithInsecureSkipVerify(),
		client.WithDialopts(
			grpc.WithPerRPCCredentials(
				tokenAuth{
					token: cfg.token,
				},
			),
		),
	)
	if err != nil {
		return nil, err
	}
	if err := c.Connect(); err != nil {
		return nil, err
	}

	r := Vector{
		Conn: vectorpb.NewExternalInterfaceClient(c.Conn()),
	}

	return &r, nil

}
