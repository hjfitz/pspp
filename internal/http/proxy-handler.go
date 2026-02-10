package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjfitz/pspp/internal/pubsub"
	"github.com/rs/zerolog"
)

type ProxyHandler struct {
	Host   string
	Pubsub pubsub.PubSub
	Logger zerolog.Logger
}

func NewProxyHandler(
	host string,
	pubsub pubsub.PubSub,
	logger zerolog.Logger,
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

	resp, err := http.Post(upstream, "application/json", buf)
	if err != nil {
		p.Logger.Error().Err(err).Msg("unable to post upstream")
	}
	p.Logger.Log().Str("path", path).Str("upstream", upstream).Msg("Request handled")

	status := resp.StatusCode
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	ct := resp.Header.Get("content-type")

	ctx.Data(status, ct, body)
}
