package main

import (
	"bufio"
	"fmt"
	"os"

	"fgo24-go-temp-convert/menu"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== Menu Konversi Suhu ===\n")
		fmt.Println("1. Convert suhu")
		fmt.Println("2. Lihat hasil convert")
		fmt.Println("0. Keluar\n")
		fmt.Print("Masukkan pilihan: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			menu.ConvertSuhu(scanner)
		case "2":
			menu.TampilkanHasil()
		case "0":
			fmt.Println("Terima kasih")
			return
		default:
			fmt.Println("Pilihan tidak ada")
		}
	}
}
