package main

import (
	"github.com/jlgore/components"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
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
		publicRds, err := rds.NewInstance(ctx, "public-rds", &rds.InstanceArgs{
			InstanceClass:       pulumi.String("db.t3.micro"),
			AllocatedStorage:    pulumi.Int(20),
			Engine:              pulumi.String("mysql"),
			EngineVersion:       pulumi.String("8.0"),
			Username:            pulumi.String("admin"),
			Password:            pulumi.String("password123"), // Note: Use secrets manager in production
			DbName:              pulumi.String("publicdb"),
			PubliclyAccessible:  pulumi.Bool(true),
			DbSubnetGroupName:   createSubnetGroup(ctx, "public-subnet-group", vpc.PublicSubnetIds),
			VpcSecurityGroupIds: pulumi.StringArray{rdsSecurityGroup.ID()},
			SkipFinalSnapshot:   pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name":        pulumi.String("PublicRDSInstance"),
				"Environment": pulumi.String("Development"),
				"Project":     pulumi.String("cfi-rds"),
				"ManagedBy":   pulumi.String("Pulumi"),
			},
		})
		if err != nil {
			return err
		}

		// Create private RDS instance
		privateRds, err := rds.NewInstance(ctx, "private-rds", &rds.InstanceArgs{
			InstanceClass:       pulumi.String("db.t3.micro"),
			AllocatedStorage:    pulumi.Int(20),
			Engine:              pulumi.String("mysql"),
			EngineVersion:       pulumi.String("8.0"),
			Username:            pulumi.String("admin"),
			Password:            pulumi.String("password456"), // Note: Use secrets manager in production
			DbName:              pulumi.String("privatedb"),
			PubliclyAccessible:  pulumi.Bool(false),
			DbSubnetGroupName:   createSubnetGroup(ctx, "private-subnet-group", vpc.PrivateSubnetIds),
			VpcSecurityGroupIds: pulumi.StringArray{rdsSecurityGroup.ID()},
			SkipFinalSnapshot:   pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name":        pulumi.String("PrivateRDSInstance"),
				"Environment": pulumi.String("Development"),
				"Project":     pulumi.String("MyProject"),
				"ManagedBy":   pulumi.String("Pulumi"),
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

func createSubnetGroup(ctx *pulumi.Context, name string, subnetIds pulumi.IDArrayOutput) pulumi.StringOutput {
	subnetGroup, err := rds.NewSubnetGroup(ctx, name, &rds.SubnetGroupArgs{
		SubnetIds: subnetIds.ApplyT(func(ids []pulumi.ID) []string {
			result := make([]string, len(ids))
			for i, id := range ids {
				result[i] = string(id)
			}
			return result
		}).(pulumi.StringArrayOutput),
	})
	if err != nil {
		return pulumi.String("").ToStringOutput()
	}
	return subnetGroup.Name
}
