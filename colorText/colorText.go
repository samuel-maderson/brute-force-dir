package colorText

import "fmt"

func Disable() {
	fmt.Printf("\033[0;m")
}

func Green(message string) {
	fmt.Printf("\033[0;32m%s\n", message)
}

func Red(message string) {
	fmt.Printf("\033[0;31m%s\n", message)
}
