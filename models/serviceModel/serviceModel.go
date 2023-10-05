package servicemodel

type Service struct {
	Id             uint   `json:"id"`
	Servicetype_id uint   `json:"servicetype_id"`
	Company_id     uint   `json:"company_id"`
	Date_start     string `json:"date_start"`
	Date_due       string `json:"date_dule"`
	Timestamps     string `json:"timestamps"`
}

type Servicetype struct {
	Id           uint   `json:"id"`
	Service_name string `json:"service_name"`
	Detail       string `json:"detail"`
	Price        uint   `json:"price"`
	Timestamps   string `json:"timestamps"`
}
