package main

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Id      int
	Name    string
	Address string
	Phone   string
}

type Product struct {
	gorm.Model
	Id          int
	Name        string
	Description string
	Price       float64
	Count       int
}

type Order struct {
	gorm.Model
	Id         int
	CustomerId int
	ProductId  int
	Quantity   int
	Total      float64
}
