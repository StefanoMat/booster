package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	name := "iPhone"
	price := 500.00
	description := "Apple iPhone"
	product, err := NewProduct(name, description, price)
	require.Nil(t, err)
	require.NotNil(t, product)
	require.Equal(t, name, product.Name)
	require.Equal(t, description, product.Description)
	require.Equal(t, price, product.Price)
}

func TestProductNameIsRequired(t *testing.T) {
	product, err := NewProduct("", "Apple iPhone", 500.00)
	require.NotNil(t, err)
	require.Nil(t, product)
	require.Equal(t, ErrNameIsRequired, err)
}

func TestProductInvalidPrice(t *testing.T) {
	product, err := NewProduct("iPhone", "Apple iPhone", -10.00)
	require.NotNil(t, err)
	require.Nil(t, product)
	require.Equal(t, ErrPriceIsRequired, err)

	product, err = NewProduct("iPhone", "Apple iPhone", 0.00)
	require.NotNil(t, err)
	require.Nil(t, product)
	require.Equal(t, ErrPriceIsRequired, err)
}
