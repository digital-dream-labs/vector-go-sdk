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

	ctx := context.Background()
	start := make(chan bool)
	stop := make(chan bool)

	go func() {
		_ = v.BehaviorControl(ctx, start, stop)
	}()

	for {
		select {
		case <-start:
			_, err := v.Conn.SetLiftHeight(
				ctx,
				&vectorpb.SetLiftHeightRequest{
					HeightMm:          250,
					MaxSpeedRadPerSec: 1,
					IdTag:             2000001,
					DurationSec:       .001,
				},
			)
			if err != nil {
				fmt.Println(err)
			}
			stop <- true
			return
		}
	}

}
