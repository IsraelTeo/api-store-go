package model

import "gorm.io/gorm"

type Stock uint

type Product struct {
	gorm.Model
	Name  string  `json:"name" validate:"required,min=2,max=50"`
	Mark  string  `json:"mark" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock `json:"stock" validate:"gt=0"`
}

//Disminuir stock
//Para el metodo stock:
//Parametros: Necesita la cantidad que hay de un producto -> esta cantidad es un uint 
//Para obtener la cantidad en uint se debe contar 1 a 1 la lista de un producto
//Para disminuir la cantidad en 1 o en N, se requiere saber la lista de productos de una Venta
//Esa cantidad se debe disminuir a la cantidad total de productos del stock
//Retorno:
