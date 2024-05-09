/*
References
- https://xebia.com/blog/automated-provisioning-of-acm-certificates-using-route53-in-cloudformation/
*/
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

func (dps *Dns) getMainId() *utils.GenIdOutput {
	return utils.GenId(&utils.GenIdInput{
		Id:        dps.StackId,
		Name:      dps.StackId,
		Type:      "dns",
		OmitStage: true, // Dns should never have stage
	})
}

func (dps *Dns) GetMainRef() string {
	dnsId := dps.getMainId()
	return cloudformation.ImportValue(dnsId.Name)
}

func (dps *Dns) CreateMain(t *cloudformation.Template, i *providers.CreateMainDnsInput) {
	// Hosted Zone

	dnsId := dps.getMainId()
	t.Resources[dnsId.Id] = &route53.HostedZone{
		Name: &i.DomainName,
	}

	t.Outputs[dnsId.Id+"Output"] = cloudformation.Output{
		Value: cloudformation.Ref(dnsId.Id),
		Export: &cloudformation.Export{
			Name: dnsId.Name,
		},
	}

	// Certificate

	valMethod := "DNS"
	dnsRef := cloudformation.Ref(dnsId.Id)

	certId := utils.GenId(&utils.GenIdInput{
		Id:        dps.StackId,
		Name:      dps.StackId,
		Type:      "cert",
		OmitStage: true, // Dns should never have stage
	})
	t.Resources[certId.Id] = &certificatemanager.Certificate{
		DomainName:              i.DomainName,
		ValidationMethod:        &valMethod,
		SubjectAlternativeNames: []string{"*." + i.DomainName},
		DomainValidationOptions: []certificatemanager.Certificate_DomainValidationOption{
			{
				DomainName:   i.DomainName,
				HostedZoneId: &dnsRef,
			},
		},
		AWSCloudFormationDependsOn: []string{
			dnsId.Id,
		},
	}
}
