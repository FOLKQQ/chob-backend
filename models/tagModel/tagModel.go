package tagModel

type Tag struct {
	Id         uint   `json:"id"`
	Company_id uint   `json:"company_id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
}
