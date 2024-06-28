package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

// func (a *AddressViaCep) String() string {
// 	return fmt.Sprintf("%s, %s - %s, %s", a.Logradouro, a.Localidade, a.Uf, a.Bairro)
// }

func main() {
	r := chi.NewRouter()
	//----//
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Heartbeat("/health"))

	r.Route("/api/cep", func(r chi.Router) {
		r.Get("/{cep}", GetCep)
	})

	//----//
	println("ListenAndServe: Inicializando servidor na porta 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		println("ListenAndServe: Erro ao iniciar o servidor")
		panic(err)
	}
}

func GetCep(rw http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel()
	//Verificar se o CEP consultado existe na nossa base de dados.
	// Se existir: Retornar os dados do CEP da base.
	// Se nao existir:
	//Salvar em uma tabela do banco de dados chamada address | id int-primary key, cep - unique-not null, logradouro - string, complemento-string, bairro-string, localidade- string - not null, uf string-not null

	//sqlite3
	//mysql no docker
	//Repository
	//Consultar módulo 1 para ver sqlite e queries SQL.

	address, err := BuscaCEP(cep, ctx)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(rw).Encode(address)
	if err != nil {
		panic(err)
	}
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
