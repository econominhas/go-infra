package vpc

import (
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/econominhas/infra/internal/utils"
)

type createEc2SecurityGroupInput struct {
	Name     string
	VpcId    string
	Template *cloudformation.Template
}

type Ip struct {
	Ec2Ip string
}

type AwsConsoleIp struct {
	UsEast1 *Ip
	UsEast2 *Ip
	UsWest1 *Ip
	UsWest2 *Ip
}

func (dps *Vpc) createEc2SecurityGroup(i *createEc2SecurityGroupInput) string {
	// Official Ips by AWS
	// https://ip-ranges.amazonaws.com/ip-ranges.json
	i.Template.Mappings["AwsConsoleIp"] = &AwsConsoleIp{
		UsEast1: &Ip{
			Ec2Ip: "18.206.107.24/29",
		},
		UsEast2: &Ip{
			Ec2Ip: "3.16.146.0/29",
		},
		UsWest1: &Ip{
			Ec2Ip: "13.52.6.112/29",
		},
		UsWest2: &Ip{
			Ec2Ip: "18.237.140.160/29",
		},
	}

	// Creates a security group that allows
	// outbound traffic and inbound traffic
	// from port 80 (http) and 443 (https)
	// and ssh access from AWS Console

	httpPort := 80
	httpsPort := 443
	sshPort := 443
	source := "0.0.0.0/0"
	awsConsoleIp := cloudformation.FindInMap(
		"AwsConsoleIp",
		utils.ToPascal(os.Getenv("AWS_REGION")),
		"Ec2Ip",
	)
	sgId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "sg",
	})
	i.Template.Resources[sgId.Id] = &ec2.SecurityGroup{
		GroupName: &sgId.Name,
		VpcId:     &i.VpcId,
		SecurityGroupIngress: []ec2.SecurityGroup_Ingress{
			{
				IpProtocol: "tcp",
				FromPort:   &httpPort,
				ToPort:     &httpPort,
				CidrIp:     &source,
			},
			{
				IpProtocol: "tcp",
				FromPort:   &httpsPort,
				ToPort:     &httpsPort,
				CidrIp:     &source,
			},
			{
				IpProtocol: "tcp",
				FromPort:   &sshPort,
				ToPort:     &sshPort,
				CidrIp:     &awsConsoleIp,
			},
		},
	}

	return cloudformation.Ref(sgId.Id)
}
