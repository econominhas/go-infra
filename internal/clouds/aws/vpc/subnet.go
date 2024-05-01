package vpc

import (
	"strconv"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/econominhas/infra/internal/utils"
)

type CreateSubnetInput struct {
	StackId    string
	Name       string
	SubnetType string
	VpcId      string
	Resources  cloudformation.Resources
}

func getIdx(subnetType string, idx int) int {
	if subnetType == privateEnum {
		return idx + 10
	}

	return idx
}

func createSubnets(i CreateSubnetInput) {
	for idx := 0; idx <= 1; idx++ {
		trueIdx := getIdx(i.SubnetType, idx)
		nbr := strconv.Itoa(trueIdx)

		azs := cloudformation.Select(trueIdx, []string{cloudformation.GetAZs("")})
		cidrBlock := "10.10." + nbr + ".0/24"
		subnetId := utils.GenId(&utils.GenIdInput{
			Id:   i.StackId,
			Name: i.Name + nbr,
			Type: i.SubnetType + "sbn",
		})
		i.Resources[subnetId] = &ec2.Subnet{
			AvailabilityZone: &azs,
			CidrBlock:        &cidrBlock,
			VpcId:            i.VpcId,
		}
	}
}
