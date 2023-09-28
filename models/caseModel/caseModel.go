package casemodel

type Case struct {
	Id              uint   `json:"id"`
	Service_id      uint   `json:"service_id"`
	Admin_id        uint   `json:"admin_id"`
	Payment_date    string `json:"payment_date"`
	Name            string `json:"name"`
	Tax_id          uint   `json:"tax_id"`
	Income_type     uint   `json:"income_type"`
	Amount          uint   `json:"amount"`
	Withholding_tax uint   `json:"withholding_tax"`
	Net_income      uint   `json:"net_income"`
	Number          uint   `json:"number"`
	Village         string `json:"village"`
	Sub_district    string `json:"sub_district"`
	District        string `json:"district"`
	Province        string `json:"province"`
	Zip_code        uint   `json:"zip_code"`
	View            uint   `json:"view"`
	Status          string `json:"status"`
	Timestamps      string `json:"timestamps"`
}
