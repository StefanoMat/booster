package repository

import (
	"github.com/stefanoMat/boost/entity"
	"gorm.io/gorm"
)

type Address struct {
	DB *gorm.DB
}

func NewAddress(db *gorm.DB) *Address {
	return &Address{DB: db}
}

func (a *Address) Create(address *entity.Address) error {
	return a.DB.Create(address).Error
}

func (a *Address) GetByCep(cep string) (*entity.Address, error) {
	var address entity.Address
	if err := a.DB.Where("cep = ?", cep).First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (a *Address) FindAll(page, limit int, sort string) ([]entity.Address, error) {
	var addresses []entity.Address
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = a.DB.Limit(limit).Offset((page - 1) * limit).Order("cep " + sort).Find(&addresses).Error
	} else {
		err = a.DB.Order("created_at" + sort).Find(&addresses).Error
	}
	return addresses, err
}
