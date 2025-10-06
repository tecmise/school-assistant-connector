package assistant

import (
	"context"

	"github.com/tecmise/connector-lib/pkg/adapters/outbound/client_rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		ChatPrompt(ctx context.Context, request any) (Response, error)
	}

	client struct {
		mapper connector.Call[Response]
		host   string
	}
)

func Rest(host string) Client {
	return &client{
		host:   host,
		mapper: client_rest.NewConnector[Response](),
	}
}

func Lambda(identifier string) Client {
	return &client{
		host:   identifier,
		mapper: client_rest.NewConnector[Response](),
	}
}

func (c client) ChatPrompt(ctx context.Context, request any) (Response, error) {
	var response Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource("assistant/chat/prompt").
		WithBody(request).
		WithMethod("POST").
		Build()
	return response, c.mapper.Create(parameter, &response)
}
