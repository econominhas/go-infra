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

func GenId(i *GenIdInput) string {
	if i.OmitStage {
		return ToPascal(
			strings.Join([]string{i.Id, os.Getenv("ENV"), i.Name, i.Type}, "-"),
		)
	}

	return ToPascal(
		strings.Join([]string{i.Id, i.Name, i.Type}, "-"),
	)
}
