package model

type ViaCep struct {
	Cep          string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	UF           string `json:"uf"`
	DDD          string `json:"ddd"`
}
