package main

import "fmt"

const usd_eur float64 = 0.87
const usd_rub float64 = 81.25

func main() {
	val, first_val, second_val := valut()
	converter(val, first_val, second_val)
}
func valut() (float64, string, string) {
	var val float64
	var first_val, secind_val string
	fmt.Printf("Введите сумму для перевода: \n")
	fmt.Scan(&val)
	fmt.Printf("Введите из какой валюты и в какую совершить перевод (EUR, RUB, USD): \n")
	fmt.Scan(&first_val, &secind_val)
	return val, first_val, secind_val
}

func converter(val float64, first_val, second_val string) {

	if first_val == "EUR" {
		if second_val == "USD" {
			res := val / usd_eur
			fmt.Printf("Итоговая сумма: %.1f", res)
		}
		if second_val == "RUB" {
			res := (val / usd_eur) * usd_rub
			fmt.Printf("Итоговая сумма: %.1f", res)
		}
	}
	if first_val == "USD" {
		if second_val == "EUR" {
			res := val * usd_eur
			fmt.Printf("Итоговая сумма: %.1f", res)
		}
		if second_val == "RUB" {
			res := val * usd_rub
			fmt.Printf("Итоговая сумма: %.1f", res)
		}
	}
	if first_val == "RUB" {
		if second_val == "EUR" {
			res := (val / usd_rub) * usd_eur
			fmt.Printf("Итоговая сумма: %.1f", res)
		}
		if second_val == "USD" {
			res := val / usd_rub
			fmt.Printf("Итоговая сумма: %.1f", res)
		}
	}
}
