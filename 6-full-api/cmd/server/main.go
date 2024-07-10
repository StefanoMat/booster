package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/stefanoMat/boost/6-full-api/configs"
	"github.com/stefanoMat/boost/6-full-api/internal/entity"
	"github.com/stefanoMat/boost/6-full-api/internal/infra/database"
	"github.com/stefanoMat/boost/6-full-api/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.Load(".")
	if err != nil {
		panic(err)
	}
	println(config.DBName)

	db, err := gorm.Open(sqlite.Open("datasource.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Post("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", r)

}
