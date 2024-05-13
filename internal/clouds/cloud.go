package clouds

import "github.com/econominhas/infra/internal/clouds/providers"

// Cloud

type Cloud interface {
	Vpc() providers.Vpc
	Dns() providers.Dns
	SqlDb() providers.SqlDb
	Iam() providers.Iam

	Website() providers.Website
	Api() providers.Api
}
