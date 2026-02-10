package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/pspp/internal/args"
	"github.com/hjfitz/pspp/internal/http"
	"github.com/hjfitz/pspp/internal/logger"
	"github.com/hjfitz/pspp/internal/pubsub"
	"github.com/hjfitz/pspp/internal/validators"
)

const SUBSCRIPTION = "projects/local/subscriptons/mocked"

func main() {
	l := logger.NewLogger()
	opts := args.GetOpts()

	if !validators.IsValidUrl(opts.Upstream) {
		l.Error().Str("upstream", opts.Upstream).Msg("Unable to start program")
		os.Exit(1)
	}

	if  !validators.IsValidPort(opts.Port) {
		l.Error().Str("port", opts.Port).Msg("Unable to start program")
		os.Exit(1)
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.Default()


	ps := pubsub.NewPubsub(SUBSCRIPTION)
	handler := http.NewProxyHandler(opts.Upstream, ps, l)

	r.Use(handler.Handle)

	l.Info().Msg("Starting server")
	addr := fmt.Sprintf(":%s", opts.Port)
	if err := r.Run(addr); err != nil {
		l.Error().Err(err).Msg("failed to run server")
	}
}
