package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	friends := []Friend{
		{1, "Alice", "Jakarta", "Software Engineer", "Ingin mempelajari bahasa pemrograman baru"},
		{2, "Bob", "Bandung", "Data Scientist", "Mencari alternatif bahasa pemrograman yang lebih cepat dan efisien"},
		{3, "Lintang", "Surabaya", "UI/UX Designer", "Ingin memperluas skill programming"},
		{4, "Ica", "Semarang", "Full-stack Developer", "Mendapatkan kesempatan belajar bahasa pemrograman baru"},
		{5, "Adam", "Medan", "Product Manager", "Ingin mempercepat proses development dengan bahasa pemrograman yang lebih baik"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <absen>")
		os.Exit(1)
	}

	absen := os.Args[1]
	for _, friend := range friends {
		if fmt.Sprintf("%d", friend.Absen) == absen {
			fmt.Printf("Nama: %s\n", friend.Nama)
			fmt.Printf("Alamat: %s\n", friend.Alamat)
			fmt.Printf("Pekerjaan: %s\n", friend.Pekerjaan)
			fmt.Printf("Alasan memilih kelas Golang: %s\n", friend.Alasan)
			return
		}
	}

	fmt.Println("Data teman tidak ditemukan")
}
