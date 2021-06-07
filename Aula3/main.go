package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
