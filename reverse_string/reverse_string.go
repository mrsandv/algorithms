package reversestring

import (
	"fmt"
)

func ReverseString(s string) string {
	runes := []rune(s) //Rune es el valor int que representa un string
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) //Esta version si respeta los caracteres complejos, como las tildes. pero la mia esta mas bonita y facil :p

	// var reversedString string
	// for i := len(s) - 1; i >= 0; i-- {
	// 	reversedString = reversedString + string(s[i])
	// 	fmt.Println(string(s[i]))
	// }
	// return reversedString
}

func Init() {
	result := ReverseString("Hola, este mensaje es una prueba, ahora con acción")
	fmt.Println(result)
}
