// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AppSecConfigurationVersionClone struct {
	pulumi.CustomResourceState

	ConfigId          pulumi.IntOutput     `pulumi:"configId"`
	CreateFromVersion pulumi.IntOutput     `pulumi:"createFromVersion"`
	RuleUpdate        pulumi.BoolPtrOutput `pulumi:"ruleUpdate"`
	// Version of cloned configuration
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewAppSecConfigurationVersionClone registers a new resource with the given unique name, arguments, and options.
func NewAppSecConfigurationVersionClone(ctx *pulumi.Context,
	name string, args *AppSecConfigurationVersionCloneArgs, opts ...pulumi.ResourceOption) (*AppSecConfigurationVersionClone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.CreateFromVersion == nil {
		return nil, errors.New("invalid value for required argument 'CreateFromVersion'")
	}
	var resource AppSecConfigurationVersionClone
	err := ctx.RegisterResource("akamai:index/appSecConfigurationVersionClone:AppSecConfigurationVersionClone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecConfigurationVersionClone gets an existing AppSecConfigurationVersionClone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecConfigurationVersionClone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecConfigurationVersionCloneState, opts ...pulumi.ResourceOption) (*AppSecConfigurationVersionClone, error) {
	var resource AppSecConfigurationVersionClone
	err := ctx.ReadResource("akamai:index/appSecConfigurationVersionClone:AppSecConfigurationVersionClone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecConfigurationVersionClone resources.
type appSecConfigurationVersionCloneState struct {
	ConfigId          *int  `pulumi:"configId"`
	CreateFromVersion *int  `pulumi:"createFromVersion"`
	RuleUpdate        *bool `pulumi:"ruleUpdate"`
	// Version of cloned configuration
	Version *int `pulumi:"version"`
}

type AppSecConfigurationVersionCloneState struct {
	ConfigId          pulumi.IntPtrInput
	CreateFromVersion pulumi.IntPtrInput
	RuleUpdate        pulumi.BoolPtrInput
	// Version of cloned configuration
	Version pulumi.IntPtrInput
}

func (AppSecConfigurationVersionCloneState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecConfigurationVersionCloneState)(nil)).Elem()
}

type appSecConfigurationVersionCloneArgs struct {
	ConfigId          int   `pulumi:"configId"`
	CreateFromVersion int   `pulumi:"createFromVersion"`
	RuleUpdate        *bool `pulumi:"ruleUpdate"`
}

// The set of arguments for constructing a AppSecConfigurationVersionClone resource.
type AppSecConfigurationVersionCloneArgs struct {
	ConfigId          pulumi.IntInput
	CreateFromVersion pulumi.IntInput
	RuleUpdate        pulumi.BoolPtrInput
}

func (AppSecConfigurationVersionCloneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecConfigurationVersionCloneArgs)(nil)).Elem()
}

type AppSecConfigurationVersionCloneInput interface {
	pulumi.Input

	ToAppSecConfigurationVersionCloneOutput() AppSecConfigurationVersionCloneOutput
	ToAppSecConfigurationVersionCloneOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneOutput
}

func (*AppSecConfigurationVersionClone) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecConfigurationVersionClone)(nil))
}

func (i *AppSecConfigurationVersionClone) ToAppSecConfigurationVersionCloneOutput() AppSecConfigurationVersionCloneOutput {
	return i.ToAppSecConfigurationVersionCloneOutputWithContext(context.Background())
}

func (i *AppSecConfigurationVersionClone) ToAppSecConfigurationVersionCloneOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecConfigurationVersionCloneOutput)
}

type AppSecConfigurationVersionCloneOutput struct {
	*pulumi.OutputState
}

func (AppSecConfigurationVersionCloneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecConfigurationVersionClone)(nil))
}

func (o AppSecConfigurationVersionCloneOutput) ToAppSecConfigurationVersionCloneOutput() AppSecConfigurationVersionCloneOutput {
	return o
}

func (o AppSecConfigurationVersionCloneOutput) ToAppSecConfigurationVersionCloneOutputWithContext(ctx context.Context) AppSecConfigurationVersionCloneOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppSecConfigurationVersionCloneOutput{})
}
