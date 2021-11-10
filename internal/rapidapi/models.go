package rapidapi

import "net/http"

type Recipe struct {
	Title          string `json:"title"`
	ReadyInMinutes int    `json:"readyInMinutes"`
	Image          string `json:"image"`
	SourceURL      string `json:"sourceUrl"`
}

func (rec *Recipe) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
