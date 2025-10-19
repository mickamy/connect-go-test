package session

import (
	"net/http"

	"connectrpc.com/connect"

	"github.com/mickamy/connecttest-example/gen/github.com/mickamy/connecttest-example/gen/auth/v1/authv1connect"
	"github.com/mickamy/connecttest-example/internal/domain/session/handler"
	"github.com/mickamy/connecttest-example/internal/domain/session/usecase"
)

func Route(mux *http.ServeMux, options ...connect.HandlerOption) {
	mux.Handle(authv1connect.NewSessionServiceHandler(handler.NewSession(
		usecase.NewCreateSession(),
	), options...))
}
