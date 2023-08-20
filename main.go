package main

import "fmt"

func main() {
	fmt.Println("Kumpulan Program Sederhana")
}

func menu() {
	var menu int

	fmt.Println("Menu: ")
	fmt.Println("1. Menghitung Luas Trapesium")
	fmt.Println("2. Menentukan Bilangan Ganjil-Genap")
	fmt.Println("3. Menentukan Grade")
	fmt.Println("4. Menentukan Faktor Bilangan")
	fmt.Println("0. Keluar")

	fmt.Print("Pilihan Menu : ")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		fmt.Println("Program Menghitung Luas Trapesium")
	case 2:
		fmt.Println("Program Menentukan Bilangan Ganjil-Genap")
	case 3:
		fmt.Println("Program Menentukan Grade")
	case 4:
		fmt.Println("Program Menentukan Faktor Bilangan")
	case 5:
		fmt.Println("Keluar")
	default:
		fmt.Println("Input Invalid")
	}
}