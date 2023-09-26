package servicemodel

type Service struct {
	Id         uint   `json:"id"`
	Company_id uint   `json:"company_id"`
	Service_id uint   `json:"service_id"`
	Date_start string `json:"date_start"`
	Date_end   string `json:"date_end"`
	Timestamps string `json:"timestamps"`
	Team_id    uint   `json:"team_id"`
}

type Servicetype struct {
	Id           uint   `json:"id"`
	Service_name string `json:"service_name"`
	Detail       string `json:"detail"`
	Price        uint   `json:"price"`
	Timestamps   string `json:"timestamps"`
}
