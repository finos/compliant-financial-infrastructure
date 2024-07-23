package components

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VPCArgs struct {
	CidrBlock          string
	EnableDnsHostnames bool
	EnableDnsSupport   bool
	InstanceType       string
	Region             string
	Tags               TagArgs
}

type TagArgs struct {
	VPC               map[string]string
	InternetGateway   map[string]string
	PublicSubnet      map[string]string
	PrivateSubnet     map[string]string
	NatInstance       map[string]string
	NatSecurityGroup  map[string]string
	PublicRouteTable  map[string]string
	PrivateRouteTable map[string]string
	Default           map[string]string
}

type VPC struct {
	pulumi.ResourceState

	VpcId            pulumi.IDOutput
	PublicSubnetIds  pulumi.IDArrayOutput
	PrivateSubnetIds pulumi.IDArrayOutput
	NatInstanceId    pulumi.IDOutput
	NatInstanceAmiId pulumi.StringOutput
}

func lookupNatAmi(ctx *pulumi.Context, provider *aws.Provider) (pulumi.StringOutput, error) {
	ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
		Owners: []string{"568608671756"},
		Filters: []ec2.GetAmiFilter{
			{
				Name:   "name",
				Values: []string{"fck-nat-al2023-*"},
			},
		},
		MostRecent: pulumi.BoolRef(true),
	}, pulumi.Provider(provider))
	if err != nil {
		return pulumi.StringOutput{}, err
	}
	return pulumi.String(ami.Id).ToStringOutput(), nil
}

func NewVPC(ctx *pulumi.Context, name string, args *VPCArgs, opts ...pulumi.ResourceOption) (*VPC, error) {
	vpc := &VPC{}
	err := ctx.RegisterComponentResource("custom:network:VPC", name, vpc, opts...)
	if err != nil {
		return nil, err
	}

	// Create a new AWS provider with the specified region
	awsProvider, err := aws.NewProvider(ctx, "aws-provider", &aws.ProviderArgs{
		Region: pulumi.String(args.Region),
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	// Look up the NAT instance AMI
	amiId, err := lookupNatAmi(ctx, awsProvider)
	if err != nil {
		return nil, err
	}

	// Helper function to merge tags
	mergeTags := func(specific map[string]string, defaultTags map[string]string, resourceName string) pulumi.StringMap {
		merged := pulumi.StringMap{}
		for k, v := range defaultTags {
			merged[k] = pulumi.String(v)
		}
		for k, v := range specific {
			merged[k] = pulumi.String(v)
		}
		merged["Name"] = pulumi.String(fmt.Sprintf("%s-%s", name, resourceName))
		return merged
	}

	// Create VPC
	vpcTags := mergeTags(args.Tags.VPC, args.Tags.Default, "vpc")
	vpcResource, err := ec2.NewVpc(ctx, fmt.Sprintf("%s-vpc", name), &ec2.VpcArgs{
		CidrBlock:          pulumi.String(args.CidrBlock),
		EnableDnsHostnames: pulumi.Bool(args.EnableDnsHostnames),
		EnableDnsSupport:   pulumi.Bool(args.EnableDnsSupport),
		Tags:               vpcTags,
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	// Create Internet Gateway
	igwTags := mergeTags(args.Tags.InternetGateway, args.Tags.Default, "igw")
	igw, err := ec2.NewInternetGateway(ctx, fmt.Sprintf("%s-igw", name), &ec2.InternetGatewayArgs{
		VpcId: vpcResource.ID(),
		Tags:  igwTags,
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	// Create public subnets
	publicSubnets := make([]pulumi.IDOutput, 2)
	for i := 0; i < 2; i++ {
		subnetTags := mergeTags(args.Tags.PublicSubnet, args.Tags.Default, fmt.Sprintf("public-subnet-%d", i))
		subnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("%s-public-subnet-%d", name, i), &ec2.SubnetArgs{
			VpcId:               vpcResource.ID(),
			CidrBlock:           pulumi.Sprintf("10.0.%d.0/24", i),
			AvailabilityZone:    pulumi.Sprintf("us-east-1%s", string(rune('a'+i))),
			MapPublicIpOnLaunch: pulumi.Bool(true),
			Tags:                subnetTags,
		}, pulumi.Parent(vpc))
		if err != nil {
			return nil, err
		}
		publicSubnets[i] = subnet.ID()
	}

	// Create private subnets
	privateSubnets := make([]pulumi.IDOutput, 2)
	for i := 0; i < 2; i++ {
		subnetTags := mergeTags(args.Tags.PrivateSubnet, args.Tags.Default, fmt.Sprintf("private-subnet-%d", i))
		subnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("%s-private-subnet-%d", name, i), &ec2.SubnetArgs{
			VpcId:            vpcResource.ID(),
			CidrBlock:        pulumi.Sprintf("10.0.%d.0/24", i+2),
			AvailabilityZone: pulumi.Sprintf("us-east-1%s", string(rune('a'+i))),
			Tags:             subnetTags,
		}, pulumi.Parent(vpc))
		if err != nil {
			return nil, err
		}
		privateSubnets[i] = subnet.ID()
	}

	// Create NAT instance security group
	natSgTags := mergeTags(args.Tags.NatSecurityGroup, args.Tags.Default, "nat-sg")
	natSg, err := ec2.NewSecurityGroup(ctx, fmt.Sprintf("%s-nat-sg", name), &ec2.SecurityGroupArgs{
		VpcId: vpcResource.ID(),
		Egress: ec2.SecurityGroupEgressArray{
			&ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("-1"),
				FromPort:   pulumi.Int(0),
				ToPort:     pulumi.Int(0),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
		Ingress: ec2.SecurityGroupIngressArray{
			&ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("-1"),
				FromPort:   pulumi.Int(0),
				ToPort:     pulumi.Int(0),
				CidrBlocks: pulumi.StringArray{pulumi.String(args.CidrBlock)},
			},
		},
		Tags: natSgTags,
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	// Create NAT instance
	natInstanceTags := mergeTags(args.Tags.NatInstance, args.Tags.Default, "nat-instance")
	natInstance, err := ec2.NewInstance(ctx, fmt.Sprintf("%s-nat-instance", name), &ec2.InstanceArgs{
		InstanceType:        pulumi.String(args.InstanceType),
		Ami:                 amiId,
		VpcSecurityGroupIds: pulumi.StringArray{natSg.ID()},
		SubnetId:            publicSubnets[0],
		SourceDestCheck:     pulumi.Bool(false),
		Tags:                natInstanceTags,
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	// Create route tables
	publicRtTags := mergeTags(args.Tags.PublicRouteTable, args.Tags.Default, "public-rt")
	publicRt, err := ec2.NewRouteTable(ctx, fmt.Sprintf("%s-public-rt", name), &ec2.RouteTableArgs{
		VpcId: vpcResource.ID(),
		Routes: ec2.RouteTableRouteArray{
			&ec2.RouteTableRouteArgs{
				CidrBlock: pulumi.String("0.0.0.0/0"),
				GatewayId: igw.ID(),
			},
		},
		Tags: publicRtTags,
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	privateRtTags := mergeTags(args.Tags.PrivateRouteTable, args.Tags.Default, "private-rt")
	privateRt, err := ec2.NewRouteTable(ctx, fmt.Sprintf("%s-private-rt", name), &ec2.RouteTableArgs{
		VpcId: vpcResource.ID(),
		Routes: ec2.RouteTableRouteArray{
			&ec2.RouteTableRouteArgs{
				CidrBlock:          pulumi.String("0.0.0.0/0"),
				NetworkInterfaceId: natInstance.PrimaryNetworkInterfaceId,
			},
		},
		Tags: privateRtTags,
	}, pulumi.Parent(vpc))
	if err != nil {
		return nil, err
	}

	// Associate route tables with subnets
	for i, subnet := range publicSubnets {
		_, err := ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("%s-public-rta-%d", name, i), &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet,
			RouteTableId: publicRt.ID(),
		}, pulumi.Parent(vpc))
		if err != nil {
			return nil, err
		}
	}

	for i, subnet := range privateSubnets {
		_, err := ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("%s-private-rta-%d", name, i), &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet,
			RouteTableId: privateRt.ID(),
		}, pulumi.Parent(vpc))
		if err != nil {
			return nil, err
		}
	}

	// Set the component outputs
	vpc.VpcId = vpcResource.ID()
	vpc.PublicSubnetIds = pulumi.ToIDArrayOutput(publicSubnets)
	vpc.PrivateSubnetIds = pulumi.ToIDArrayOutput(privateSubnets)
	vpc.NatInstanceId = natInstance.ID()

	return vpc, nil
}
