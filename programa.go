package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const tempoEmSegundosDeDelay = 5 * time.Second
const quantidadeDeReposicoes = 5
const caminhoSitesMonitorados = "sites-monitorados.txt"
const caminhoLogs = "logs.txt"

func main() {
	nome, _ := develveNomeEIdade()
	exibeIntroducao(nome)

	for {
		exibeMenu()
		comandoEscolhido := leComando()
		switch comandoEscolhido {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo ...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
		fmt.Println("")
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

	sites := leSitesDoArquivo()

	for i := 0; i < quantidadeDeReposicoes; i++ {
		for i, urlSite := range sites {
			fmt.Println("Testando site", i, ":", urlSite)
			testaSite(urlSite)
		}
		time.Sleep(tempoEmSegundosDeDelay)
		fmt.Println("")
	}
}

func testaSite(urlSite string) {
	response, err := http.Get(urlSite)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", urlSite, "foi carregado com sucesso")
		registraLog(urlSite, true)
	} else {
		fmt.Println("Site:", urlSite, "esta com problemas. Status Code:", response.StatusCode)
		registraLog(urlSite, false)
	}
}

func leSitesDoArquivo() []string {
	arquivo, err := os.ReadFile(caminhoSitesMonitorados)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	linhas := strings.Split(string(arquivo), "\n")
	return linhas
}

func registraLog(site string, status bool) {
	data := time.Now().Format("02/01/2006 15:04:05")
	logCompleto := fmt.Sprintf("Data: %s, site: %s, status: %v\n", data, site, status)
	arquivo, err := os.OpenFile(caminhoLogs, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return
	}

	arquivo.WriteString(logCompleto)

	defer arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile(caminhoLogs)

	if err != nil {
		fmt.Println("Erro ao imprimir logs", err)
		return
	}

	linhas := strings.Split(string(arquivo), "\n")
	for _, linha := range linhas {
		fmt.Println(linha)
	}
}

// func exibeNomes() {
// 	nomes := []string{"Bruno", "Ana", "Marcos"}

// 	fmt.Println("O meu slice contem", len(nomes), "com capacidade", cap(nomes))

// 	nomes = append(nomes, "Anderson")

// 	fmt.Println("O meu slice contem", len(nomes), "com capacidade", cap(nomes))
// 	fmt.Println(nomes)
// }
