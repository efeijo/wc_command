package operations

import (
	"bytes"
	"io"
	"strings"
)

type OperationType int
type Operation func(reader io.Reader) ([]int, error)

const (
	NumOfBytes OperationType = iota
	CountLines
	CountWords
	CountChars
	All
)

var Map = map[OperationType]Operation{
	NumOfBytes: numOfBytes,
	CountLines: numOfLines,
	CountWords: numOfWords,
	CountChars: numOfChars,
	All:        all,
}

func numOfBytes(reader io.Reader) ([]int, error) {
	fileContents, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []int{len(fileContents)}, nil
}

func numOfLines(reader io.Reader) ([]int, error) {
	fileContents, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	count := bytes.Count(fileContents, []byte("\n"))
	return []int{count}, nil
}

func numOfWords(reader io.Reader) ([]int, error) {
	fileContents, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []int{len(strings.Fields(string(fileContents)))}, nil
}

func numOfChars(reader io.Reader) ([]int, error) {
	fileContents, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []int{len(string(fileContents))}, nil
}

func all(reader io.Reader) ([]int, error) {
	b1 := &bytes.Buffer{}
	b2 := &bytes.Buffer{}
	teeReader := io.TeeReader(reader, b1)
	teeReader2 := io.TeeReader(b1, b2)

	chars, _ := numOfChars(teeReader)
	lines, _ := numOfLines(teeReader2)
	words, _ := numOfWords(b2)

	lines = append(lines, chars...)
	lines = append(lines, words...)
	return lines, nil
}
