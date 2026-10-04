package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	nome, _ := develveNomeEIdade()
	fmt.Println(nome)

	exibeIntroducao(nome)

	for {
		exibeMenu()
		comandoEscolhido := leComando()
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
}

func develveNomeEIdade() (string, int) {
	nome := "Bruno"
	idade := 29
	return nome, idade
}

func exibeIntroducao(nome string) {
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
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", response.StatusCode)
	}
}
