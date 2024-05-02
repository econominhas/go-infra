package clouds

import "github.com/econominhas/infra/internal/clouds/providers"

// Cloud

type Cloud interface {
	Vpc() providers.Vpc
	Dns() providers.Dns
}
