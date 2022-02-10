package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Get banner name from cmd line
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Print("EX: go run . something standard")
		return
	}

	bName := os.Args[2] + ".txt"

	file, err := os.Open(bName)
	if err != nil {
		log.Fatal(err)
	}

	// file to slice of string by line
	var lttrlines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lttrlines = append(lttrlines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// find a string in a scanned file
	// Splits on newlines by default.
	ln1 := ""
	ln2 := ""
	ln3 := ""
	ln4 := ""
	ln5 := ""
	ln6 := ""
	ln7 := ""
	ln8 := ""

	line := 0
	Output := os.Args[1]
	SlcOutput := []rune(Output)
	dash := strings.Index(Output, `\n`)

	for i := 0; i < len(SlcOutput); i++ {
		if i == dash && dash >= 0 {
			if ln1 != "" && ln2 != "" && ln3 != "" && ln4 != "" && ln5 != "" && ln6 != "" && ln7 != "" && ln8 != "" {
				fmt.Println(ln1)
				fmt.Println(ln2)
				fmt.Println(ln3)
				fmt.Println(ln4)
				fmt.Println(ln5)
				fmt.Println(ln6)
				fmt.Println(ln7)
				fmt.Println(ln8)
			} else {
				fmt.Println()
			}

			ln1 = ""
			ln2 = ""
			ln3 = ""
			ln4 = ""
			ln5 = ""
			ln6 = ""
			ln7 = ""
			ln8 = ""
			dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
			i++
		} else {
			for j, s := range lttrlines {
				line = j
				if len(s) > 0 && s == string(SlcOutput[i]) {
					break
				}

			}

			for i := line + 1; i <= line+8; i++ {
				if i == line+1 {
					ln1 = ln1 + lttrlines[i]
				}
				if i == line+2 {
					ln2 = ln2 + lttrlines[i]
				}
				if i == line+3 {
					ln3 = ln3 + lttrlines[i]
				}
				if i == line+4 {
					ln4 = ln4 + lttrlines[i]
				}
				if i == line+5 {
					ln5 = ln5 + lttrlines[i]
				}
				if i == line+6 {
					ln6 = ln6 + lttrlines[i]
				}
				if i == line+7 {
					ln7 = ln7 + lttrlines[i]
				}
				if i == line+8 {
					ln8 = ln8 + lttrlines[i]
				}

			}
			if i == len(Output)-1 {
				fmt.Println(ln1)
				fmt.Println(ln2)
				fmt.Println(ln3)
				fmt.Println(ln4)
				fmt.Println(ln5)
				fmt.Println(ln6)
				fmt.Println(ln7)
				fmt.Println(ln8)
				ln1 = ""
				ln2 = ""
				ln3 = ""
				ln4 = ""
				ln5 = ""
				ln6 = ""
				ln7 = ""
				ln8 = ""
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		// Handle the error
	}
	file.Close()
}
