package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

const URL_VIA_CEP = "https://viacep.com.br/ws/%s/json"

func main() {
	e := echo.New()
	e.GET("/viacep/:cep", viacep)
	e.Logger.Fatal(e.Start(":8080"))
}

func viacep(c echo.Context) error {
	cep := c.Param("cep")
	url := fmt.Sprintf(URL_VIA_CEP, cep)

	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var address AddressViaCep
	if err = json.Unmarshal(b, &address); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, address)
}

type AddressViaCep struct {
	Cep          string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	UF           string `json:"uf"`
	DDD          string `json:"ddd"`
}
