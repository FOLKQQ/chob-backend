package dashboardmodel

type Case struct {
	Id           int
	Type_company string
	Company_name string
	Service_name string
	Date_start   string
	Date_end     string
	Sbt_tax      []Sbt_tax
}
type Sbt_tax struct {
	E_tax_name string `json:"e_tax_name"`
	Status     string `json:"status"`
}

type Statusworks struct {
	Backlog int
	Ready   int
	Doing   int
	Done    int
}
