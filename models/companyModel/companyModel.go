package companyModel

type Company struct {
	ID                 int    `json:"id"`
	TypeCompany        string `json:"type_company"`
	CompanyName        string `json:"company_name"`
	CodeIdentification string `json:"code_identification"`
	Signature          string `json:"signature"`
	StartComDate       string `json:"startComDate"`
	StartVatDate       string `json:"startVatDate"`
	BusinessType       string `json:"business_type"`
	Id_dbd             string `json:"id_dbd"`
	Pass_dbd           string `json:"pass_dbd"`
	Id_filing          string `json:"id_filing"`
	Pass_filing        string `json:"pass_filing"`
	Email              string `json:"email"`
	Tal                string `json:"tal"`
	Status             string `json:"status"`
	Address            string `json:"address"`
	Addressextra       string `json:"addressextra"`
	Subdistrict        string `json:"subdistrict"`
	District           string `json:"district"`
	Province           string `json:"province"`
	Zipcode            string `json:"zipcode"`
	Timestamps         string `json:"timestamps"`
}
