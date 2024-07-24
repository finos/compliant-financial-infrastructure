package main

import (
	"github.com/jlgore/components"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create VPC
		vpc, err := components.NewVPC(ctx, "my-vpc", &components.VPCArgs{
			CidrBlock:          "10.0.0.0/16",
			EnableDnsHostnames: true,
			EnableDnsSupport:   true,
			InstanceType:       "t3.micro", // For NAT instance
			Region:             "us-east-1",
			Tags: components.TagArgs{
				Default: map[string]string{
					"Project":     "cfi-rds",
					"Environment": "Sandbox",
				},
				VPC: map[string]string{
					"Type": "Main",
				},
				NatInstance: map[string]string{
					"Role": "NAT",
				},
				NatSecurityGroup: map[string]string{
					"Purpose": "Allow NAT traffic",
				},
				PublicSubnet: map[string]string{
					"Accessibility": "Public",
				},
				PrivateSubnet: map[string]string{
					"Accessibility": "Private",
				},
			},
		})
		if err != nil {
			return err
		}

		// Create a security group for RDS
		rdsSecurityGroup, err := ec2.NewSecurityGroup(ctx, "rds-sg", &ec2.SecurityGroupArgs{
			VpcId: vpc.VpcId,
			Ingress: ec2.SecurityGroupIngressArray{
				ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(3306),
					ToPort:     pulumi.Int(3306),
					CidrBlocks: pulumi.StringArray{pulumi.String("10.0.0.0/16")},
				},
			},
		})
		if err != nil {
			return err
		}

		// Create public RDS instance
		publicRds, err := components.NewRDSInstance(ctx, "cfi-public-rds", &components.RDSInstanceArgs{
			InstanceClass:      "db.t3.micro",
			AllocatedStorage:   20,
			Engine:             "mysql",
			EngineVersion:      "8.0",
			Username:           "admin",
			Password:           "password123", // Use secrets manager in production
			DbName:             "publicdb",
			PubliclyAccessible: true,
			SubnetIds: vpc.PublicSubnetIds.ApplyT(func(ids []pulumi.ID) []string {
				result := make([]string, len(ids))
				for i, id := range ids {
					result[i] = string(id)
				}
				return result
			}).(pulumi.StringArrayOutput),
			VpcSecurityGroupIds: pulumi.StringArray{rdsSecurityGroup.ID()},
			CreateKey:           false,
			Encryption:          false,
			Tags: map[string]string{
				"Name":        "PublicRDSInstance",
				"Environment": "Development",
				"Project":     "cfi-rds",
				"ManagedBy":   "Pulumi",
			},
		})
		if err != nil {
			return err
		}

		// Create private RDS instance
		privateRds, err := components.NewRDSInstance(ctx, "cfi-private-rds", &components.RDSInstanceArgs{
			InstanceClass:      "db.t3.micro",
			AllocatedStorage:   20,
			Engine:             "mysql",
			EngineVersion:      "8.0",
			Username:           "admin",
			Password:           "password456", // Use secrets manager in production
			DbName:             "privatedb",
			PubliclyAccessible: false,
			SubnetIds: vpc.PrivateSubnetIds.ApplyT(func(ids []pulumi.ID) []string {
				result := make([]string, len(ids))
				for i, id := range ids {
					result[i] = string(id)
				}
				return result
			}).(pulumi.StringArrayOutput),
			VpcSecurityGroupIds: pulumi.StringArray{rdsSecurityGroup.ID()},
			CreateKey:           true,
			Encryption:          true,
			Tags: map[string]string{
				"Name":        "PrivateRDSInstance",
				"Environment": "Development",
				"Project":     "cfi-rds",
				"ManagedBy":   "Pulumi",
			},
		})
		if err != nil {
			return err
		}
		// Export the VPC and RDS information
		ctx.Export("vpcId", vpc.VpcId)
		ctx.Export("publicSubnetIds", vpc.PublicSubnetIds)
		ctx.Export("privateSubnetIds", vpc.PrivateSubnetIds)
		ctx.Export("publicRdsEndpoint", publicRds.Endpoint)
		ctx.Export("privateRdsEndpoint", privateRds.Endpoint)

		return nil
	})
}
