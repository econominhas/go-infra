/*
References
- https://dev.to/tiamatt/aws-project-module-1-host-a-static-website-on-aws-s3-via-cloudformation-2pa2
- https://dev.to/tiamatt/aws-project-module-3-use-your-custom-domain-for-static-website-on-aws-s3-via-route-53-and-cloudformation-34cn
- https://stackoverflow.com/questions/40865710/declaring-an-iam-access-key-resource-by-cloudformation#40866799
*/
package website

import (
	"os"

	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/route53"
	"github.com/awslabs/goformation/v7/cloudformation/s3"
	"github.com/econominhas/infra/internal/clouds/providers"
	"github.com/econominhas/infra/internal/utils"
)

type Website struct {
	StackId string
}

type Region struct {
	S3HostedZoneId string
}

type RegionMap struct {
	UsEast1 *Region
	UsWest1 *Region
	UsWest2 *Region
}

func (dps *Website) CreateStatic(t *cloudformation.Template, i *providers.CreateStaticWebsiteInput) {
	// Bucket

	indexDocument := "index.html"
	errorDocument := "error.html"
	blockPublicAccess := false
	maxCorsAge := 3000
	objectOwnership := "ObjectWriter"

	bucketId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "s3",
	})
	t.Resources[bucketId.Id] = &s3.Bucket{
		BucketName: &i.FullDomain,
		WebsiteConfiguration: &s3.Bucket_WebsiteConfiguration{
			IndexDocument: &indexDocument,
			ErrorDocument: &errorDocument,
		},
		PublicAccessBlockConfiguration: &s3.Bucket_PublicAccessBlockConfiguration{
			BlockPublicAcls:       &blockPublicAccess,
			BlockPublicPolicy:     &blockPublicAccess,
			IgnorePublicAcls:      &blockPublicAccess,
			RestrictPublicBuckets: &blockPublicAccess,
		},
		OwnershipControls: &s3.Bucket_OwnershipControls{
			Rules: []s3.Bucket_OwnershipControlsRule{
				{
					ObjectOwnership: &objectOwnership,
				},
			},
		},
		CorsConfiguration: &s3.Bucket_CorsConfiguration{
			CorsRules: []s3.Bucket_CorsRule{
				{
					AllowedHeaders: []string{"authorization"},
					AllowedMethods: []string{"GET"},
					AllowedOrigins: []string{"*"},
					ExposedHeaders: []string{},
					MaxAge:         &maxCorsAge,
				},
			},
		},
	}
	bucketRef := cloudformation.Ref(bucketId.Id)

	// Bucket Policy

	bucketPId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "s3p",
	})
	t.Resources[bucketPId.Id] = &s3.BucketPolicy{
		Bucket: bucketRef,
		PolicyDocument: &providers.PolicyDocument{
			Version: "2012-10-17",
			Statement: []providers.Statement{
				{
					Sid:       "PublicReadGetObject",
					Effect:    "Allow",
					Principal: "*",
					Action:    []string{"s3:GetObject"},
					Resource: cloudformation.Join("", []string{
						"arn:aws:s3:::",
						bucketRef,
						"/*",
					}),
				},
				{
					Sid:    "AllowDeployUser",
					Effect: "Allow",
					Principal: providers.AwsPrincipal{
						AWS: i.DeployUserArn,
					},
					Action: []string{"s3:ListBucket"},
					Resource: cloudformation.Join("", []string{
						"arn:aws:s3:::",
						bucketRef,
					}),
				},
				{
					Sid:    "AllowDeployUserPutObject",
					Effect: "Allow",
					Principal: providers.AwsPrincipal{
						AWS: i.DeployUserArn,
					},
					Action: []string{
						"s3:PutObject",
						"s3:PutObjectAcl",
						"s3:GetObject",
						"s3:GetObjectAcl",
						"s3:DeleteObject",
					},
					Resource: cloudformation.Join("", []string{
						"arn:aws:s3:::",
						bucketRef,
						"/*",
					}),
				},
			},
		},
		AWSCloudFormationDependsOn: []string{
			bucketId.Id,
		},
	}

	// Subdomain

	t.Mappings["RegionMap"] = &RegionMap{
		UsEast1: &Region{
			S3HostedZoneId: "Z3AQBSTGFYJSTF",
		},
		UsWest1: &Region{
			S3HostedZoneId: "Z2F56UZL2M1ACD",
		},
		UsWest2: &Region{
			S3HostedZoneId: "Z3BJ6K6RIION7M",
		},
	}

	rsgId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "rsg",
	})
	t.Resources[rsgId.Id] = &route53.RecordSetGroup{
		HostedZoneId: &i.DnsRef,
		RecordSets: []route53.RecordSetGroup_RecordSet{
			{
				Name: i.FullDomain,
				Type: "A",
				AliasTarget: &route53.RecordSetGroup_AliasTarget{
					DNSName: cloudformation.Sub("s3-website-${AWS::Region}.amazonaws.com"),
					HostedZoneId: cloudformation.FindInMap(
						"RegionMap",
						utils.ToPascal(os.Getenv("AWS_REGION")),
						"S3HostedZoneId",
					),
				},
			},
		},
		AWSCloudFormationDependsOn: []string{
			bucketId.Id,
			bucketPId.Id,
		},
	}
}
