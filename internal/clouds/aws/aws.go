package aws

import (
	"github.com/econominhas/infra/internal/clouds"
	"github.com/econominhas/infra/internal/clouds/aws/dns"
	"github.com/econominhas/infra/internal/clouds/aws/vpc"
	"github.com/econominhas/infra/internal/clouds/aws/website"
	"github.com/econominhas/infra/internal/clouds/providers"
)

type Aws struct {
	StackId string
}

func (aws *Aws) Vpc() providers.Vpc {
	return &vpc.Vpc{
		StackId: aws.StackId,
	}
}

func (aws *Aws) Dns() providers.Dns {
	return &dns.Dns{
		StackId: aws.StackId,
	}
}

func (aws *Aws) Website() providers.Website {
	return &website.Website{
		StackId: aws.StackId,
	}
}

func NewAws(stackId string) clouds.Cloud {
	return &Aws{
		StackId: stackId,
	}
}
