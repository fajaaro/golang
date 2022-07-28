package main

import "fmt"

func main() {
	pecahanArr := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	pecahan := map[int]int{
		100000: 0,
		50000:  0,
		20000:  0,
		10000:  0,
		5000:   0,
		2000:   0,
		1000:   0,
		500:    0,
		200:    0,
		100:    0,
	}

	var harga int
	bayar := 0
	fmt.Print("Input : ")
	_, err := fmt.Scanf("%d", &harga)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	for _, p := range pecahanArr {
		if harga >= p {
			sisa := harga - bayar
			totalPecahan := sisa / p
			pecahan[p] = totalPecahan
			if totalPecahan > 0 {
				bayar += totalPecahan * p

				fmt.Println("Total Pecahan:", p, "x", totalPecahan)
			}
		}
	}

	if bayar < harga {
		fmt.Println("Total Pecahan: 100 x 1")
		pecahan[100] = 1
	}

	fmt.Println(pecahan)
}
