package model

type Postmon struct {
	Cep          string `json:"cep"`
	State        string `json:"estado"`
	City         string `json:"cidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}
