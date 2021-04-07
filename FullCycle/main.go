package main

import "github.com/VictordaSilvaf/FullCycleGoLang/model"

func main() {
	produto1 := model.NewProduct()
	produto1.Name = "Refrigerante"

	produto2 := model.NewProduct()
	produto2.Name = "H2O"

	products := model.Products{}
	products.Add(produto1)
	products.Add(produto2)
}
