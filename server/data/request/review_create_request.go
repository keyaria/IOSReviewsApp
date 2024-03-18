package request

type ReviewCreateRequest struct {
	Name    string `validate:"required min=1,max=100" json:"name"`
	Updated string `json: "updated"`
	Rating  string `json: "im:rating"`
	Title   string `json: "title"`
	Content string `json: "content"`
}
