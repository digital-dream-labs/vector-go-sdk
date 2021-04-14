package main

import (
	"context"
	"log"
	"os"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vector"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

func main() {

	v, err := vector.NewEP(
		vector.WithTarget(os.Getenv("BOT_TARGET")),
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
			_, _ = v.Conn.SayText(
				ctx,
				&vectorpb.SayTextRequest{
					Text:           "hello, hello, hello",
					UseVectorVoice: true,
					DurationScalar: 1.0,
				},
			)
			stop <- true
			return
		}
	}
}
