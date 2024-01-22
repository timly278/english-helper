package service

import (
	"fmt"
	"os"
	"strings"
)

const (
	SEPARATE_STRING = "\n- "
)

func GetRandomDash(data string) string {
	dashes := strings.Split(string(data), SEPARATE_STRING)

	return dashes[RandomInt(1, len(dashes))]
}

func ReadFile(path string) []byte{
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error read file: %v\n", err)
		os.Exit(1)
	}
	return data
}

func PrintOut(data string){
	fmt.Printf("\n******************************\n\n")
	fmt.Println(data)
	fmt.Printf("\n******************************\n")
}