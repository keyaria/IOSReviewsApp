package router

import (
	// "fmt"
	// "net/http"
	"reviewsApp/server/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(reviewController *controller.ReviewController) *httprouter.Router {
	router := httprouter.New()

	// router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	fmt.Fprint(w, "Welcome Home")
	// })

	router.GET("/api/reviews", reviewController.FindAll)
	// TODO: Modify the api url to be clearer
	//router.POST("/api/reviews/create", reviewController.CreateAll)
	router.POST("/api/review", reviewController.Create)
	//router.PATCH("/api/review/:reviewId", reviewController.Update)
	//router.DELETE("/api/review/:reviewId", reviewController.Delete)

	return router
}
