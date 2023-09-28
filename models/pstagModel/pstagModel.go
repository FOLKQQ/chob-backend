package pstagmodel

type Pstag struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Permis_tag string `json:"permis_tag"`
	Status     string `json:"status"`
	Timestamps string `json:"timestamps"`
}
