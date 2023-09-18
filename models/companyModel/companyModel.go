package companyModel

type Company struct {
	Id                  uint   `json:"id"`
	Type_company        string `json:"type_company"`
	Company_name        string `json:"company_name"`
	Code_Identification string `json:"code_identification"`
	Signature           string `json:"signature"`
	Regis_company       string `json:"regis_company"`
	Regis_vat           string `json:"regis_vat"`
	Business_type       string `json:"business_type"`
	Id_dbd              string `json:"id_dbd"`
	Pass_dbd            string `json:"pass_dbd"`
	Id_filing           string `json:"id_filing"`
	Pass_filing         string `json:"pass_filing"`
	Email               string `json:"email"`
	Tel                 string `json:"tel"`
	Timestamps          string `json:"timestamps"`
}
