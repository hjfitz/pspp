package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/pspp/internal/pubsub"
	"go.uber.org/zap"
)

type ProxyHandler struct {
	Host   string
	Pubsub pubsub.PubSub
	Logger *zap.SugaredLogger
}

func NewProxyHandler(
	host string,
	pubsub pubsub.PubSub,
	logger *zap.SugaredLogger,
) ProxyHandler {
	return ProxyHandler{
		Host:   host,
		Pubsub: pubsub,
		Logger: logger,
	}
}

func (p *ProxyHandler) Handle(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	upstream := fmt.Sprintf("%s%s", p.Host, path)

	jsonData, _ := io.ReadAll(ctx.Request.Body)
	payload := p.Pubsub.ToPubSubPayload(jsonData)
	reqPayload, _ := json.Marshal(payload)
	buf := bytes.NewBuffer(reqPayload)

	_, err := http.Post(upstream, "application/json", buf)
	if err != nil {
		p.Logger.Errorw("unable to post upstream",
			"err", err,
		)
	}
	p.Logger.Infow("Request handled",
		"path", path,
		"upstream", upstream,
		"payload", string(reqPayload),
	)

	// todo: proxy response
	ctx.JSON(http.StatusOK, gin.H{"text": "ok"})
}
