package taskModel

type Task struct {
	Id              uint   `json:"id"`
	Company_id      uint   `json:"company_id"`
	Title           string `json:"title"`
	Tax_status      string `json:"tax_status"`
	Tasklist_status string `json:"tasklist_status"`
	Timestamps      string `json:"timestamps"`
}

type Subtask struct {
	Id             uint   `json:"id"`
	Task_id        uint   `json:"task_id"`
	Title          string `json:"title"`
	Subtask_status string `json:"subtask_status"`
	Timestamps     string `json:"timestamps"`
}

type Taskdue struct {
	Id         uint   `json:"id"`
	Task_id    uint   `json:"task_id"`
	Date_due   string `json:"date_due"`
	Timestamps string `json:"timestamps"`
}

type Taskassignees struct {
	Id       uint `json:"id"`
	Task_id  uint `json:"task_id"`
	Assignee uint `json:"assignee"`
	User_id  uint `json:"user_id"`
}
