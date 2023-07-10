package main

import "fmt"

func main() {
	fmt.Print("Digite seu nome: ")
	var nome string
	fmt.Scanln(&nome)

	fmt.Printf("Olá, %s! Bem-vindo ao mundo do Go.\n", nome)
}
