package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mvssba/cep-brazilian/model"
	"io/ioutil"
	"net/http"
)

const URL_VIA_CEP = "https://viacep.com.br/ws/%s/json"
const URL_POSTMON = "http://api.postmon.com.br/v1/cep/%s"

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/viacep/:cep", viacep)
	e.GET("/postmon/:cep", postmon)
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

	var address model.ViaCep
	if err = json.Unmarshal(b, &address); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, address)
}

func postmon(c echo.Context) error {
	cep := c.Param("cep")
	url := fmt.Sprintf(URL_POSTMON, cep)

	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	var address model.Postmon
	if err = json.Unmarshal(b, &address); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, address)
}
