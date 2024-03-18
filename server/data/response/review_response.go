package response

type ReviewResponse struct {
	Id      int    `json:"id"`
	Name    string `json: "name"`
	Updated string `json: "updated"`
	Rating  string `json: "im:rating"`
	Title   string `json: "title"`
	Content string `json: "content"`
}
