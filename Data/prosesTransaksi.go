package data

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func FindItemByName(name string) *Item {
	for _, item := range DataItems {
		if strings.EqualFold(item.Name, name) {
			return &item
		}
	}
	return nil
}

func ProsesTransaksi(transaksi []Item) Invoice {
	var total, tax, grandTotal int

	for _, item := range transaksi {
		subTotal := item.Quantity * item.Price
		total += subTotal
	}

	tax = total * 10 / 100

	grandTotal = total + tax

	invoice := Invoice{
		IdInvoice:  generateInvoiceID(),
		Date:       time.Now(),
		Items:      transaksi,
		Total:      total,
		Tax:        tax,
		GrandTotal: grandTotal,
	}
	return invoice
}

func generateInvoiceID() int {
	return rand.Intn(1000) + 1000
}

func SaveInvoice(invoice Invoice, filename string) error {
	content := fmt.Sprintf("Invoice Number: %d\n", invoice.IdInvoice)
	content += fmt.Sprintf("Date: %s\n", invoice.Date.Format("02 January 2006"))
	content += "-------------------------------\n"
	content += "Items:\n"
	for _, item := range invoice.Items {
		content += fmt.Sprintf("- %s (x%d) - Rp. %d\n", item.Name, item.Quantity, item.Price)
	}
	content += "-------------------------------\n"
	content += fmt.Sprintf("Total: Rp. %.2f\n", float64(invoice.Total)/100)
	content += fmt.Sprintf("Tax: Rp. %.2f\n", float64(invoice.Tax)/100)
	content += fmt.Sprintf("Grand Total: Rp. %.2f\n", float64(invoice.GrandTotal)/100)
	path := "Transaksi/" + filename
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddItem(id, quantity, price int, name string) string {
	if id < 1 || quantity < 1 || name == "" || price < 0 {
		return "Inputan Salah"
	}
	var item = Item{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
	DataItems = append(DataItems, item)
	return fmt.Sprintf("Barang %s Berhasil Ditambahkan", name)
}
