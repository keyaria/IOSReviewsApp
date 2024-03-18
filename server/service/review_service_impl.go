package service

import (
	"context"
	"net/http"
	"net/url"
	"reviewsApp/server/data/request"
	"reviewsApp/server/data/response"

	"reviewsApp/server/models"
	"reviewsApp/server/repository"
)

type ReviewServiceImpl struct {
	ReviewRepository repository.ReviewRepository
}

func NewReviewServiceImpl(ReviewRepository repository.ReviewRepository) ReviewService {
	return &ReviewServiceImpl{ReviewRepository: ReviewRepository}
}

// Create implements ReviewService
func (r *ReviewServiceImpl) Create(ctx context.Context, request request.ReviewCreateRequest) {
	review := models.Review{
		Name:    request.Name,
		Updated: request.Updated,
		Rating:  request.Rating,
		Title:   request.Title,
		Content: request.Content,
	}
	r.ReviewRepository.Save(ctx, review)
}

// FindAll implements ReviewService
func (r *ReviewServiceImpl) FindAll(w http.ResponseWriter, ctx context.Context, res url.Values) []response.ReviewResponse {
	reviews := r.ReviewRepository.FindAll(w, ctx, res)

	var reviewResp []response.ReviewResponse

	for _, value := range reviews {
		review := response.ReviewResponse{Id: value.Id, Name: value.Name, Updated: value.Updated, Rating: value.Rating, Title: value.Title, Content: value.Content}
		reviewResp = append(reviewResp, review)
	}
	return reviewResp

}

func (r *ReviewServiceImpl) CreateAll(request request.ReviewCreateRequest) {
	// review := models.Review{
	// 	Name:    request.Name,
	// 	Updated: request.Updated,
	// 	Rating:  request.Rating,
	// 	Title:   request.Title,
	// 	Content: request.Content,
	// }
	r.ReviewRepository.SaveAll(request)
}

// TODO: Update API
