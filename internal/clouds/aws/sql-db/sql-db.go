/*
References
- https://dev.to/aws/provisioning-an-rds-database-with-cloudformation-part-2-i6n
*/
package sqldb

import (
	"github.com/awslabs/goformation/v7/cloudformation"
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
		DBSecurityGroups: []string{
			cloudformation.Ref(sbngId.Id),
		},
	}
}
