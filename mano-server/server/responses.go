package server

type SuggestionsResponse struct {
	Suggestions string `json:"suggestions"`
}

type ErrorResponse struct {
	Type    string
	Message string
}
