/*
References:
- https://dev.to/femilawal/creating-an-ecs-cluster-with-cloudformation-54cg
- https://dev.to/aws-builders/mastering-aws-ecs-with-cloudformation-a-comprehensive-guide-aid
- https://towardsaws.com/deploying-mern-app-with-ecr-ecs-and-fargate-using-cloudformation-b54c03de707d
- https://sakyasumedh.medium.com/setup-application-load-balancer-and-point-to-ecs-deploy-to-aws-ecs-fargate-with-load-balancer-4b5f6785e8f
- https://sakyasumedh.medium.com/deploy-backend-application-to-aws-ecs-with-application-load-balancer-step-by-step-guide-part-3-b8125ca27177
- https://medium.com/@vladkens/aws-ecs-cluster-on-ec2-with-terraform-2023-fdb9f6b7db07
- https://medium.com/@sivajyothi.linga/ecs-cluster-with-ec2-launch-type-using-terraform-b5fbf535cb67
*/
package api

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/econominhas/infra/internal/clouds/providers"
)

type Api struct {
	StackId string
}

func (dps *Api) CreateBasic(t *cloudformation.Template, i *providers.CreateBasicApiInput) {}
