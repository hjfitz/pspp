package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/pspp/internal/http"
	"github.com/hjfitz/pspp/internal/pubsub"
	"go.uber.org/zap"
)

const SUBSCRIPTION = "projects/local/subscriptons/mocked"

func main() {
	host := os.Args[3]

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.Default()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()

	ps := pubsub.NewPubsub(SUBSCRIPTION)
	handler := http.NewProxyHandler(host, ps, l)

	r.Use(handler.Handle)

	if err := r.Run(); err != nil {
		l.Errorw("failed to run server", "err", err)
	}
}
