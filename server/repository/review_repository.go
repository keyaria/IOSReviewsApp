package repository

import (
	"context"
	"net/http"
	"net/url"
	"reviewsApp/server/data/request"
	"reviewsApp/server/models"
)

type ReviewRepository interface {
	Save(ctx context.Context, review models.Review)
	SaveAll(review request.ReviewCreateRequest)
	FindAll(w http.ResponseWriter, ctx context.Context, res url.Values) []models.Review
}
