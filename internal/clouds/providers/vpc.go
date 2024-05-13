package providers

import "github.com/awslabs/goformation/v7/cloudformation"

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
