package repository

import (
	"fmt"
	"testing"

	"github.com/stefanoMat/boost/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindAllAddresses(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"))
	require.NoError(t, err)
	db.AutoMigrate(&entity.Address{})
	addressRepository := NewAddress(db)
	for i := 1; i < 25; i++ {
		address := entity.NewAddress(fmt.Sprintf("cep %d", i), fmt.Sprintf("rua %d", i), fmt.Sprintf("Complemento %d", i), fmt.Sprintf("bairro %d", i), fmt.Sprintf("localidade %d", i), fmt.Sprintf("uf %d", i))
		addressRepository.Create(address)
	}

	addresses, err := addressRepository.FindAll(1, 10, "asc")
	require.Nil(t, err)
	assert.Len(t, addresses, 10)
	require.Equal(t, "cep 1", addresses[0].Cep)
}
