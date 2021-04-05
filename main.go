package main

import (
	"github.com/lucasAzS/legendary-computing-machine/http"
	"github.com/lucasAzS/legendary-computing-machine/model"
	uuid "github.com/satori/go.uuid"
)

func main() {
	produto1 := model.Product{
		ID: uuid.NewV4().String(),
		Name: "Carrinho",
	}

	produto2 := model.Product{
		ID: uuid.NewV4().String(),
		Name: "Boneca da Xuxa",
	}
	produto2.Name = "Boneca da xuxa"

	products := model.Products{}

	products.Add(produto1)
	products.Add(produto2)

	server := http.NewWebServer()
	server.Products = &products

	server.Serve()

}