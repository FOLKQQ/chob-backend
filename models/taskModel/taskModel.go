package taskModel

import "time"

type Task struct {
	Id              uint   `json:"id"`
	Company_id      uint   `json:"company_id"`
	Title           string `json:"title"`
	Tax_status      string `json:"tax_status"`
	Tasklist_status string `json:"tasklist_status"`
	Timestamps      string `json:"timestamps"`
}

// worker is array id and name
type Worker struct {
	Id uint
}

type Taskadd struct {
	ProjectName string    `json:"projectname"`
	CompanyName string    `json:"company_name"`
	StartDate   time.Time `json:"startdate"`
	EndDate     time.Time `json:"enddate"`
	Repeat      string    `json:"repeat"`
	Cyclemonth  string    `json:"cyclemonth"`
	Worker      []int     `json:"worker"`
	Checker     []int     `json:"checker"`
	Service     []int     `json:"service"`
}

type Subtask struct {
	Id             uint   `json:"id"`
	Task_id        uint   `json:"task_id"`
	Title          string `json:"title"`
	Subtask_status string `json:"subtask_status"`
	Timestamps     string `json:"timestamps"`
}

/*type SubTasklist struct {
Id             uint   `json:"id"`
subtask_id        uint   `json:"task_id"`
Title          string `json:"title"`
Subtasklist_status string `json:"subtask_status"`
Timestamps     string `json:"timestamps"`

*/

/*type SubTaskdue struct {
	Id         uint   `json:"id"`
	Subtask_id    uint   `json:"task_id"`
	Date_start string `json:"date_start"`
	Date_due   string `json:"date_due"`
	Timestamps string `json:"timestamps"`
}*/

type Taskassignees struct {
	Id         uint   `json:"id"`
	Task_id    uint   `json:"task_id"`
	User_id    uint   `json:"user_id"`
	Timestamps string `json:"timestamps"`
}
