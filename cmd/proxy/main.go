package main

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/pspp/internal/http"
	"github.com/hjfitz/pspp/internal/pubsub"
	"go.uber.org/zap"
)

const PROXY_HOST = "http://localhost:3000"
const SUBSCRIPTION = "projects/local/subscriptons/mocked"

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	r := gin.Default()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := logger.Sugar()

	ps := pubsub.NewPubsub(SUBSCRIPTION)
	handler := http.NewProxyHandler(
		PROXY_HOST,
		ps,
		l,
	)

	r.Use(handler.Handle)

	if err := r.Run(); err != nil {
		l.Errorw("failed to run server", "err", err)
	}
}
