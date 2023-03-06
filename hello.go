package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const n_monitoring = 3
const delay = 5

func main() {

	showIntroduction()
	registerLog("site", false)

	ShowMenu()
	comando := readComando()

	switch comando {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Showing Logs...")
		showLogs()
	case 0:
		fmt.Println("Exiting Program...")
		os.Exit(0)
	default:
		fmt.Println("Do not know this command")
		os.Exit(-1)
	}

}

func showIntroduction() {
	name := "Guilherme" /*var nome string = "Guilherme"*/
	version := 1.2      /* var age float32 = 1.2 */
	fmt.Println("Hello World! This is", name)
	fmt.Println("The version of this program is", version)
}

func readComando() int {
	var comando_read int
	fmt.Scan(&comando_read)
	fmt.Println("Chose", comando_read)
	return comando_read
}

func ShowMenu() {
	fmt.Println("1 - Initiate Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")
}

func initMonitoring() {
	fmt.Println("Monitoring...")

	sites := readExternal()

	for i := 0; i < n_monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Test", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Print("")

}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error occured.", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was correctly loaded")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "has problems. Status Code",
			resp.StatusCode)
		registerLog(site, false)
	}
}

func readExternal() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error occured:", err)
	}

	reader := bufio.NewReader(file)
	for {
		site, err := reader.ReadString('\n')
		site = strings.TrimSpace(site)
		if site != "" {
			sites = append(sites, site)
		}

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt",
	os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site +
	"- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs(){
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
