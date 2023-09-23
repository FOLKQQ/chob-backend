package jwtverifyModel

type Jwtverify struct {
	Username string `json:"username"`
	Role_id  string `json:"role_id"`
	Pstag_id string `json:"pstag_id"`
	Team_id  string `json:"team_id"`
	Exp      string `json:"exp"`
}
