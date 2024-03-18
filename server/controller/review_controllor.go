package controller

import (
	"net/http"
	"reviewsApp/server/data/request"
	"reviewsApp/server/data/response"
	"reviewsApp/server/helpers"
	"reviewsApp/server/service"

	"github.com/julienschmidt/httprouter"
)

type ReviewController struct {
	ReviewService service.ReviewService
}

func NewReviewController(ReviewService service.ReviewService) *ReviewController {
	return &ReviewController{ReviewService: ReviewService}
}

func (controller *ReviewController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	ReviewCreateRequest := request.ReviewCreateRequest{}
	helpers.ReadRequestBody(requests, &ReviewCreateRequest)

	controller.ReviewService.Create(requests.Context(), ReviewCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helpers.WriteResponseBody(writer, webResponse)
}

func (controller *ReviewController) CreateAll(review request.ReviewCreateRequest) {
	//ReviewCreateRequest := request.ReviewCreateRequest{}

	controller.ReviewService.CreateAll(review)
	// webResponse := response.WebResponse{
	// 	Code:   200,
	// 	Status: "Ok",
	// 	Data:   nil,
	// }

	// helpers.WriteResponseBody(writer, webResponse)
}

func (controller *ReviewController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	result := controller.ReviewService.FindAll(writer, requests.Context(), requests.URL.Query())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helpers.WriteResponseBody(writer, webResponse)
}
