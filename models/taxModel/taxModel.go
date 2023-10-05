package taxModel

type Tax_from struct {
	Task_id        uint   `json:"task_id"`
	Payment_date   string `json:"payment_date"`
	Name           string `json:"name"`
	Tax_id         uint   `json:"tax_id"`
	Income_type    string `json:"income_type"`
	Amount         uint   `json:"amount"`
	Witholding_tax uint   `json:"witholding_tax"`
	Net_income     uint   `json:"net_income"`
	Number         string `json:"number"`
	Village        string `json:"village"`
	Sub_district   string `json:"sub_district"`
	District       string `json:"district"`
	Province       string `json:"province"`
	Zipcode        string `json:"zipcode"`
	View           uint   `json:"view"`
	Status         string `json:"status"`
	Timestamps     string `json:"timestamps"`
}

type Tax_30 struct {
	Task_id                             uint   `json:"task_id"`
	Excess_tax_paid                     uint   `json:"excess_tax_paid"`
	Excess_tax_paid_from_previous_month uint   `json:"excess_tax_paid_from_previous_month"`
	Excess_tax_not_to_be_paid           uint   `json:"excess_tax_not_to_be_paid"`
	Net_tax_payable                     uint   `json:"net_tax_payable"`
	Sales_total                         uint   `json:"sales_total"`
	Sales_tax                           uint   `json:"sales_tax"`
	Purchase_total                      uint   `json:"purchase_total"`
	Purchase_tax                        uint   `json:"purchase_tax"`
	Timestamp                           string `json:"timestamp"`
}
