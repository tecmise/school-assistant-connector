package assistant

import (
	"context"

	"github.com/tecmise/connector-lib/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/rest"
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
		mapper: rest.NewConnector[Response](),
	}
}

func Lambda(identifier string) Client {
	return &client{
		host:   identifier,
		mapper: lambda.NewConnector[Response](),
	}
}

func (c client) ChatPrompt(ctx context.Context, request any) (Response, error) {
	var school Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("assistant/chat/prompt").
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("POST").
		Build()
	return school, c.mapper.Create(parameter, &school)
}
