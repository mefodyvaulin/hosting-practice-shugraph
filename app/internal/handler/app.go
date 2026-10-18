package handler

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/feytox/shugraph/backend/graph-api/internal/config"
)

type AppHandler struct {
	config config.AppConfig
}

func NewAppHandler(config config.AppConfig) *AppHandler {
	return &AppHandler{
		config: config,
	}
}
func (h *AppHandler) Register(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "appConfig",
		Summary:     "Return app configuration",
		Method:      http.MethodGet,
		Path:        "/",
	}, h.AppConfig)
}

type AppConfigOutput struct {
	Body struct {
		ID         string
		AppMessage string
		AppToken   bool
	}
}

func (h *AppHandler) AppConfig(_ context.Context, _ *struct{}) (*AppConfigOutput, error) {
	response := &AppConfigOutput{}
	response.Body.ID = "Mefody"
	response.Body.AppMessage = h.config.AppMessage
	response.Body.AppToken = h.config.AppToken != ""
	return response, nil

}
