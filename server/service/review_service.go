package service

import (
	"context"
	"net/http"
	"net/url"
	"reviewsApp/server/data/request"
	"reviewsApp/server/data/response"
)

type ReviewService interface {
	Create(ctx context.Context, request request.ReviewCreateRequest)
	CreateAll(request request.ReviewCreateRequest)
	FindAll(w http.ResponseWriter, ctx context.Context, res url.Values) []response.ReviewResponse
}
