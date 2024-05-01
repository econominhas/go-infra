package main

import (
	"errors"
	"os"

	"github.com/econominhas/infra/internal/projects/econominhas"
)

func getStack(stackName string) ([]byte, error) {
	switch stackName {
	case "econominhas-global":
		return econominhas.Global()
	}

	return []byte(""), errors.New("stack not found")
}

func main() {
	stack, err := getStack("")
	if err != nil {
		panic(1)
	}

	os.WriteFile("./cloudformation.yaml", stack, 0644)
}
