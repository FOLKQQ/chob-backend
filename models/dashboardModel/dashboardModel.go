package dashboardmodel

type Dashboard_case struct {
	Id                 uint `json:"id"`
	Admin_id           uint `json:"admin_id"`
	Dashboard_sbt_taxs []Dashboard_sbt_tax
}

type Dashboard_service struct {
	Id             uint   `json:"id"`
	Company_id     uint   `json:"company_id"`
	Servicetype_id uint   `json:"servicetype_id"`
	Date_start     string `json:"date_start"`
	Date_end       string `json:"date_end"`
}

type Dashboard_company struct {
	Id           uint   `json:"id"`
	Type_company uint   `json:"type_company"`
	Company_name string `json:"company_name"`
}

type Dashboard_servicetype struct {
	Id           uint   `json:"id"`
	Service_name string `json:"service_name"`
}

type Dashboard_sbt_tax struct {
	Case_id    uint   `json:"case_id"`
	E_tax_name string `json:"e_tax_name"`
	Status     string `json:"status"`
}
