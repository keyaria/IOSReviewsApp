package models

type APIResponse struct {
	Feed struct {
		Author struct {
			Name struct {
				Label string `json:"label"`
			} `json:"name"`
			URI struct {
				Label string `json:"label"`
			} `json:"uri"`
		} `json:"author"`
		Entry []struct {
			Author struct {
				URI struct {
					Label string `json:"label"`
				} `json:"uri"`
				Name struct {
					Label string `json:"label"`
				} `json:"name"`
				Label string `json:"label"`
			} `json:"author"`
			Updated struct {
				Label string `json:"label"`
			} `json:"updated"`
			ImRating struct {
				Label string `json:"label"`
			} `json:"im:rating"`
			ImVersion struct {
				Label string `json:"label"`
			} `json:"im:version"`
			ID struct {
				Label string `json:"label"`
			} `json:"id"`
			Title struct {
				Label string `json:"label"`
			} `json:"title"`
			Content struct {
				Label      string `json:"label"`
				Attributes struct {
					Type string `json:"type"`
				} `json:"attributes"`
			} `json:"content"`
			Link struct {
				Attributes struct {
					Rel  string `json:"rel"`
					Href string `json:"href"`
				} `json:"attributes"`
			} `json:"link"`
			ImVoteSum struct {
				Label string `json:"label"`
			} `json:"im:voteSum"`
			ImContentType struct {
				Attributes struct {
					Term  string `json:"term"`
					Label string `json:"label"`
				} `json:"attributes"`
			} `json:"im:contentType"`
			ImVoteCount struct {
				Label string `json:"label"`
			} `json:"im:voteCount"`
		} `json:"entry"`
		Updated struct {
			Label string `json:"label"`
		} `json:"updated"`
		Rights struct {
			Label string `json:"label"`
		} `json:"rights"`
		Title struct {
			Label string `json:"label"`
		} `json:"title"`
		Icon struct {
			Label string `json:"label"`
		} `json:"icon"`
		Link []struct {
			Attributes struct {
				Rel  string `json:"rel"`
				Type string `json:"type"`
				Href string `json:"href"`
			} `json:"attributes,omitempty"`
			Attributes0 struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"attributes,omitempty"`
			Attributes1 struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"attributes,omitempty"`
			Attributes2 struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"attributes,omitempty"`
			Attributes3 struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"attributes,omitempty"`
			Attributes4 struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"attributes,omitempty"`
		} `json:"link"`
		ID struct {
			Label string `json:"label"`
		} `json:"id"`
	} `json:"feed"`
}

type Review struct {
	Id      int    `json:"id"`
	Name    string `json: "name"`
	Updated string `json: "updated"`
	Rating  string `json: "im:rating"`
	Title   string `json: "title"`
	Content string `json: "content"`
}

type ReviewResponse struct {
	Name    string `json: "name"`
	Updated string `json: "updated"`
	Rating  string `json: "im:rating"`
	Title   string `json: "title"`
	Content string `json: "content"`
}
