package main

import "fmt"

func main() {
	const usd_eur float64 = 0.87
	const usd_rub float64 = 81.25
	var eur_rub float64 = usd_rub / usd_eur
	fmt.Println(eur_rub)
}
