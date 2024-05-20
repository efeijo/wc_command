package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"wc_command/internal/operations"
)

func main() {
	countBytes := flag.Bool("c", false, "-c")
	countLines := flag.Bool("l", false, "-l")
	countWords := flag.Bool("w", false, "-w")
	countChars := flag.Bool("m", false, "-m")

	flag.Parse()

	fs := flag.Args()

	var operation operations.Operation

	switch {
	case *countBytes:
		operation = operations.Map[operations.NumOfBytes]
	case *countLines:
		operation = operations.Map[operations.CountLines]
	case *countWords:
		operation = operations.Map[operations.CountWords]
	case *countChars:
		operation = operations.Map[operations.CountChars]
	default:
		operation = operations.Map[operations.All]
	}

	file, err := os.Open(fs[0])
	if err != nil {
		log.Fatal(err)
	}

	c, err := operation(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prettyPrint(fs[0], c...))
}

func prettyPrint(fileName string, nums ...int) string {
	var resString strings.Builder
	for _, n := range nums {
		resString.WriteString(fmt.Sprintf("%d\t", n))
	}
	resString.WriteString(fmt.Sprintf("%s", fileName))
	return resString.String()
}
