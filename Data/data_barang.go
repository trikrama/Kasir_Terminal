package data

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    int
}
type Invoice struct {
	IdInvoice  int
	Date       time.Time
	Items      []Item
	Total      int
	Tax        int
	GrandTotal int
}

var DataItems = []Item{
	{ID: 1, Name: "Apel", Quantity: 5, Price: 2500},
	{ID: 2, Name: "Roti", Quantity: 3, Price: 1750},
	{ID: 3, Name: "Susu", Quantity: 2, Price: 4500},
	{ID: 4, Name: "Kentang", Quantity: 10, Price: 750},
	{ID: 5, Name: "Ayam", Quantity: 2, Price: 12500},
	{ID: 6, Name: "Pisang", Quantity: 4, Price: 1200},
	{ID: 7, Name: "Teh", Quantity: 8, Price: 2000},
	{ID: 8, Name: "Telur", Quantity: 6, Price: 500},
	{ID: 9, Name: "Nasi", Quantity: 2, Price: 1000},
	{ID: 10, Name: "Ikan", Quantity: 3, Price: 8750},
	{ID: 11, Name: "Tomat", Quantity: 7, Price: 800},
	{ID: 12, Name: "Kacang", Quantity: 5, Price: 1500},
	{ID: 13, Name: "Sosis", Quantity: 4, Price: 3250},
	{ID: 14, Name: "Kue", Quantity: 6, Price: 2500},
	{ID: 15, Name: "Minyak", Quantity: 3, Price: 5000},
	{ID: 16, Name: "Beras", Quantity: 4, Price: 9000},
	{ID: 17, Name: "Gula", Quantity: 5, Price: 2000},
	{ID: 18, Name: "Kopi", Quantity: 2, Price: 3500},
	{ID: 19, Name: "Coklat", Quantity: 3, Price: 4000},
	{ID: 20, Name: "Keju", Quantity: 4, Price: 6000},
}

func ListBarang() {
	fmt.Println("Daftar Barang:")
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%-10s %-10s %-10s %-10s\n", "Id", "Nama", "Jumlah", "Harga")
	fmt.Println(strings.Repeat("-", 45))
	for _, item := range DataItems {
		fmt.Printf("%-10s %-10s %-10s %-10s\n", strconv.Itoa(item.ID), item.Name, "  "+strconv.Itoa(item.Quantity), "Rp. "+strconv.Itoa(item.Price))
	}
	fmt.Println(strings.Repeat("-", 45))
}
