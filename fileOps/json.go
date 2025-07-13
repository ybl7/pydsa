package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ParseJson(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	d := &Data{}
	json.Unmarshal(b, d)

	fmt.Printf("%+v", d)
}

type Data struct {
	Users    []User    `json:"users"`
	Products []Product `json:"products"`
	Orders   []Order   `json:"orders"`
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	City   string `json:"city"`
	Active bool   `json:"active"`
}
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

type Order struct {
	OrderID   string  `json:"order_id"`
	UserID    int     `json:"user_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	Date      string  `json:"date"`
}
