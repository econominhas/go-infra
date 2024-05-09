package utils

import (
	"os"
	"strings"
)

type GenIdInput struct {
	Id        string
	Type      string
	Name      string
	OmitStage bool
}

type GenIdOutput struct {
	Id   string
	Name string
}

func GenId(i *GenIdInput) *GenIdOutput {
	var name string

	if i.OmitStage {
		name = strings.Join([]string{i.Id, i.Name, i.Type}, "-")
	} else {
		name = strings.Join([]string{i.Id, os.Getenv("ENV"), i.Name, i.Type}, "-")
	}

	return &GenIdOutput{
		Id:   ToPascal(name),
		Name: name,
	}
}
