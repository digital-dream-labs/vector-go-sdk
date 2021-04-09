package vector

import (
	"context"
	"flag"
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

// NewEP returns either a vector struct for escape pod vector, or an error on failure
func NewEP(opts ...Option) (*Vector, error) {
	cfg := options{}

	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.target == "" {
		var host = flag.String("host", "", "Vector's IP address")
		flag.Parse()
		if *host == "" {
			return nil, fmt.Errorf("please set the BOT_TARGET env variable, or use the -host argument and set it to your robots IP address")
		}
		cfg.target = fmt.Sprintf("%s:443", *host)
	}
	if cfg.target == "" {
		return nil, fmt.Errorf("configuration options missing")
	}

	c, err := client.New(
		client.WithTarget(cfg.target),
		client.WithInsecureSkipVerify(),
	)
	if err != nil {
		return nil, err
	}
	if err := c.Connect(); err != nil {
		return nil, err
	}

	vc := vectorpb.NewExternalInterfaceClient(c.Conn())

	login, err := vc.UserAuthentication(context.Background(),
		&vectorpb.UserAuthenticationRequest{
			UserSessionId: []byte("anything1"),
			ClientName:    []byte("anything2"),
		},
	)
	if err != nil {
		return nil, err
	}

	return New(
		WithTarget(cfg.target),
		WithToken(string(login.ClientTokenGuid)),
	)
}
