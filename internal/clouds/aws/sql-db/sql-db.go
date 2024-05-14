/*
References
- https://dev.to/aws/provisioning-an-rds-database-with-cloudformation-part-2-i6n
*/
package sqldb

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/ec2"
	"github.com/awslabs/goformation/v7/cloudformation/rds"
	"github.com/econominhas/infra/internal/clouds/providers"
	"github.com/econominhas/infra/internal/utils"
)

type SqlDb struct {
	StackId string
}

func (dps *SqlDb) CreateMain(t *cloudformation.Template, i *providers.CreateMainSqlDbInput) {
	dbId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "sqldb",
	})

	// Subnet Group

	sbngId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "sbng",
	})
	t.Resources[sbngId.Id] = &rds.DBSubnetGroup{
		DBSubnetGroupName:        &sbngId.Name,
		DBSubnetGroupDescription: dbId.Name + " subnet group",
		SubnetIds:                i.SubnetIds,
	}

	// Security Group

	// Creates a security group that only allows
	// inbound traffic from the SecurityGroupRef
	// and doesn't allow any outbound traffic

	dbPort := 5432
	sgId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "sg",
	})
	t.Resources[sgId.Id] = &ec2.SecurityGroup{
		GroupName: &sgId.Name,
		SecurityGroupIngress: []ec2.SecurityGroup_Ingress{
			{
				IpProtocol:              "tcp",
				FromPort:                &dbPort,
				ToPort:                  &dbPort,
				SourceSecurityGroupName: &i.Ec2SgRef,
			},
		},
	}

	// Database

	engine := "postgres"
	instanceClass := "db.t2.micro"
	multiAz := false

	allocatedStorage := "20"
	maxAllocatedStorage := 120

	allowMajorVersionUpgrade := false
	autoMinorVersionUpgrade := true

	backupRetention := 0
	deleteAutomatedBackups := true

	publiclyAccessible := false

	storageEncrypted := true
	storageType := "gp2"

	deletionProtection := false

	t.Resources[dbId.Id] = &rds.DBInstance{
		DBName:                   &dps.StackId,
		DBInstanceIdentifier:     &dbId.Name,
		Engine:                   &engine,
		DBInstanceClass:          &instanceClass,
		MultiAZ:                  &multiAz,
		AllocatedStorage:         &allocatedStorage,
		MaxAllocatedStorage:      &maxAllocatedStorage,
		AllowMajorVersionUpgrade: &allowMajorVersionUpgrade,
		AutoMinorVersionUpgrade:  &autoMinorVersionUpgrade,
		BackupRetentionPeriod:    &backupRetention,
		DeleteAutomatedBackups:   &deleteAutomatedBackups,
		DeletionProtection:       &deletionProtection,
		PubliclyAccessible:       &publiclyAccessible,
		StorageEncrypted:         &storageEncrypted,
		StorageType:              &storageType,
		DBSubnetGroupName:        &sbngId.Name,
		DBSecurityGroups: []string{
			cloudformation.Ref(sgId.Id),
		},
	}
}
