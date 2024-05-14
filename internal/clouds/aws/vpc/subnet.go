package vpc

import (
	"strconv"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/econominhas/infra/internal/utils"
)

type createSubnetInput struct {
	Name       string
	SubnetType string
	VpcId      string
	Resources  cloudformation.Resources
}

const AMOUNT_OF_SUBNETS int = 2

func getIdx(subnetType string, idx int) int {
	if subnetType == privateEnum {
		return idx + 10
	}

	return idx
}

func (dps *Vpc) createSubnets(i *createSubnetInput) []string {
	ids := make([]string, AMOUNT_OF_SUBNETS)

	for idx := 0; idx < AMOUNT_OF_SUBNETS; idx++ {
		trueIdx := getIdx(i.SubnetType, idx)
		nbr := strconv.Itoa(trueIdx)

		azs := cloudformation.Select(trueIdx, []string{cloudformation.GetAZs("")})
		cidrBlock := "10.10." + nbr + ".0/24"
		subnetId := utils.GenId(&utils.GenIdInput{
			Id:   dps.StackId,
			Name: i.Name + nbr,
			Type: i.SubnetType + "sbn",
		})
		i.Resources[subnetId.Id] = &ec2.Subnet{
			AvailabilityZone: &azs,
			CidrBlock:        &cidrBlock,
			VpcId:            i.VpcId,
		}

		ids = append(ids, cloudformation.Ref(subnetId.Id))
	}

	return ids
}
