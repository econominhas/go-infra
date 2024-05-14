/*
References
- https://www.infoq.com/articles/aws-vpc-cloudformation/
- https://blog.devops.dev/how-to-use-aws-cloudformation-to-create-a-vpc-10dbd70a3677
*/
package vpc

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/econominhas/infra/internal/clouds/providers"
	"github.com/econominhas/infra/internal/utils"
)

type Vpc struct {
	StackId string
}

const (
	publicEnum  = "public"
	privateEnum = "private"
)

func (dps *Vpc) CreateMain(t *cloudformation.Template, i *providers.CreateMainVpcInput) *providers.CreateMainVpcOutput {
	// Vpc

	cidrBlock := "10.10.0.0/16"
	enableDns := true
	vpcId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "vpc",
	})
	t.Resources[vpcId.Id] = &ec2.VPC{
		CidrBlock:          &cidrBlock,
		EnableDnsSupport:   &enableDns,
		EnableDnsHostnames: &enableDns,
	}
	vpcRef := cloudformation.Ref(vpcId.Id)

	// Public Subnets

	publicSubnetIds := dps.createSubnets(&createSubnetInput{
		Resources:  t.Resources,
		Name:       i.Name,
		SubnetType: publicEnum,
		VpcId:      vpcId.Id,
	})

	// Private Subnets

	privateSubnetIds := dps.createSubnets(&createSubnetInput{
		Resources:  t.Resources,
		Name:       i.Name,
		SubnetType: privateEnum,
		VpcId:      vpcId.Id,
	})

	// Internet Gateway

	igId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "ig",
	})
	t.Resources[igId.Id] = &ec2.InternetGateway{}
	igRef := cloudformation.Ref(igId.Id)

	igaId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "iga",
	})
	t.Resources[igaId.Id] = &ec2.VPCGatewayAttachment{
		InternetGatewayId: &igRef,
		VpcId:             vpcRef,
	}

	// Public Route Table

	dps.createPublicRouteTable(&createPublicRouteTableInput{
		Resources: t.Resources,
		Name:      i.Name,
		VpcId:     vpcId.Id,
	})

	// Ec2 Security Group

	ec2SgRef := dps.createEc2SecurityGroup(&createEc2SecurityGroupInput{
		Name:     i.Name,
		VpcId:    vpcId.Id,
		Template: t,
	})

	// Return

	return &providers.CreateMainVpcOutput{
		PublicSubnetsIds:  publicSubnetIds,
		PrivateSubnetsIds: privateSubnetIds,
		Ec2SgRef:          ec2SgRef,
	}
}
