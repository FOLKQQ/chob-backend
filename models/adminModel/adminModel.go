package adminModel

type Admin struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Fistname   string `json:"fistname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Tal        string `json:"tal"`
	Image      string `json:"image"`
	Status     string `json:"status"`
	Timestamps string `json:"timestamps"`
	Role_id    uint   `json:"role_id"`
	Pstag_id   uint   `json:"pstag_id"`
	Team_id    uint   `json:"team_id"`
	Token_link string `json:"token_link"`
	User_id    string `json:"user_id"`
}

type AddAdmin struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	RoleID    int    `json:"role_id"`
	PstagID   int    `json:"pstag_id"`
	TeamID    int    `json:"team_id"`
}
