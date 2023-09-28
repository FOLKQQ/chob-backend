package dashboardmodel

type Dashboard_case struct {
	Id                uint `json:"id"`
	Service_id        uint `json:"service_id"`
	Dashboard_sbt_tax []Dashboard_sbt_tax
}

type Dashboard_service struct {
	Id             uint   `json:"id"`
	Company_id     uint   `json:"company_id"`
	Servicetype_id uint   `json:"servicetype_id"`
	Date_start     string `json:"date_start"`
	Date_end       string `json:"date_end"`
}

type Dashboard_company struct {
	Id                uint   `json:"id"`
	Type_company_id   uint   `json:"type_company_id"`
	Company_name      string `json:"company_name"`
	Dashboard_service []Dashboard_service
}

type Dashboard_servicetype struct {
	Id           uint   `json:"id"`
	Service_name string `json:"service_name"`
}

type Dashboard_sbt_tax struct {
	Id         uint   `json:"id"`
	E_tax_name string `json:"e_tax_name"`
	Status     string `json:"status"`
}
