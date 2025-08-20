package create

type CreateInputDto struct {
	Name     string   `json:"name"`
	Document string   `json:"document"`
	Contacts []string `json:"contacts"`
	PersonType int      `json:"personType"`
}

type CreateOutputDto struct {
	Uuid string `json:"uuid"`
}