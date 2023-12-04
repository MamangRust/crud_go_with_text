package main

import (
	"crud_with_text/repository"
	"crud_with_text/service"
	"fmt"
	"os"
)

func displayMenu() {
	fmt.Println("===== Data Management CLI =====")
	fmt.Println("1. Tampilkan Data")
	fmt.Println("2. Tambah Data")
	fmt.Println("3. Find id Data")
	fmt.Println("4. Update Data")
	fmt.Println("5. Hapus Data")
	fmt.Println("6. Keluar")
	fmt.Println("===============================")
}

func main() {

	filePath := "data.txt"
	repo := repository.NewTextFileRepository(filePath)
	services := service.NewTextFileService(repo)

	for {
		displayMenu()

		var choice string

		fmt.Print("Pilih operasi yang anda lakukan: ")

		fmt.Scanln(&choice)

		switch choice {
		case "1":
			allData := services.ReadAllData()

			fmt.Println("\nData yang tersimpan")

			for idx, data := range allData {
				fmt.Printf("%d: %s\n", idx, data)
			}
		case "2":
			var newData string
			fmt.Print("Masukkan data baru: ")
			fmt.Scanln(&newData)

			services.CreateData(newData)

			fmt.Println("Data Berhasil ditambahkan")

		case "3":
			fmt.Print("Masukkan data id yang dicari: ")
			var findID int

			fmt.Scanln(&findID)

			foundData := services.FindByID(findID)

			if foundData != "" {
				fmt.Printf("Data ditemukkan: %s\n", foundData)
			} else {
				fmt.Println("Data tidak ditemukkan")
			}

		case "4":
			fmt.Print("Masukkan indexs data yang ingin diupdate")

			var index int

			fmt.Scanln(&index)

			var updatedData string

			fmt.Print("Masukkan data baru: ")
			fmt.Scanln(&updatedData)

			services.UpdateData(index, updatedData)

			fmt.Println("Data berhasil diupdate")
		case "5":
			fmt.Print("Masukkan index data yang dihapus: ")

			var index int

			fmt.Scanln(&index)

			services.DeleteData(index)

			fmt.Println("Data berhasil dihapus")

		case "6":
			fmt.Println("Terima kasih! Sampai jumpa!")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak valid. Silahkan pilih operasi yang sesuai")
		}
	}

}
