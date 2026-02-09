package pubsub

type PubSubPushRequest struct {
	Message      PubSubMessage `json:"message"`
	Subscription string        `json:"subscription"`
}

type PubSubMessage struct {
	Data        string            `json:"data"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	MessageID   string            `json:"messageId"`
	PublishTime string            `json:"publishTime"`
}
