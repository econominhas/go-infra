package iam

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/iam"
	"github.com/awslabs/goformation/v7/cloudformation/secretsmanager"
	"github.com/econominhas/infra/internal/clouds/providers"
	"github.com/econominhas/infra/internal/utils"
)

type Iam struct {
	StackId string
}

func (dps *Iam) CreateDeployUser(t *cloudformation.Template, i *providers.CreateDeployUserInput) *providers.CreateDeployUserOutput {
	userId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "user",
	})
	t.Resources[userId.Id] = &iam.User{
		UserName: &userId.Name,
	}

	keyId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "key",
	})
	t.Resources[keyId.Id] = &iam.AccessKey{
		UserName: userId.Name,
		AWSCloudFormationDependsOn: []string{
			userId.Id,
		},
	}

	keySecretId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "keysecret",
	})
	secretName := keySecretId.Name
	secretValue := cloudformation.GetAtt(keyId.Id, "SecretAccessKey")
	t.Resources[keySecretId.Id] = &secretsmanager.Secret{
		Name:         &secretName,
		SecretString: &secretValue,
		AWSCloudFormationDependsOn: []string{
			userId.Id,
			keyId.Id,
		},
	}

	return &providers.CreateDeployUserOutput{
		Arn: cloudformation.GetAtt(userId.Id, "Arn"),
	}
}
