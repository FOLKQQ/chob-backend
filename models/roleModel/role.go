package roleModel

type Role struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Project     string `json:"project"`
	Managerrole string `json:"manager_role"`
	Addtags     string `json:"addtags"`
	Report      string `json:"report"`
	Timestamps  string `json:"timestamps"`
}
