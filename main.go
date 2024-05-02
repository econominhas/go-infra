package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/econominhas/infra/internal/projects/econominhas"
	"github.com/manifoldco/promptui"
)

func getStack(stackName string) ([]byte, error) {
	switch stackName {
	case "econominhas-global":
		return econominhas.Global()
	case "econominhas-webapp":
		return econominhas.Webapp()
	default:
		return []byte(""), errors.New("stack not found")
	}
}

func main() {
	envQuestion := promptui.Select{
		Label: "Select Environment",
		Items: []string{
			"dev",
			"prod",
		},
	}
	_, env, err := envQuestion.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	projectQuestion := promptui.Select{
		Label: "Select Project",
		Items: []string{
			"econominhas",
		},
	}
	_, project, err := projectQuestion.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	productQuestion := promptui.Select{
		Label: "Select Product",
		Items: []string{
			"global",
			"webapp",
		},
	}
	_, product, err := productQuestion.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	os.Setenv("ENV", env)

	stack, err := getStack(project + "-" + product)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	os.WriteFile("./cloudformation.yaml", stack, 0644)
}
