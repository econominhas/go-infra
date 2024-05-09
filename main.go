package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

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

func getFullStackName(project string, env string, product string) string {
	if product == "global" {
		return strings.Join([]string{project, product}, "-")
	}

	return strings.Join([]string{project, env, product}, "-")
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

	fullStackName := getFullStackName(project, env, product)
	filePath := "./outputs/" + fullStackName + ".yaml"

	const perm fs.FileMode = 0770
	os.Mkdir("./outputs", perm)
	err = os.WriteFile(filePath, stack, perm)
	if err != nil {
		fmt.Print(err)
		panic(1)
	}

	fmt.Print("Please run the following command:\n")
	fmt.Printf(
		"aws cloudformation deploy --template-file %s --stack-name %s\n",
		filePath,
		fullStackName,
	)
}
