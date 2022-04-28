package perhitung

import "fmt"

func Averange(data []int) {
	hasilpenjumlahan := 0
	for _, angka := range data {
		hasilpenjumlahan = hasilpenjumlahan + angka
	}
	// decimal := float64(hasilpenjumlahan)
	rataRata := float64(hasilpenjumlahan) / float64(len(data))
	fmt.Println("ini rata-rata : ", rataRata)
}

func LebihBesar(data []int) {
	lebihBesar := []int{}
	for _, angka := range data {
		if angka >= 90 {
			lebihBesar = append(lebihBesar, angka)
		}
	}
	fmt.Print("ini lebih besar 90 :", lebihBesar)
}
