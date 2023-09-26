package casemodel

type Case struct {
	Id           uint   `json:"id"`
	Service_id   uint   `json:"service_id"`
	Admin_id     uint   `json:"admin_id"`
	Payment_date string `json:"payment_date"`
	Name         string `json:"name"`
	Tex_id       string `json:"tex_id"`
	Income_type  string `json:"income_type"`
	Amount       uint   `json:"amount"`
	Withholding  uint   `json:"withholding"`
	Net_income   uint   `json:"net_income"`
	Vaillge      string `json:"vaillge"`
	Sub_district string `json:"sub_district"`
	District     string `json:"district"`
	Province     string `json:"province"`
	Zip_code     uint   `json:"zip_code"`
	View         uint   `json:"view"`
	Status       uint   `json:"status"`
	Timestamps   string `json:"timestamps"`
}
