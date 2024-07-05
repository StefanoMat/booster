package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stefanoMat/boost/entity"
	"github.com/stefanoMat/boost/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("consulta.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Address{})

	r := chi.NewRouter()
	//----//
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Heartbeat("/health"))

	r.Route("/api/cep", func(r chi.Router) {
		r.Get("/", GetCeps)
		r.Get("/{cep}", GetCep)
	})

	//----//
	println("ListenAndServe: Inicializando servidor na porta 3000")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		println("ListenAndServe: Erro ao iniciar o servidor")
		panic(err)
	}
}

func GetCep(rw http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 900*time.Millisecond)
	defer cancel()

	addressFind, err := repository.NewAddress(db).GetByCep(cep)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	if addressFind != nil {
		println("Encontrado no banco de dados")
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		err = json.NewEncoder(rw).Encode(addressFind)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		}
		return
	}

	address, err := BuscaCEP(cep, ctx)
	if err != nil {
		panic(err)
	}
	addressEntity := entity.NewAddress(address.Cep, address.Logradouro, address.Complemento, address.Bairro, address.Localidade, address.Uf)
	err = repository.NewAddress(db).Create(addressEntity)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	err = json.NewEncoder(rw).Encode(address)
	if err != nil {
		panic(err)
	}
}

func GetCeps(rw http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	addresses, err := repository.NewAddress(db).FindAll(pageInt, limitInt, sort)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	json.NewEncoder(rw).Encode(addresses)

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
