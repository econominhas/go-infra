package econominhas

import (
	"github.com/awslabs/goformation/v7/cloudformation"

	"github.com/econominhas/infra/internal/clouds/aws"
	"github.com/econominhas/infra/internal/clouds/providers"
)

func Webapp() ([]byte, error) {
	stackId := PROJECT_ID
	name := "webapp"

	template := cloudformation.NewTemplate()

	cloud := aws.NewAws(stackId)

	// Global Stack

	globalStack := aws.NewAws(PROJECT_ID)

	globalDnsRef := globalStack.Dns().GetMainRef()

	// Deploy User

	iam := cloud.Iam()

	deployUser := iam.CreateDeployUser(template, &providers.CreateDeployUserInput{
		Name: name,
	})

	// Website

	website := cloud.Website()

	website.CreateStatic(template, &providers.CreateStaticWebsiteInput{
		Name:          name,
		DnsRef:        globalDnsRef,
		DeployUserArn: deployUser.Arn,
		FullDomain:    "app.econominhas.com.br",
	})

	return template.YAML()
}
