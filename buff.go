package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	data := "hello what is up in the sky? blooming like a flower ! haha haha! okkayyyy"

	ds := strings.NewReader(data)
	scanner := bufio.NewScanner(ds)

	fmt.Println(scanner)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		line := scanner.Text()
		fmt.Println(line)
	}
}
