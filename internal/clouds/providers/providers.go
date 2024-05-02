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
	DomainName string
}

type Dns interface {
	GetMainRef() string
	CreateMain(t *cloudformation.Template, i *CreateMainDnsInput)
}

// Website

type CreateStaticWebsiteInput struct {
	Name          string
	DnsRef        string
	SubDomainName string
}

type Website interface {
	CreateStatic(t *cloudformation.Template, i *CreateStaticWebsiteInput)
}
