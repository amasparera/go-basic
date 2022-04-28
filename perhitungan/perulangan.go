package perhitung

import (
	"fmt"
)

func CetakGanjil(kata string) {
	for index, huruf := range kata {
		if index%2 == 0 {
			fmt.Println("letter :", string(huruf), "index :", index)
		}
	}
}

func CetakFocal(kata string) {
	for _, huruf := range kata {
		abjad := string(huruf)
		switch abjad {
		case "a", "i", "u", "e", "o":
			printHuruf(abjad)

		}
	}
}

func printHuruf(huruf string) {
	fmt.Println("ini huruf :", huruf)
}
