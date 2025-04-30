package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringCycles = 3
const delay = 1

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		command := lerComando()

		switch command {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	var nome string = "Cezar"
	var idade int            // Se não tem valor, o valor será zero (Branco para string)
	var versao float32 = 1.1 //Float caso vc remova o tipo na declaração, ele vai assumir o maior

	fmt.Println("Hello, world.", nome, "Sua idade é", idade)
	fmt.Println("Current version", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("")
	fmt.Println("Selected option:", command)

	return command
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := readSitesFromFile()

	// fmt.Println(sites)

	for i := 0; i < monitoringCycles; i++ {
		for i, site := range sites {
			fmt.Println("Testing ", i, ":", site)
			testSite(site)
		}

		fmt.Println("====================================")
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")
}

func testSite(site string) {
	// site = "https://httpbin.org/status/404" //(Workaround)
	// Skipping error for now
	resp, err := http.Get(site)

	verifyError(err)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		writeLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		writeLog(site, false)
	}
}

func readSitesFromFile() []string {
	file, err := os.Open("sites.txt")
	// file, err := os.ReadFile("sites.txt")

	verifyError(err)

	reader := bufio.NewReader(file)

	var sites []string

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		} else {
			verifyError(err)
		}
	}

	file.Close()

	fmt.Println(sites)
	return sites

}

func verifyError(err error) {
	if err != nil {
		fmt.Println("Error!:", err)
	}
}

func writeLog(site string, onlineStatus bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	verifyError(err)

	file.WriteString(
		time.Now().Format("02/01/2006 15:04:05") + " - " +
			site + "- online: " + strconv.FormatBool(onlineStatus) + "\n")

	file.Close()
}

func imprimeLogs() {
	file, err := os.ReadFile("log.txt")

	verifyError(err)

	fmt.Println(string(file))

}
