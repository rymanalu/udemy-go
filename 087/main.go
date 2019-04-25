package main

import "fmt"

func main() {
	provinces := make([]string, 34, 34)
	provinces = []string{
		"Aceh", "Bali", "Banten", "Bengkulu", "Gorontalo", "Jakarta", "Jambi",
		"Jawa Barat", "Jawa Tengah", "Jawa Timur", "Kalimantan Barat",
		"Kalimantan Selatan", "Kalimantan Tengah", "Kalimantan Timur",
		"Kalimantan Utara", "Kepulauan Bangka Belitung", "Kepulauan Riau",
		"Lampung", "Maluku", "Maluku Utara", "Nusa Tenggara Barat",
		"Nusa Tenggara Timur", "Papua", "Papua Barat", "Riau", "Sulawesi Barat",
		"Sulawesi Selatan", "Sulawesi Tengah", "Sulawesi Tenggara",
		"Sulawesi Utara", "Sumatra Barat", "Sumatra Selatan", "Sumatra Utara",
		"Yogyakarta",
	}

	totalProvinces := len(provinces)

	fmt.Println("length:", totalProvinces)
	fmt.Println("capacity:", cap(provinces))

	for i := 0; i < totalProvinces; i++ {
		fmt.Printf("Index: %v\tProvince: %v\n", i, provinces[i])
	}
}
