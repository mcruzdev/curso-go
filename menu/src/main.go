package main

import (
	"fmt"
	"net/http"
	"os"
)

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
		showLog()
	default:
		messageInvalidOption()
		os.Exit(-1)
	}
}

func messageInvalidOption() {
	fmt.Println("Invalid option")
}

func showLog() {
	fmt.Println("Logging ...")
}

func initMonitoring() {
	fmt.Println("Monitoring ...")

	site := "https://www.alura.com"

	res, _ := http.Get(site)
	fmt.Println("Website:", site, res.Status)
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
