package sbttaxmodel

type Sbttax struct {
	Id                                  uint   `json:"id"`
	Case_id                             uint   `json:"case_id"`
	E_tax_name                          uint   `json:"e_tax_name"`
	Sales_total                         uint   `json:"sales_total"`
	Sales_tax                           uint   `json:"sales_tax"`
	Purchase_total                      uint   `json:"purchase_total"`
	Purchase_tax                        uint   `json:"purchase_tax"`
	Excess_tax_paid                     uint   `json:"excess_tax_paid"`
	Excess_tax_paid_from_previous_month uint   `json:"excess_tax_paid_from_previous_month"`
	Excess_tax_not_to_be_paid           uint   `json:"excess_tax_not_to_be_paid"`
	Net_tax_payable                     uint   `json:"net_tax_payable"`
	Timestamps                          string `json:"timestamps"`
	Status                              string `json:"status"`
}
