package cmd

import (
	"context"
	"go-skeleton/cmd/app"
)

func StartServer(c context.Context) {
	ctx, cancel := context.WithCancel(c)

	startServer(ctx, cancel)

	<-ctx.Done()
}

func startServer(ctx context.Context, cancel context.CancelFunc) {
	s := app.New()
	s.Start(ctx, cancel)
}
