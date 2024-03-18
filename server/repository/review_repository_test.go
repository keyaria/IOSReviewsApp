package repository_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"reviewsApp/server/models"

	"testing"

	"github.com/stretchr/testify/mock"
)

func (u *ReviewMockRepository) GetReviews(ctx context.Context) ([]*models.Review, error) {
	args := u.Called(ctx)
	return args.Get(0).([]*models.Review), args.Error(1)
}

type ReviewMockRepository struct {
	mock.Mock
}
type ErrorMockResponseWriter struct {
	Body io.ReadCloser
	Code int
}

// TODO: MOCK the http write
func TestGetAll(t *testing.T) {

	tests := []struct {
		name   string
		output struct {
			reviews []*models.Review
			err     error
		}
	}{
		{
			name: "get all reviews without error",
			output: struct {
				reviews []*models.Review
				err     error
			}{
				reviews: []*models.Review{},
				err:     nil,
			},
		},
		{
			name: "get all reviews with error",
			output: struct {
				reviews []*models.Review
				err     error
			}{
				reviews: nil,
				err:     errors.New("Something got wrong"),
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepository := new(ReviewMockRepository)

			//responseWriter := errorMockResponseWriter()

			mockRepository.On("FindAll", mock.Anything).Return(tt.output.reviews, tt.output.err)

			//retrieving := repository.NewReviewRepository(&sql.DB{})

			//data := retrieving.FindAll(Write(), context.Background(), url.Values{"key": {"Value"}, "id": {"123"}})

			//assert.Equal(t, tt.output.err, data)

			mockRepository.AssertNumberOfCalls(t, "FindAll", 1)
			mockRepository.AssertExpectations(t)
		})
	}
}

func (e *ErrorMockResponseWriter) Header() http.Header {
	return http.Header{}
}
func (e *ErrorMockResponseWriter) Write(data []byte) {
	e.Body = io.NopCloser(bytes.NewReader(data))

}

func (e *ErrorMockResponseWriter) WriteHeader(statusCode int) {
	e.Code = statusCode
}
