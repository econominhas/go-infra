package econominhas

import (
	"github.com/awslabs/goformation/v7/cloudformation"

	"github.com/econominhas/infra/internal/clouds/aws"
	"github.com/econominhas/infra/internal/clouds/aws/dns"
	"github.com/econominhas/infra/internal/clouds/aws/vpc"
)

func Global() ([]byte, error) {
	stackId := "econominhas"

	template := cloudformation.NewTemplate()

	cloud := aws.NewAws(stackId, template.Resources)

	// Dns

	cloud.Dns.CreateMain(&dns.CreateMainDnsInput{
		Name:       stackId,
		DomainName: "econominhas.com.br",
	})

	// Vpc

	cloud.Vpc.CreateMain(&vpc.CreateMainVpcInput{
		Name: stackId,
	})

	return template.YAML()
}
