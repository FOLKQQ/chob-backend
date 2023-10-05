package adminModel

type Admin struct {
	ID         int    `json:"id"`
	RoleID     int    `json:"role_id"`
	TeamID     int    `json:"team_id"`
	UserID     string `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Image      string `json:"image"`
	Tal        string `json:"tal"`
	Token_link string `json:"token_link"`
	Status     string `json:"status"`
	Timestamps string `json:"timestamps"`
}
type AddAdmin struct {
	RoleID     int    `json:"role_id"`
	TeamID     int    `json:"team_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Tal        string `json:"tal"`
	Token_link string `json:"token_link"`
	Status     string `json:"status"`
}
