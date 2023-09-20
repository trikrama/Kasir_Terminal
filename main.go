package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	data "github.com/trikrama/Latihan_Kasir_App/Data"
	helper "github.com/trikrama/Latihan_Kasir_App/Helper"
)

func main() {

	for {
		helper.ClearScreen()
		fmt.Println("Selamat Datang")
		fmt.Println("\nPilih menu:")
		fmt.Println("1. Transaksi")
		fmt.Println("2. List Barang")
		fmt.Println("3. Tambah Barang")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var jumlah int
			var transaksi []data.Item
			for {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Masukkan Nama Barang (kosongkan atau Enter untuk selesai): ")
				name, _ := reader.ReadString('\n')
				name = strings.TrimSpace(name)
				if name == "" {
					break
				}
				fmt.Print("Masukkan Jumlah: ")
				fmt.Scanln(&jumlah)
				item := data.FindItemByName(name)
				if item != nil {
					if jumlah > item.Quantity {
						fmt.Println("Jumlah Barang Tidak Mencukupi")
					}
					transaksi = append(transaksi, data.Item{
						ID:       item.ID,
						Name:     item.Name,
						Quantity: jumlah,
						Price:    item.Price,
					})
					for i := range data.DataItems {
						if data.DataItems[i].ID == item.ID {
							data.DataItems[i].Quantity -= jumlah
							break
						}
					}
					fmt.Println("Item Di tambahkan Ke Transaksi")
				} else {
					fmt.Println("Data tidak Ditemukan")
				}
			}
			invoice := data.ProsesTransaksi(transaksi)
			fmt.Println("Transaksi Selesai!")
			fmt.Println("ID Invoice:", invoice.IdInvoice)
			fmt.Println("Tanggal:", invoice.Date.Format("02 January 2006"))
			fmt.Println("Total: Rp.", invoice.Total)
			fmt.Println("Pajak: Rp.", invoice.Tax)
			fmt.Println("Total Keseluruhan: Rp.", invoice.GrandTotal)
			err := data.SaveInvoice(invoice, "file_transaksi_"+strconv.Itoa(invoice.IdInvoice)+".txt")
			if err != nil {
				log.Fatal(err)
			}
			helper.ShowMenu()

		case "2":
			helper.ClearScreen()
			data.ListBarang()
			helper.ShowMenu()

		case "3":
			helper.ClearScreen()
			var id, quantity, price int
			var name string
			fmt.Print("Masukkan ID: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan Nama Barang: ")
			fmt.Scanln(&name)
			fmt.Print("Masukkan Jumlah Barang: ")
			fmt.Scanln(&quantity)
			fmt.Print("Masukkan Harga: ")
			fmt.Scanln(&price)
			fmt.Println(data.AddItem(id, quantity, price, name))
			helper.ShowMenu()
		case "4":
			fmt.Println("Terima kasih telah Berbelanja.")
			return
		default:
			fmt.Println("Menu Tidak Tersedia")
			helper.Delay(3)
		}
	}
}
