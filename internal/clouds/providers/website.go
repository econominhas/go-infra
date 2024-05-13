package providers

import "github.com/awslabs/goformation/v7/cloudformation"

type CreateStaticWebsiteInput struct {
	Name          string
	DnsRef        string
	FullDomain    string
	DeployUserArn string
}

type Website interface {
	CreateStatic(t *cloudformation.Template, i *CreateStaticWebsiteInput)
}
