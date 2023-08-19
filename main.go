package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
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
}