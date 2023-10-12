package tagModel

type Tag struct {
	Id      uint   `json:"id"`
	Task_id uint   `json:"task_id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
}
