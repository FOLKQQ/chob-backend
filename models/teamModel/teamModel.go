package teammodel

type Team struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Timestamps string `json:"timestamps"`
}
