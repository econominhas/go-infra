package econominhas

import (
	"github.com/awslabs/goformation/v7/cloudformation"

	"github.com/econominhas/infra/internal/clouds/aws"
	"github.com/econominhas/infra/internal/clouds/providers"
)

func Webapp() ([]byte, error) {
	stackId := "webapp"

	template := cloudformation.NewTemplate()

	cloud := aws.NewAws(stackId)

	// Global Stack

	globalStack := aws.NewAws(PROJECT_ID)

	globalDnsRef := globalStack.Dns().GetMainRef()

	// Website

	website := cloud.Website()

	website.CreateStatic(template, &providers.CreateStaticWebsiteInput{
		Name:   stackId,
		DnsRef: globalDnsRef,
	})

	return template.YAML()
}
