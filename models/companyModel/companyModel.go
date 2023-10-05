package companyModel

type Company struct {
	ID                 int    `json:"id"`
	TypeCompany        string `json:"type_company"`
	CompanyName        string `json:"company_name"`
	CodeIdentification string `json:"code_identification"`
	Signature          string `json:"signature"`
	RegisCompany       string `json:"regis_company"`
	RegisVat           string `json:"regis_vat"`
	BusinessType       string `json:"business_type"`
	IDDBD              string `json:"id_dbd"`
	PassDBD            string `json:"pass_dbd"`
	IDFiling           string `json:"id_filing"`
	PassFiling         string `json:"pass_filing"`
	Email              string `json:"email"`
	Tal                string `json:"tal"`
	Status             string `json:"status"`
	Timestamps         string `json:"timestamps"`
}
