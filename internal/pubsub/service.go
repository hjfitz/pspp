package pubsub

import (
	"encoding/base64"
	"time"
)

type PubSub struct {
	subscription string
}

func NewPubsub(subscription string) PubSub {
	return PubSub{subscription}
}

func (p *PubSub) ToPubSubPayload(rawBody []byte) PubSubPushRequest {
	now := time.Now().UTC().Format(time.RFC3339Nano)
	messageID := genMessageID()

	return PubSubPushRequest{
		Subscription: p.subscription,
		Message: PubSubMessage{
			Data:        base64.StdEncoding.EncodeToString(rawBody),
			Attributes:  map[string]string{},
			MessageID:   messageID,
			PublishTime: now,
		},
	}
}
