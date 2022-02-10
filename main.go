package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return file
}

func sliceFile(file *os.File) []string {
	var lttrlines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lttrlines = append(lttrlines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lttrlines
}

func Createmap(slice []string) map[int][]string {
	mapslice := make(map[int][]string)
	ascii := 32
	count := 0
	split := 0
	empty := []string{}

	for split < len(slice)-9 {
		split = count * 9
		for i := split; i < split+9; i++ {
			empty = append(empty, slice[i])
		}
		mapslice[ascii] = empty
		empty = []string{}
		ascii++
		count++
	}

	return mapslice
}

func Printascii(mapascii map[int][]string, word string) string {
	toprint := ""
	for i := 1; i <= 8; i++ {
		for _, c := range word {
			toprint += mapascii[int(c)][i]
		}
		toprint += "\n"
	}
	return toprint
}

func Condition(mapascii map[int][]string, a string) {
	splitA := strings.Split(a, "\\n")
	for _, w := range splitA {
		if w != "" {
			fmt.Print(Printascii(mapascii, w))
		} else {
			fmt.Println()
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 3 {
		// error for missing or extra arguments
		fmt.Print("Usage: go run . [STRING] [BANNER] EX: go run . something standard")
	}
	if args[2] == "standard" || args[2] == "shadow" || args[2] == "thinkertoy" {
		// get banner name form arguments
		filename := "banner/" + args[2] + ".txt"
		// open banner file
		file := openFile(filename)
		// slice banner by ascii art
		lttrlines := sliceFile(file)
		// create map of ascii art
		mapslice := Createmap(lttrlines)
		// get string from arguments and print as ascii art
		Output := os.Args[1]
		Condition(mapslice, Output)
		file.Close()
	} else {
		// error for incorrect banner name
		fmt.Print("Usage: go run . [STRING] [BANNER] EX: go run . something standard")
	}
}
