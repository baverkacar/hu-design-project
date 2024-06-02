package request

import (
	"hu-design-project/application/handler/request/handler_impl"
	"hu-design-project/application/repository"
)

func InitializeRequestHandler(
	repository repository.RequestRepository,
) *Handler {
	requestHandler := Handler{}
	requestHandler.GetRequests = handler_impl.NewGetRequestHandler(repository)
	return &requestHandler
}
