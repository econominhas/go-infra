package econominhas

import (
	"github.com/awslabs/goformation/v7/cloudformation"

	"github.com/econominhas/infra/internal/clouds/aws"
	"github.com/econominhas/infra/internal/clouds/providers"
)

func Global() ([]byte, error) {
	stackId := "econominhas"

	template := cloudformation.NewTemplate()

	cloud := aws.NewAws(stackId)

	// Dns

	dns := cloud.Dns()

	dns.CreateMain(template, &providers.CreateMainDnsInput{
		Name:       stackId,
		DomainName: "econominhas.com.br",
	})

	// Vpc

	vpc := cloud.Vpc()

	vpc.CreateMain(template, &providers.CreateMainVpcInput{
		Name: stackId,
	})

	return template.YAML()
}
