package main

import (
	"fmt"
	"os"
)

func printpenguin() {

	orangepenguin := `   			
		
		   .---.
		   |o_o  |
		   |:_/  |
		 //   \  \
		(|     |  )
		/'\_    /' \
		\___)==(___/ `

	fmt.Println(orangepenguin)
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("error: bitte geben sie ein Argument ein")
		os.Exit(0)
	}

	textinput := os.Args[1]

	for k := 0; k < len(textinput); k++ {
		fmt.Printf("_____")
	}
	fmt.Println()
	fmt.Println(textinput)

	for k := 0; k < len(textinput); k++ {
		fmt.Printf("_____")

	}

	printpenguin()

}
