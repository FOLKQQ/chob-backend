package templatemodel

type Template struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Templatetask struct {
	Id              uint   `json:"id"`
	Tmp_id          uint   `json:"tmp_id"`
	Title           string `json:"title"`
	Tax_status      string `json:"tax_status"`
	Tasklist_status string `json:"tasklist_status"`
}

type Templatesubtask struct {
	Id             uint   `json:"id"`
	Tmptask_id     uint   `json:"tmptask_id"`
	Title          string `json:"title"`
	Detail         string `json:"detail"`
	Subtask_status string `json:"subtask_status"`
}

type Templatesubtasklist struct {
	Id                 uint   `json:"id"`
	Tmpsubtask_id      uint   `json:"tmpsubtask_id"`
	Title              string `json:"title"`
	Subtasklist_status string `json:"subtasklist_status"`
}
