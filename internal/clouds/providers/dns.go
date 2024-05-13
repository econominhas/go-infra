package providers

import "github.com/awslabs/goformation/v7/cloudformation"

type CreateMainDnsInput struct {
	DomainName string
}

type Dns interface {
	GetMainRef() string
	CreateMain(t *cloudformation.Template, i *CreateMainDnsInput)
}
