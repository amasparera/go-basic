package perhitung

import "errors"

func Sum(list []int) int {

	hasil := 0
	for _, angka := range list {
		hasil = hasil + angka
	}
	return hasil
}

func Calculasi(number1 int, number2 int, tipe string) (int, error) {
	// if tipe == "+" {
	// 	return number1 + number2, nil
	// } else if tipe == "-" {
	// 	return number1 - number2, nil
	// } else if tipe == "*" {
	// 	return number1 * number2, nil
	// } else if tipe == "/" {
	// 	return number1 / number2, nil
	// } else {
	// 	return 0, errors.New("operasi salah")
	// }

	var resul int
	var errorSalah error

	switch tipe {
	case "+":
		resul = number1 + number2
	case "-":
		resul = number1 - number2
	case "*":
		resul = number1 * number2
	case "/":
		resul = number1 / number2
	default:
		errorSalah = errors.New("operasi salah")
	}

	return resul, errorSalah

}
