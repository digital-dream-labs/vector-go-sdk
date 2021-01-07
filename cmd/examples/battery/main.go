package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vector"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

func main() {

	v, err := vector.New(
		vector.WithTarget(os.Getenv("BOT_TARGET")),
		vector.WithToken(os.Getenv("BOT_TOKEN")),
	)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := v.Conn.BatteryState(
		context.Background(),
		&vectorpb.BatteryStateRequest{},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)

}
