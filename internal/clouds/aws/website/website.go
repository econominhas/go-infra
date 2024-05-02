// References
// https://awstip.com/deploy-a-static-website-to-aws-s3-in-seconds-with-cloudformation-ac489158054f
// https://www.golinuxcloud.com/how-to-host-static-website-on-s3/

package website

import (
	"github.com/awslabs/goformation/v7/cloudformation"
	"github.com/awslabs/goformation/v7/cloudformation/s3"
	"github.com/econominhas/infra/internal/clouds/providers"
	"github.com/econominhas/infra/internal/utils"
)

type Website struct {
	StackId string
}

func (dps *Website) CreateStatic(t *cloudformation.Template, i *providers.CreateStaticWebsiteInput) {
	// Bucket

	accessControl := "PublicRead"
	indexDocument := "index.html"
	errorDocument := "error.html"
	blockPublicAccess := false
	maxCorsAge := 3000

	bucketId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "s3",
	})
	t.Resources[bucketId] = &s3.Bucket{
		BucketName:    &bucketId,
		AccessControl: &accessControl,
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
		CorsConfiguration: &s3.Bucket_CorsConfiguration{
			CorsRules: []s3.Bucket_CorsRule{
				{
					AllowedHeaders: []string{"authorization"},
					AllowedMethods: []string{"GET"},
					AllowedOrigins: []string{
						cloudformation.GetAtt(bucketId, "WebsiteURL"),
					},
					ExposedHeaders: []string{},
					MaxAge:         &maxCorsAge,
				},
			},
		},
	}
	bucketRef := cloudformation.Ref(bucketId)

	// Bucket Policy

	bucketPId := utils.GenId(&utils.GenIdInput{
		Id:   dps.StackId,
		Name: i.Name,
		Type: "s3p",
	})
	t.Resources[bucketPId] = &s3.BucketPolicy{
		Bucket:         bucketRef,
		PolicyDocument: "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Sid\":\"PublicReadGetObject\",\"Effect\":\"Allow\",\"Principal\":\"*\",\"Action\":[\"s3:GetObject\"],\"Resource\":[\"arn:aws:s3:::" + bucketId + "/*\"]}]}",
	}

	// Subdomain

	// ttl := "900"

	// rsgId := utils.GenId(&utils.GenIdInput{
	// 	Id:   dps.StackId,
	// 	Name: i.Name,
	// 	Type: "rsg",
	// })
	// t.Resources[rsgId] = &route53.RecordSetGroup{
	// 	HostedZoneId: &i.DnsRef,
	// 	RecordSets: []route53.RecordSetGroup_RecordSet{
	// 		{
	// 			Name: rsgId,
	// 			Type: "CNAE",
	// 			TTL:  &ttl,
	// 			ResourceRecords: []string{
	// 				cloudformation.Join(
	// 					"",
	// 					cloudformation.Split(
	// 						"http://",
	// 						cloudformation.GetAtt(bucketId, "WebsiteURL"),
	// 					),
	// 				),
	// 			},
	// 		},
	// 	},
	// }
}
