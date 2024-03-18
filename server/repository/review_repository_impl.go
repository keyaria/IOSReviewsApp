package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"reviewsApp/server/data/request"
	"reviewsApp/server/helpers"
	"reviewsApp/server/models"
	"time"
)

type ReviewRepositoryImpl struct {
	Db *sql.DB
}

func NewReviewRepository(Db *sql.DB) ReviewRepository {
	return &ReviewRepositoryImpl{Db: Db}
}

// FindAll implements ReviewRepository
func (r *ReviewRepositoryImpl) FindAll(w http.ResponseWriter, ctx context.Context, res url.Values) []models.Review {

	//Get parameters from URL
	filters := res.Get("time")
	tx, err := r.Db.Begin()

	helpers.ThrowIfError(err)
	defer helpers.CommitOrRollback(tx)

	// Get the current time and calculate the time 48 hours ago
	now := time.Now()
	calc, err := time.ParseDuration(filters + "h")
	past48Hours := now.Add(-calc)
	layout := "2006-01-02T15:04:05-07:00"

	// Format the current time using the layout
	formattedTime := past48Hours.Format(layout)

	SQL := "select id,name,updated,rating, title, content from review where updated >= $1"

	result, errQuery := tx.QueryContext(ctx, SQL, formattedTime)

	helpers.ThrowIfError(errQuery)
	defer result.Close()

	var reviews []models.Review

	for result.Next() {
		review := models.Review{}
		err := result.Scan(&review.Id, &review.Name, &review.Updated, &review.Rating, &review.Title, &review.Content)
		helpers.ThrowIfError(err)

		reviews = append(reviews, review)
	}

	return reviews
}

// SaveAll implements ReviewRepository
// while Not an API Call putting here to keep DB calls in one location for now
func (r *ReviewRepositoryImpl) SaveAll(review request.ReviewCreateRequest) {

	tx, err := r.Db.Begin()
	helpers.ThrowIfError(err)
	defer helpers.CommitOrRollback(tx)

	fmt.Println("review", review)
	SQL := "insert into review(name, updated, rating, title, content) values ($1, $2, $3, $4, $5)"
	layout := "2006-01-02T15:04:05-07:00"
	timestamp, errTime := time.Parse(layout, review.Updated)
	fmt.Println("timestamp", review.Updated)
	if errTime != nil {
		fmt.Println("Error parsing timestamp:", errTime)
		return
	}

	_, err = tx.Exec(SQL, review.Name, timestamp, review.Rating, review.Title, review.Content)

	helpers.ThrowIfError(err)

	fmt.Println("Finished Adding to Database")

}

// Save implements ReviewRepository, Intially was for testing
func (r *ReviewRepositoryImpl) Save(ctx context.Context, review models.Review) {
}
