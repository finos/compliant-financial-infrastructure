package components

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RDSInstanceArgs struct {
	InstanceClass       string
	AllocatedStorage    int
	Engine              string
	EngineVersion       string
	Username            string
	Password            string
	DbName              string
	PubliclyAccessible  bool
	SubnetIds           []string
	VpcSecurityGroupIds []string
	Tags                map[string]string
}

type RDSInstance struct {
	pulumi.ResourceState

	Instance *rds.Instance
	Endpoint pulumi.StringOutput
}

func NewRDSInstance(ctx *pulumi.Context, name string, args *RDSInstanceArgs, opts ...pulumi.ResourceOption) (*RDSInstance, error) {
	rdsInstance := &RDSInstance{}
	err := ctx.RegisterComponentResource("custom:database:RDSInstance", name, rdsInstance, opts...)
	if err != nil {
		return nil, err
	}

	// Create a subnet group for the RDS instance
	subnetGroup, err := rds.NewSubnetGroup(ctx, name+"-subnet-group", &rds.SubnetGroupArgs{
		SubnetIds: pulumi.ToStringArray(args.SubnetIds),
	}, pulumi.Parent(rdsInstance))
	if err != nil {
		return nil, err
	}

	// Create the RDS instance
	instance, err := rds.NewInstance(ctx, name, &rds.InstanceArgs{
		InstanceClass:       pulumi.String(args.InstanceClass),
		AllocatedStorage:    pulumi.Int(args.AllocatedStorage),
		Engine:              pulumi.String(args.Engine),
		EngineVersion:       pulumi.String(args.EngineVersion),
		Username:            pulumi.String(args.Username),
		Password:            pulumi.String(args.Password),
		DbName:              pulumi.String(args.DbName),
		SkipFinalSnapshot:   pulumi.Bool(true),
		PubliclyAccessible:  pulumi.Bool(args.PubliclyAccessible),
		DbSubnetGroupName:   subnetGroup.Name,
		VpcSecurityGroupIds: pulumi.ToStringArray(args.VpcSecurityGroupIds),
		Tags:                pulumi.ToStringMap(args.Tags),
	}, pulumi.Parent(rdsInstance))

	if err != nil {
		return nil, err
	}

	rdsInstance.Instance = instance
	rdsInstance.Endpoint = instance.Endpoint

	ctx.RegisterResourceOutputs(rdsInstance, pulumi.Map{
		"endpoint": rdsInstance.Endpoint,
	})

	return rdsInstance, nil
}
