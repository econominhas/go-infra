package econominhas

import (
	"github.com/awslabs/goformation/v7/cloudformation"

	"github.com/econominhas/infra/internal/clouds/aws"
	"github.com/econominhas/infra/internal/clouds/providers"
)

func Api() ([]byte, error) {
	stackId := "api"

	template := cloudformation.NewTemplate()

	cloud := aws.NewAws(stackId)

	// Global Stack

	// globalStack := aws.NewAws(PROJECT_ID)

	// globalDnsRef := globalStack.Dns().GetMainRef()

	// Api

	api := cloud.Api()

	api.CreateBasic(template, &providers.CreateBasicApiInput{})

	return template.YAML()
}
