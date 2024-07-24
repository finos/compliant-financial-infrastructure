package components

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
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
	SubnetIds           pulumi.StringArrayInput
	VpcSecurityGroupIds pulumi.StringArrayInput
	Tags                map[string]string
	Encryption          bool
	CreateKey           bool
}

type RDSInstance struct {
	pulumi.ResourceState

	Instance *rds.Instance
	Endpoint pulumi.StringOutput
	KMSKey   *kms.Key
}

func NewRDSInstance(ctx *pulumi.Context, name string, args *RDSInstanceArgs, opts ...pulumi.ResourceOption) (*RDSInstance, error) {
	rdsInstance := &RDSInstance{}
	err := ctx.RegisterComponentResource("custom:database:RDSInstance", name, rdsInstance, opts...)
	if err != nil {
		return nil, err
	}

	// Create a subnet group for the RDS instance
	subnetGroup, err := rds.NewSubnetGroup(ctx, name+"-subnet-group", &rds.SubnetGroupArgs{
		SubnetIds: args.SubnetIds,
	}, pulumi.Parent(rdsInstance))
	if err != nil {
		return nil, err
	}

	instanceArgs := &rds.InstanceArgs{
		InstanceClass:       pulumi.String(args.InstanceClass),
		AllocatedStorage:    pulumi.Int(args.AllocatedStorage),
		Engine:              pulumi.String(args.Engine),
		EngineVersion:       pulumi.String(args.EngineVersion),
		Username:            pulumi.String(args.Username),
		Password:            pulumi.String(args.Password),
		DbName:              pulumi.String(args.DbName),
		PubliclyAccessible:  pulumi.Bool(args.PubliclyAccessible),
		DbSubnetGroupName:   subnetGroup.Name,
		VpcSecurityGroupIds: args.VpcSecurityGroupIds,
		SkipFinalSnapshot:   pulumi.Bool(true),
		Tags:                pulumi.ToStringMap(args.Tags),
	}

	// Handle encryption
	if args.Encryption {
		instanceArgs.StorageEncrypted = pulumi.BoolPtr(true)
		if args.CreateKey {
			key, err := kms.NewKey(ctx, name+"-kms-key", &kms.KeyArgs{
				Description: pulumi.String("KMS key for RDS instance " + name),
				Tags:        pulumi.ToStringMap(args.Tags),
			}, pulumi.Parent(rdsInstance))
			if err != nil {
				return nil, err
			}
			instanceArgs.KmsKeyId = key.Arn
		}
	}

	// Create the RDS instance
	instance, err := rds.NewInstance(ctx, name, instanceArgs, pulumi.Parent(rdsInstance))
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
