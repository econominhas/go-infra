// References
// https://xebia.com/blog/automated-provisioning-of-acm-certificates-using-route53-in-cloudformation/

package dns

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/certificatemanager"
	"github.com/awslabs/goformation/v7/cloudformation/route53"
	"github.com/econominhas/infra/internal/utils"
)

type Deps struct {
	StackId   string
	Resources cloudformation.Resources
}

type CreateMainDnsInput struct {
	Name       string
	DomainName string
}

func (dps *Deps) CreateMain(i *CreateMainDnsInput) {
	// Hosted Zone

	dnsId := utils.GenId(&utils.GenIdInput{
		Id:        dps.StackId,
		Name:      i.Name,
		Type:      "dns",
		OmitStage: true, // Dns should never have stage
	})
	dps.Resources[dnsId] = &route53.HostedZone{
		Name: &i.DomainName,
	}

	// Certificate

	valMethod := "DNS"
	dnsRef := cloudformation.Ref(dnsId)

	certId := utils.GenId(&utils.GenIdInput{
		Id:        dps.StackId,
		Name:      i.Name,
		Type:      "cert",
		OmitStage: true, // Dns should never have stage
	})
	dps.Resources[certId] = &certificatemanager.Certificate{
		DomainName:              i.DomainName,
		ValidationMethod:        &valMethod,
		SubjectAlternativeNames: []string{"*." + i.DomainName},
		DomainValidationOptions: []certificatemanager.Certificate_DomainValidationOption{
			{
				DomainName:   i.DomainName,
				HostedZoneId: &dnsRef,
			},
		},
	}
}
