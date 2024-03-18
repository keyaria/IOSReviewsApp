package router

import (
	"fmt"
	"net/http"
	"reviewsApp/server/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(reviewController *controller.ReviewController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "API Is Working :)", "available endpoints: GET /api/reviews")
	})

	router.GET("/api/reviews", reviewController.FindAll)
	router.POST("/api/review", reviewController.Create)

	return router
}
