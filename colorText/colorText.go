package colorText

import "fmt"

func Green(message string) {
	fmt.Printf("\033[0;32m%s\n", message)
}

func Red(message string) {
	fmt.Printf("\033[0;31m%s\n", message)
}

func White(message string) {
	fmt.Printf("\033[0;37m%s\n", message)
}
