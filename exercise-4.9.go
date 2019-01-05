package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	result := map[string]int{}
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	// remove all special character
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(input, " ")
	fmt.Println(processedString)
	//
	scanner := bufio.NewScanner(strings.NewReader(processedString))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		result[scanner.Text()]++
	}
	fmt.Println(result)

}
