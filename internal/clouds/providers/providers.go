package providers

import (
	"github.com/awslabs/goformation/v7/cloudformation"
)

// Vpc

type CreateMainVpcInput struct {
	Name string
}

type Vpc interface {
	CreateMain(t *cloudformation.Template, i *CreateMainVpcInput)
}

// Dns

type CreateMainDnsInput struct {
	Name       string
	DomainName string
}

type Dns interface {
	CreateMain(t *cloudformation.Template, i *CreateMainDnsInput)
}
