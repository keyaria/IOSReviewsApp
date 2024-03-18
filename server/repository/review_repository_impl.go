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

	//w.Header().Set("Content-Type", "application/json")

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

// Save implements ReviewRepository
func (r *ReviewRepositoryImpl) Save(ctx context.Context, review models.Review) {

	// resp, err := http.Get("https://itunes.apple.com/us/rss/customerreviews/id=595068606/sortBy=mostRecent/page=1/json")
	// if err != nil {
	// 	helpers.ThrowIfError(err)
	// }
	// defer resp.Body.Close()

	// Check if the request was successful (status code 200)
	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Println("Error: unexpected status code", resp.StatusCode)
	// 	return
	// }
	// tx, err := r.Db.Begin()
	// helpers.ThrowIfError(err)
	// defer helpers.CommitOrRollback(tx)

	// // resp, err := http.Get("https://itunes.apple.com/us/rss/customerreviews/id=595068606/sortBy=mostRecent/page=1/json")
	// // if err != nil {
	// // 	log.Fatalln(err)
	// // }
	// // //We Read the response body on the line below.
	// // body, err := io.ReadAll(resp.Body)

	// // if err != nil {
	// // 	log.Fatalln(err)
	// // }

	// // var response models.AutoGenerated
	// // if err := json.Unmarshal(body, &response); err != nil {
	// // 	fmt.Println("Error unmarshalling JSON:", err)
	// // 	return
	// // }

	// // //fmt.Print(response)

	// // for _, entry := range response.Feed.Entry {
	// // 	fmt.Println(entry.Title.Label)
	// // }

	// SQL := "insert into review(name, updated, rating, title, content) values ($1, $2, $3, $4, $5)"
	// _, err = tx.ExecContext(ctx, SQL, review.Name, review.Updated, review.Rating, review.Title, review.Content)
	// helpers.ThrowIfError(err)
}

// // Update implements BookRepository
// func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
// 	tx, err := b.Db.Begin()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	SQL := "update book set name=$1 where id=$2"
// 	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)
// 	helper.PanicIfError(err)
// }
