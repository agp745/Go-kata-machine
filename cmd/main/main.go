package main

import (
	"fmt"
	"os"

	"github.com/agp745/Go-kata-machine/internal/generate"
	// "github.com/agp745/Go-kata-machine/internal/generate"
)

func readFile() {
	data, err := os.ReadFile("./internal/dsa.json")

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(string(data))
}

func main() {
	generate.Generate()
	// readFile()
}
