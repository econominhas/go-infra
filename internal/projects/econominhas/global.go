package econominhas

// ---------------------------------
//
//  DO NOT USE THE GLOBAL STACK AS
//  REFERENCE TO CREATE OTHER
//  NON-GLOBAL STACKS!!!!
//
// ---------------------------------

import (
	"github.com/awslabs/goformation/v7/cloudformation"

	"github.com/econominhas/infra/internal/clouds/aws"
	"github.com/econominhas/infra/internal/clouds/providers"
)

// The project ID should always be declared
// at the global stack so it can be used
// by the other stacks
const PROJECT_ID = "econominhas"

// The global stack is responsible for creating
// every resource that is not used by a product
// in specific, but can be used by all/multiple
// products
func Global() ([]byte, error) {
	stackId := PROJECT_ID

	template := cloudformation.NewTemplate()

	cloud := aws.NewAws(stackId)

	// Dns

	dns := cloud.Dns()

	dns.CreateMain(template, &providers.CreateMainDnsInput{
		DomainName: "econominhas.com.br",
	})

	// Vpc

	// vpc := cloud.Vpc()

	// vpcOutput := vpc.CreateMain(template, &providers.CreateMainVpcInput{
	// 	Name: stackId,
	// })

	// Sql

	// sqlDb := cloud.SqlDb()

	// sqlDb.CreateMain(template, &providers.CreateMainSqlDbInput{
	// 	Name:      stackId,
	// 	SubnetIds: vpcOutput.PrivateSubnetsIds,
	// })

	return template.YAML()
}
