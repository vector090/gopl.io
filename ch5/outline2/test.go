package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune { return r + 1 }

func main() {

//	var x rune = rune(1)
//	fmt.Println(x)

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}
