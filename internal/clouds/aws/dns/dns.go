// References
// https://xebia.com/blog/automated-provisioning-of-acm-certificates-using-route53-in-cloudformation/

package dns

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/certificatemanager"
	"github.com/awslabs/goformation/v7/cloudformation/route53"
	"github.com/econominhas/infra/internal/clouds/providers"
	"github.com/econominhas/infra/internal/utils"
)

type Dns struct {
	StackId string
}

func (dps *Dns) CreateMain(t *cloudformation.Template, i *providers.CreateMainDnsInput) {
	// Hosted Zone

	dnsId := utils.GenId(&utils.GenIdInput{
		Id:        dps.StackId,
		Name:      i.Name,
		Type:      "dns",
		OmitStage: true, // Dns should never have stage
	})
	t.Resources[dnsId] = &route53.HostedZone{
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
	t.Resources[certId] = &certificatemanager.Certificate{
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
