package main

import (
	"github.com/lucasAzS/legendary-computing-machine/model"
)

func main() {
	produto1 := model.NewProduct()
	produto1.Name = "Carrinho"

	produto2 := model.NewProduct()
	produto2.Name = "Boneca da xuxa"

	products := model.Products{}

	products.Add(produto1)
	products.Add(produto2)
}