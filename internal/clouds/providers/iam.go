package providers

import "github.com/awslabs/goformation/v7/cloudformation"

type AwsPrincipal struct {
	AWS string
}

type Statement struct {
	Sid       string
	Effect    string
	Principal interface{}
	Action    []string
	Resource  string
}

type PolicyDocument struct {
	Version   string
	Statement []Statement
}

type CreateDeployUserInput struct {
	Name string
}

type CreateDeployUserOutput struct {
	Arn string
}

type Iam interface {
	CreateDeployUser(t *cloudformation.Template, i *CreateDeployUserInput) *CreateDeployUserOutput
}
