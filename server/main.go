package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reviewsApp/server/config"
	"reviewsApp/server/controller"
	"reviewsApp/server/data/request"
	"reviewsApp/server/helpers"
	"reviewsApp/server/models"
	"reviewsApp/server/repository"
	"reviewsApp/server/router"
	"reviewsApp/server/service"
	"time"

	"github.com/rs/cors"
)

// TODO: replace httpRouter with native or explain why added
// add pagination, polling, add sort by var

func main() {
	fmt.Printf("Starting server")

	// database
	db := config.DatabaseConnection()

	createReviewTable(db)

	// repository
	reviewRepository := repository.NewReviewRepository(db)

	// service
	reviewService := service.NewReviewServiceImpl(reviewRepository)

	// controller
	reviewController := controller.NewReviewController(reviewService)

	// Fill DB with API
	InitializeData(reviewController)

	// router
	routes := router.NewRouter(reviewController)

	handler := cors.Default().Handler(routes)

	server := http.Server{Addr: ":8080", Handler: handler}

	err := server.ListenAndServe()

	helpers.ThrowIfError(err)

}

func InitializeData(c *controller.ReviewController) {
	apiUrl := "https://itunes.apple.com/us/rss/customerreviews/id=466965151/sortBy=mostRecent/page=1/json"

	// TODO: In Order to do Polling, create a timer function
	httpClient := http.Client{
		Timeout: time.Second * 10, // Adjust timeout as needed
	}

	resp, err := httpClient.Get(apiUrl)
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var response models.APIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, entry := range response.Feed.Entry {
		reviews := request.ReviewCreateRequest{Name: entry.Author.Name.Label, Updated: entry.Updated.Label, Rating: entry.ImRating.Label, Title: entry.Title.Label, Content: entry.Content.Label}

		c.CreateAll(reviews)

	}
}

// TODO: migrations is way to update this
func createReviewTable(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS "review" (
		"id" bigserial PRIMARY KEY,
		"name" varchar NOT NULL,
		"updated" timestamp NOT NULL,
		"rating" varchar NOT NULL,
		"title" varchar NOT NULL,
		"content" varchar NOT NULL
	  );
	  
	  CREATE INDEX ON "review" ("updated");
	  
	  CREATE INDEX ON "review" ("rating");
	  
	  COMMENT ON COLUMN "review"."rating" IS 'Can not be negative';`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}
