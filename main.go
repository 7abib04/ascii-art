package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Error : invalid arguments")
		os.Exit(1)
	}
	txt := args[1]

	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
	ascii(slice, textSlice)

}

func ascii(slice []string, txtslice []string) {
	for j, txt := range txtslice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					fmt.Print(slice[firstLine])
				}
				fmt.Println()
			}
		} else if j != len(txtslice)-1 {
			fmt.Println("")
		}
	}
}

func charValidation(str string) bool {
	slice := []rune(str)
	for _, char := range slice {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
