package main

import (
	"brute-force-dir/colorText"
	"brute-force-dir/readWordlist"
	"brute-force-dir/request"
	"fmt"
	"os"
	"strconv"
)

type IReturnMessage interface {
	ReturnMessage([]string) []string
}

func DoReturnMessage(m IReturnMessage, message []string) []string {
	return message
}

func main() {

	if len(os.Args) < 2 {
		colorText.Red("None arguments are given, use --help for options")
		os.Exit(1)
	}

	if os.Args[1] == "--help" {
		fmt.Println("url - Destination url to be tested\nExample: go run main.go http://domain")
		os.Exit(0)
	} else {
		colorText.Green("[+] Starting program...")
		url := os.Args[1]

		list := readWordlist.ReadFile("wordlist.txt")

		for _, path := range list["lines"] {

			path = "/" + path

			req := request.Target{Url: url, Path: path}
			resp, err := req.Action("GET", url, path)
			returnSlice := []string{url, path, strconv.Itoa(resp.StatusCode)}

			message := DoReturnMessage(&req, returnSlice)

			if err != nil {
				fmt.Println("Error:", err, "requesting to site:", url)
			}

			colorText.White("Running: " + message[0] + message[1] + " | Code: " + message[2])

		}
	}

}
