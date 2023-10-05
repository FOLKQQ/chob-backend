package chatModel

type Chat_team struct {
	Id         uint   `json:"id"`
	User_id    uint   `json:"user_id"`
	Company_id uint   `json:"company_id"`
	Comment    string `json:"comment"`
	Timestamps string `json:"timestamps"`
}

type Chat_task struct {
	Id         uint   `json:"id"`
	Task_id    uint   `json:"task_id"`
	User_id    uint   `json:"user_id"`
	Comment    string `json:"comment"`
	Timestamps string `json:"timestamps"`
}
