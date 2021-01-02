package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insorsiyon(arr []int, puh int) {
	var i, key, j, x, y int
	x = 1
	for i = 1; i < puh; i++ {
		y = 1
		key = arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
			fmt.Printf("Ara-%6v. Adım : ", y)
			fmt.Println(arr)
			y++
		}
		arr[j+1] = key
		fmt.Printf("\n%10v. ADIM : ", x)
		fmt.Println(arr)
		fmt.Println("")
		x++
	}
}
func main() {
	var arr [20]int
	for i := 0; i < 20; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Print("Sıralanacak dizi: ")
	fmt.Println(arr)
	start := time.Now()
	insorsiyon(arr[:], 20)
	finish := time.Since(start)
	fmt.Printf("Şu kadan zaman aldı= %s //\n", finish)
	fmt.Scanln()
}
