package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gophers-latam/small-template/internal"
	"github.com/gophers-latam/small-template/models"
	"github.com/gophers-latam/small-template/services"
	"github.com/gophers-latam/small-template/storage"
)

func main() {

	logger := internal.NewLogger(os.Stdout, log.Ldate|log.Ltime)
	repository := storage.NewMemoryRepository()
	service := services.NewProductService(repository, logger)
	product := &models.Product{
		Name:  "Gophers LATAM t-shirt",
		Price: 20,
	}

	service.Save(product)
	p := service.Get(1)
	fmt.Printf("Gotten product; Id: %d, Name: %s, Price: %f", p.Id, p.Name, p.Price)
}
