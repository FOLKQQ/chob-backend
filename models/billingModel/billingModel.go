package billingModel

type Billing struct {
	Id           uint   `json:"id"`
	Client_id    uint   `json:"client_id"`
	Invoice      string `json:"invoice"`
	Invoice_date string `json:"invoice_date"`
	Date_due     string `json:"date_due"`
	Date_paid    string `json:"date_paid"`
	Totol        string `json:"totol"`
	Status       string `json:"status"`
	Timestamps   string `json:"timestamps"`
}

type Billing_tax struct {
	Id      uint `json:"id"`
	Task_id uint `json:"task_id"`
	Price   uint `json:"price"`
}
