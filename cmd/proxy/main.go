package main

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/pspp/internal/args"
	"github.com/hjfitz/pspp/internal/http"
	"github.com/hjfitz/pspp/internal/pubsub"
	"go.uber.org/zap"
)

const SUBSCRIPTION = "projects/local/subscriptons/mocked"

func main() {
	opts := args.GetOpts()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.Default()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()

	ps := pubsub.NewPubsub(SUBSCRIPTION)
	handler := http.NewProxyHandler(opts.Upstream, ps, l)

	r.Use(handler.Handle)

	if err := r.Run(); err != nil {
		l.Errorw("failed to run server", "err", err)
	}
}
