package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()

	comandoEscolhido := leComando()

	// if comando == 1 {
	// 	fmt.Println("Monitorando")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo ...")
	// } else {
	// 	fmt.Println("Não conheço esse comando")
	// }

	switch comandoEscolhido {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo ...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Bruno"
	versao := 0.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("O programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")

	site := "https://www.alura.com.br"
	response, error := http.Get(site)
	if error == nil {
		fmt.Println(response.StatusCode)
	} else {
		fmt.Println(error.Error())
	}
}
