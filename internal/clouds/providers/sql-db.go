package providers

import "github.com/awslabs/goformation/v7/cloudformation"

type CreateMainSqlDbInput struct {
	Name      string
	SubnetIds []string
}

type SqlDb interface {
	CreateMain(t *cloudformation.Template, i *CreateMainSqlDbInput)
}
