package entity

import (
	"errors"
	"time"

	"github.com/stefanoMat/boost/6-full-api/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("ID is required")
	ErrInvalidID       = errors.New("ID is invalid")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID          entity.ID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewProduct(name, description string, price float64) (*Product, error) {
	product := &Product{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
	}
	if err := product.IsValid(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) IsValid() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrPriceIsRequired
	}
	return nil
}
