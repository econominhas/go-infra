package aws

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/econominhas/infra/internal/clouds/aws/dns"
	"github.com/econominhas/infra/internal/clouds/aws/vpc"
)

type AwsStack struct {
	Dns dns.Deps
	Vpc vpc.Deps
}

func NewAws(stackId string, resources cloudformation.Resources) *AwsStack {
	return &AwsStack{
		Dns: dns.Deps{
			StackId:   stackId,
			Resources: resources,
		},
		Vpc: vpc.Deps{
			StackId:   stackId,
			Resources: resources,
		},
	}
}
