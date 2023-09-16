package roleModel

type Role struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Widgets    string `json:"widgets"`
	Status     string `json:"status"`
	Timestamps string `json:"timestamps"`
}

type AddRole struct {
	Name    string `json:"name"`
	Widgets string `json:"widgets"`
	Status  string `json:"status"`
}
