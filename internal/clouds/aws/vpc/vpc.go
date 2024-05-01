// References
// https://www.infoq.com/articles/aws-vpc-cloudformation/
// https://blog.devops.dev/how-to-use-aws-cloudformation-to-create-a-vpc-10dbd70a3677

package vpc

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/econominhas/infra/internal/utils"
)

type Deps struct {
	StackId   string
	Resources cloudformation.Resources
}

type CreateMainVpcInput struct {
	Name string
}

const (
	publicEnum  = "public"
	privateEnum = "private"
)

func (dps *Deps) CreateMain(i *CreateMainVpcInput) {
	// Vpc

	cidrBlock := "10.10.0.0/16"
	enableDns := true
	vpcId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "vpc",
	})
	dps.Resources[vpcId] = &ec2.VPC{
		CidrBlock:          &cidrBlock,
		EnableDnsSupport:   &enableDns,
		EnableDnsHostnames: &enableDns,
	}
	vpcRef := cloudformation.Ref(vpcId)

	// Public Subnets

	createSubnets(CreateSubnetInput{
		StackId:    dps.StackId,
		Resources:  dps.Resources,
		Name:       i.Name,
		SubnetType: publicEnum,
		VpcId:      vpcId,
	})

	// Private Subnets

	createSubnets(CreateSubnetInput{
		StackId:    dps.StackId,
		Resources:  dps.Resources,
		Name:       i.Name,
		SubnetType: privateEnum,
		VpcId:      vpcId,
	})

	// Internet Gateway

	igId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "ig",
	})
	dps.Resources[igId] = &ec2.InternetGateway{}
	igRef := cloudformation.Ref(igId)

	igaId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "iga",
	})
	dps.Resources[igaId] = &ec2.VPCGatewayAttachment{
		InternetGatewayId: &igRef,
		VpcId:             vpcRef,
	}

	// Public Route Table

	createPublicRouteTable(CreatePublicRouteTableInput{
		StackId:   dps.StackId,
		Resources: dps.Resources,
		Name:      i.Name,
		VpcId:     vpcId,
	})
}
