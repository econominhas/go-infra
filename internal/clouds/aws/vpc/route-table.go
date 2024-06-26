package vpc

import (
	"strconv"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/econominhas/infra/internal/utils"
)

type createPublicRouteTableInput struct {
	Name      string
	VpcId     string
	IgRef     string
	Resources cloudformation.Resources
}

func (dps *Vpc) createPublicRouteTable(i *createPublicRouteTableInput) {
	// Route Table

	rtId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: publicEnum + "rt",
	})
	i.Resources[rtId.Id] = &ec2.RouteTable{
		VpcId: i.VpcId,
	}
	rtRef := cloudformation.Ref(rtId.Id)

	destCidrBlock := "0.0.0.0/0"
	routeId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: publicEnum + "rt",
	})
	i.Resources[routeId.Id] = &ec2.Route{
		DestinationCidrBlock: &destCidrBlock,
		GatewayId:            &i.IgRef,
		RouteTableId:         rtRef,
	}
	routeRef := cloudformation.Ref(routeId.Id)

	// Route Table Association

	for idx := 0; idx <= 1; idx++ {
		subnetId := utils.GenId(&utils.GenIdInput{
			Id:   dps.StackId,
			Name: i.Name + strconv.Itoa(idx),
			Type: publicEnum + "sbn",
		})
		subnetRef := cloudformation.Ref(subnetId.Id)

		rtSb1Id := utils.GenId(&utils.GenIdInput{
			Id:   dps.StackId,
			Name: i.Name + "0",
			Type: "sbnrta",
		})
		i.Resources[rtSb1Id.Id] = &ec2.SubnetRouteTableAssociation{
			RouteTableId: routeRef,
			SubnetId:     subnetRef,
		}
	}
}
