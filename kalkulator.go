package main

import (
	"fmt"
)

func main() {
	var operasi string
	var bil1, bil2 float64

	fmt.Println("Pilih Operasi! Perkalian, Pembagian, Penambahan, Pengurangan.")
	fmt.Scanln(&operasi)

	if operasi != "*" && operasi != "/" && operasi != "+" && operasi != "-" {
		fmt.Println("Kalkulator tidak mengerti input user!")

		return
	}
	fmt.Print("Bilangan pertama: ")
	fmt.Scanln(&bil1)
	fmt.Print("Bilangan kedua: ")
	fmt.Scanln(&bil2)

	switch operasi {
	case "*":
		fmt.Print("Hasil: ", bil1*bil2)
	case "/":
		fmt.Print("Hasil: ", bil1/bil2)
	case "+":
		fmt.Print("Hasil: ", bil1+bil2)
	case "-":
		fmt.Print("hasil: ", bil1-bil2)
	}

}
