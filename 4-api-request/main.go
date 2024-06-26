package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type AddressViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel()
	address, err := BuscaCEP("42700-130", ctx)
	if err != nil {
		panic(err)
	}
	println(address.Localidade)

}

func BuscaCEP(cep string, ctx context.Context) (*AddressViaCep, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var address AddressViaCep
	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
