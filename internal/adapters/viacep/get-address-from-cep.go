package viacep

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/adapters"
)

type cepFromViaCep struct {
	Cep         string `json:"cep"`         // PostalCode
	Logradouro  string `json:"logradouro"`  // Part 1 of Line 1
	Complemento string `json:"complemento"` // Part of line 2 (Apartment Number, Extra information)
	Unidade     string `json:"unidade"`     // Part of line 2 (Building)
	Bairro      string `json:"bairro"`      // Part 2 of Line 2 (Block)
	Localidade  string `json:"localidade"`  // City
	Uf          string `json:"uf"`          // Acronym for the State
	Estado      string `json:"estado"`      // State
	Regiao      string `json:"regiao"`      // Country region
	Ibge        string `json:"ibge"`        // ???
	Gia         string `json:"gia"`         // ???
	Ddd         string `json:"ddd"`         // Initial digits of phone numbers
	Siafi       string `json:"siafi"`       // ???
}

func (self *viacepAdapter) GetAddressFromCep(cep string) (*adapters.AddressData, error) {
	req, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}

	var result cepFromViaCep
	err = json.NewDecoder(req.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var extraData *string = nil
	if result.Unidade != "" || result.Complemento != "" {
		extraDataParts := make([]string, 0, 2)

		if result.Complemento != "" {
			extraDataParts = append(extraDataParts, result.Complemento)
		}
		if result.Unidade != "" {
			extraDataParts = append(extraDataParts, result.Unidade)
		}

		if len(extraDataParts) > 0 {
			extraDataComplete := strings.Join(extraDataParts, ", ")
			extraData = &extraDataComplete
		}
	}

	return &adapters.AddressData{
		PostalCode: result.Cep,
		Street:     result.Logradouro,
		Block:      result.Bairro,
		ExtraData:  extraData,
		City:       result.Localidade,
		State:      result.Estado,
		Country:    "BR",
	}, nil
}
