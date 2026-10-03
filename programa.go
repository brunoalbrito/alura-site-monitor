package main

import "fmt"

func main() {
	nome := "Bruno"
	versao := 0.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("O programa esta na versao", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escondido foi", comando)
}
