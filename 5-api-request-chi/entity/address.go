package entity

import "time"

type Address struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Cep         string    `json:"cep"`
	Logradouro  string    `json:"logradouro"`
	Complemento string    `json:"complemento"`
	Bairro      string    `json:"bairro"`
	Localidade  string    `json:"localidade"`
	UF          string    `json:"uf"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewAddress(cep, logradouro, complemento, bairro, localidade, uf string) *Address {
	return &Address{
		Cep:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Bairro:      bairro,
		Localidade:  localidade,
		UF:          uf,
		CreatedAt:   time.Now(),
	}
}
