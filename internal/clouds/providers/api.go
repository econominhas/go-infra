package providers

import (
	"github.com/awslabs/goformation/v7/cloudformation"
)

type CreateBasicApiInput struct {
}

type Api interface {
	CreateBasic(t *cloudformation.Template, i *CreateBasicApiInput)
}
