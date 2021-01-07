package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/digital-dream-labs/hugh/grpc/client"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

const (
	// These can be set to whatever you want
	sessionID  = "id01"
	clientName = "id02"
)

func main() {

	var host = flag.String("host", "", "Vector's IP address")
	flag.Parse()

	if *host == "" {
		log.Fatal("please use the -host argument and set it to your robots IP address")
	}

	c, err := client.New(
		client.WithTarget(
			fmt.Sprintf("%s:443", *host),
		),
		client.WithInsecureSkipVerify(),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	vc := vectorpb.NewExternalInterfaceClient(c.Conn())

	login, err := vc.UserAuthentication(context.Background(),
		&vectorpb.UserAuthenticationRequest{
			UserSessionId: []byte(sessionID),
			ClientName:    []byte(clientName),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(login)

}
