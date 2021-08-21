package vector

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/digital-dream-labs/hugh/grpc/client"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
)

// Vector is the struct containing info about Vector
type Vector struct {
	Conn vectorpb.ExternalInterfaceClient
}

// New returns either a vector struct, or an error on failure
func New(opts ...Option) (*Vector, error) {

	cfg := options{}

	homedir, _ := os.UserHomeDir()
	dirname := filepath.Join(homedir, ".anki_vector", "sdk_config.ini")

	if initData, _ := ini.Load(dirname); initData != nil {
		sec, _ := initData.GetSection(ini.DefaultSection)
		sec.MapTo(&cfg)
	}

	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.Target == "" || cfg.Token == "" {
		return nil, fmt.Errorf("configuration options missing")
	}

	c, err := client.New(
		client.WithTarget(cfg.Target),
		client.WithInsecureSkipVerify(),
		client.WithDialopts(
			grpc.WithPerRPCCredentials(
				tokenAuth{
					token: cfg.Token,
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
