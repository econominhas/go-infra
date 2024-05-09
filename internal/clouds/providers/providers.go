package providers

import (
	"github.com/awslabs/goformation/v7/cloudformation"
)

// Vpc

type CreateMainVpcInput struct {
	Name string
}

type CreateMainVpcOutput struct {
	PublicSubnetsIds  []string
	PrivateSubnetsIds []string
}

type Vpc interface {
	CreateMain(t *cloudformation.Template, i *CreateMainVpcInput) *CreateMainVpcOutput
}

// Dns

type CreateMainDnsInput struct {
	DomainName string
}

type Dns interface {
	GetMainRef() string
	CreateMain(t *cloudformation.Template, i *CreateMainDnsInput)
}

// SqlDb

type CreateMainSqlDbInput struct {
	Name      string
	SubnetIds []string
}

type SqlDb interface {
	CreateMain(t *cloudformation.Template, i *CreateMainSqlDbInput)
}

// Website

type CreateStaticWebsiteInput struct {
	Name       string
	DnsRef     string
	FullDomain string
}

type Website interface {
	CreateStatic(t *cloudformation.Template, i *CreateStaticWebsiteInput)
}

// Api

type CreateBasicApiInput struct {
}

type Api interface {
	CreateBasic(t *cloudformation.Template, i *CreateBasicApiInput)
}
