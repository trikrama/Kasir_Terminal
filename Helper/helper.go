package helper

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Delay(duration int) {
	for i := duration; i >= 1; i-- {
		fmt.Printf("\r%d seconds...", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r")
}

func ShowMenu() {
	var pilihan string
	for {
		fmt.Println("Klik x untuk Kembali ke Menu Utama")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)
		if pilihan == "x" {
			break
		}
		handlePilihan(pilihan)
	}
}

func handlePilihan(pilihan string) {
	fmt.Printf("Anda memilih pilihan %s\n", pilihan)
}
