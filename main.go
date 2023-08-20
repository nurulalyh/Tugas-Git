package main

import "fmt"

func main() {
	developerName := "Nurul Aliyah"
	copyRight := "August 2023"

	fmt.Println("Nama: ", developerName)
	fmt.Println(copyRight)
	fmt.Println("")

	fmt.Println("\t\tKumpulan Program Sederhana")
}

func menu() {
	var menu int

	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Menu: ")
	fmt.Println("1. Menghitung Luas Trapesium")
	fmt.Println("2. Menentukan Bilangan Ganjil-Genap")
	fmt.Println("3. Menentukan Grade")
	fmt.Println("4. Menentukan Faktor Bilangan")
	fmt.Println("0. Keluar")

	fmt.Print("Pilihan Menu : ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		trapesium()
	case 2:
		ganjilgenap()
	case 3:
		grade()
	case 4:
		faktorbil()
	case 0:
		fmt.Println("------------------------------------------------------------------")
		fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
		fmt.Println("------------------------------------------------------------------")
		break
	default:
		fmt.Println("Input Invalid")
		main()
	}
}

func trapesium() {
	var a, b, t, L float32
	var tryAgain string

	fmt.Println("\t\tProgram Menghitung Luas Trapesium")
	fmt.Println("------------------------------------------------------------------")
	fmt.Print("Masukkan panjang sisi a (cm): ")
	fmt.Scanln(&a)
	fmt.Println("")
	fmt.Print("Masukkan panjang sisi b (cm): ")
	fmt.Scanln(&b)
	fmt.Println("")
	fmt.Print("Masukkan tinggi trapesium (cm): ")
	fmt.Scanln(&t)
	fmt.Println("------------------------------------------------------------------")

	L = 0.5 * (a + b) * t
	fmt.Printf("Luas Trapesium adalah %v cm\n", L)
	fmt.Println("------------------------------------------------------------------")

	fmt.Print("Apakah Anda ingin menghitung luas trapesium lagi? (y/n): ")
	fmt.Scanln(&tryAgain)
	fmt.Println("------------------------------------------------------------------")

	if tryAgain != "y" && tryAgain != "Y" {
		fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
		fmt.Println("------------------------------------------------------------------")
	} else {
		main()
	}
}

func ganjilgenap() {
	var bil int16
	var tryAgain string

	fmt.Println("\tProgram Menentukan Bilangan Ganjil atau Genap")
	fmt.Println("-------------------------------------------------------------")
	fmt.Print("Masukkan Bilangan Bulat: ")
	fmt.Scanln(&bil)
	fmt.Println("")
	fmt.Println("-------------------------------------------------------------")

	if bil%2 == 0 {
		fmt.Printf("%v adalah bilangan genap\n", bil)
		fmt.Println("-------------------------------------------------------------")
	} else {
		fmt.Printf("%v adalah bilangan ganjil\n", bil)
		fmt.Println("-------------------------------------------------------------")
	}

	fmt.Print("Apakah Anda ingin mencoba lagi? (y/n): ")
	fmt.Scanln(&tryAgain)
	fmt.Println("-------------------------------------------------------------")

	if tryAgain != "y" && tryAgain != "Y" {
		fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
		fmt.Println("-------------------------------------------------------------")
	} else {
		main()
	}

}

func grade() {
	var nilai float32
	var tryAgain string

	fmt.Println("\t\tProgram Menentukan Grade Nilai")
	fmt.Println("-------------------------------------------------------------")
	fmt.Print("Masukkan Nilai (ex : 90.00): ")
	fmt.Scanln(&nilai)
	fmt.Println("")
	fmt.Println("-------------------------------------------------------------")

	if nilai <= 100 && nilai >= 80 {
		fmt.Printf("Nilai anda %v, Grade anda adalah A\n", nilai)
		fmt.Println("-------------------------------------------------------------")
	} else if nilai < 80 && nilai >= 65 {
		fmt.Printf("Nilai anda %v, Grade anda adalah B\n", nilai)
		fmt.Println("-------------------------------------------------------------")
	} else if nilai < 65 && nilai >= 50 {
		fmt.Printf("Nilai anda %v, Grade anda adalah C\n", nilai)
		fmt.Println("-------------------------------------------------------------")
	} else if nilai < 50 && nilai >= 35 {
		fmt.Printf("Nilai anda %v, Grade anda adalah D\n", nilai)
		fmt.Println("-------------------------------------------------------------")
	} else if nilai < 35 && nilai >= 0 {
		fmt.Printf("Nilai anda %v, Grade anda adalah E\n", nilai)
		fmt.Println("-------------------------------------------------------------")
	} else {
		fmt.Println("Nilai Invalid")
		fmt.Println("-------------------------------------------------------------")
	}

	fmt.Print("Apakah Anda ingin mencoba lagi? (y/n): ")
	fmt.Scanln(&tryAgain)
	fmt.Println("-------------------------------------------------------------")

	if tryAgain != "y" && tryAgain != "Y" {
		fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
		fmt.Println("-------------------------------------------------------------")
	} else {
		main()
	}
}

func faktorbil() {
	var bil int
	var tryAgain string

	fmt.Println("\t\tProgram Mencetak Faktor Bilangan")
	fmt.Println("----------------------------------------------------------------")

	fmt.Print("Masukkan bil: ")
	fmt.Scanln(&bil)
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("Faktor dari %d adalah: \n", bil)
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
	fmt.Println("----------------------------------------------------------------")

	fmt.Print("Apakah Anda ingin mencoba lagi? (y/n): ")
	fmt.Scanln(&tryAgain)
	fmt.Println("----------------------------------------------------------------")

	if tryAgain != "y" && tryAgain != "Y" {
		fmt.Println("Terima Kasih Telah Menggunakan Program Ini 😊")
		fmt.Println("----------------------------------------------------------------")
	} else {
		main()
	}
}
