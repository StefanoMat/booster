package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/stefanoMat/boost/6-full-api/internal/entity"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("iPhone", "Apple iPhone", 500.00)
	require.NoError(t, err)
	productDB := NewProduct(db)
	productDB.Create(product)
	productFound, err := productDB.FindByID(product.ID.String())
	require.Nil(t, err)
	require.Equal(t, "iPhone", product.Name)
	require.Equal(t, 500.00, product.Price)
	require.Equal(t, product.ID, productFound.ID)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("iPhone", "Apple iPhone", 500.00)
	require.NoError(t, err)
	productDB := NewProduct(db)
	productDB.Create(product)
	product.Name = "Galaxy S30"
	productDB.Update(product)
	productFound, err := productDB.FindByID(product.ID.String())
	require.Nil(t, err)
	require.Equal(t, product.Name, productFound.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("iPhone", "Apple iPhone", 500.00)
	require.NoError(t, err)
	productDB := NewProduct(db)
	productDB.Create(product)
	productDB.Delete(product.ID.String())
	_, err = productDB.FindByID(product.ID.String())
	require.Error(t, err)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	require.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)
	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), fmt.Sprintf("Description %d", i), rand.Float64()*100)
		require.NoError(t, err)
		productDB.Create(product)
	}

	products, err := productDB.FindAll(1, 10, "asc")
	require.Nil(t, err)
	require.Len(t, products, 10)
	require.Equal(t, "Product 1", products[0].Name)

}
