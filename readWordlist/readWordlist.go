package readWordlist

import (
	"bufio"
	"fmt"
	"os"
)

type WordList struct {
	lines []string
}

func ReadFile(filename string) map[string][]string {

	var wl *WordList
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err, " when opening:", file)
	}

	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)

	wl = new(WordList)
	for fileScan.Scan() {
		wl.lines = append(wl.lines, fileScan.Text())
	}

	return map[string][]string{
		"lines": wl.lines,
	}
}
