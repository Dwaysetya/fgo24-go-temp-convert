package menu

import (
	"bufio"
	"fmt"
	"strconv"

	"fgo24-go-temp-convert/calculate"
)

var results []string

func ConvertSuhu(scanner *bufio.Scanner) {
	fmt.Print("Masukkan suhu dalam celcius: ")
	scanner.Scan()
	input := scanner.Text()
	celcius, _ := strconv.Atoi(input)

	fmt.Println("Pilih konversi:")
	fmt.Println("1. Ke Kelvin")
	fmt.Println("2. Ke Fahrenheit")
	fmt.Println("3. Ke Reamur")
	fmt.Print("Pilihan: ")
	scanner.Scan()
	pilihan := scanner.Text()

	var result string
	switch pilihan {
	case "1":
		result = calculate.CkeK(celcius)
	case "2":
		result = calculate.CkeF(celcius)
	case "3":
		result = calculate.CkeR(celcius)
	default:
		fmt.Println("Pilihan tidak ada")
		return
	}

	fmt.Println("Hasil konversi:", result)
	results = append(results, result)
}


func TampilkanHasil() {
	fmt.Println("\n=== Hasil Konversi Suhu Sebelumnya ===\n")
	if len(results) == 0 {
		fmt.Println("Belum ada hasil")
		return
	}

	for i, res := range results {
		fmt.Printf("%d. %s\n", i+1, res)
	}
}
