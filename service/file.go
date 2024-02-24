package service

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	SEPARATE_STRING = "- "
)

// GetRandomDash
func GetRandomDash(data []byte, n int) ([]string, error) {
	temp := strings.Trim(string(data), "\n")
	dashes := strings.Split(temp, SEPARATE_STRING)

	if n > len(dashes) {
		return nil, errors.New("too few data in the setup, please add more")
	}

	islice := RandomIntSlice(n, 0, len(dashes)-1)
	result := make([]string, len(dashes))

	for i, v := range islice {
		result[i] = dashes[v]
	}

	return result, nil
}

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error read file: %v\n", err)
		os.Exit(1)
	}
	return data
}

func PrintOut(data string) {
	fmt.Printf("\n******************************\n\n")
	fmt.Println(data)
	fmt.Printf("\n******************************\n")
}

func RemoveEmptyLine(data []string) []string {

	var result []string
	for _, v := range data {
		if v[0] != '\n' {
			result = append(result, v)
		}
	}

	return result
}
