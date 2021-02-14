package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const monitors = 3
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()
		option := read()
		makeDecision(option)
	}
}

func makeDecision(option int) {
	switch option {

	case 0:
		fmt.Println("Exiting ...")
		os.Exit(0)
	case 1:
		initMonitoring()
	case 2:
		printLog()
	default:
		messageInvalidOption()
		os.Exit(-1)
	}
}

func messageInvalidOption() {
	fmt.Println("Invalid option")
}

func initMonitoring() {
	fmt.Println("Monitoring ...")

	sites := readFile()

	for i := 0; i < monitors; i++ {
		for _, site := range sites {
			fmt.Println("Testando site:", site)
			checkSite(site)
		}

		time.Sleep(delay * time.Second)
	}
}

func showIntroduction() {
	nome := "Matheus"
	version := 1.1
	fmt.Println("Welcome,", nome)
	fmt.Println("You are using the application with version:", version)
}

func showMenu() {
	fmt.Println("1 - Monitor")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Exit")
	fmt.Println("Select an option")
}

func read() int {
	var option int
	fmt.Scan(&option)
	return option
}

func checkSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Error getting site status")
	}

	fmt.Println("Website:", site, res.Status)
	log(site, res.StatusCode == 200)
}

func readFile() []string {

	workDir, err := os.Getwd()

	var sites []string

	if err != nil {
		fmt.Println("Error getting work directory")
	}

	file, err := os.Open(filepath.Join(workDir, "sites.txt"))

	if err != nil {
		fmt.Println("Error opening file")
	}

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	if err != nil {
		fmt.Println("Error opening file:")
		fmt.Println(err.Error())
	}

	file.Close()

	return sites
}

func log(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error open log")
	}

	now := time.Now().Format("02/01/2006 15:04:05")

	file.WriteString(now + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLog() {

	fmt.Println("Logging ...")

	workDir, err := os.Getwd()

	if err != nil {
		fmt.Println()
	}

	file, err := ioutil.ReadFile(filepath.Join(workDir, "log.txt"))

	fmt.Println(string(file))
}
