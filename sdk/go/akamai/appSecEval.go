// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `AppSecEval` resource to perform evaluation mode operations such as Start, Stop, Restart, Update, or Complete.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/v2/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Security_configuration
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &akamai.LookupAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		evalOperation, err := akamai.NewAppSecEval(ctx, "evalOperation", &akamai.AppSecEvalArgs{
// 			ConfigId:         pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId: pulumi.Any(_var.Security_policy_id),
// 			EvalOperation:    pulumi.Any(_var.Eval_operation),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("evalModeEvaluatingRuleset", evalOperation.EvaluatingRuleset)
// 		ctx.Export("evalModeExpirationDate", evalOperation.ExpirationDate)
// 		ctx.Export("evalModeCurrentRuleset", evalOperation.CurrentRuleset)
// 		ctx.Export("evalModeStatus", evalOperation.EvalStatus)
// 		return nil
// 	})
// }
// ```
type AppSecEval struct {
	pulumi.CustomResourceState

	// The ID of the security configuration to use.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// The set of rules currently in effect.
	CurrentRuleset pulumi.StringOutput `pulumi:"currentRuleset"`
	// The operation to perform: START, STOP, RESTART, UPDATE, or COMPLETE.
	EvalOperation pulumi.StringOutput `pulumi:"evalOperation"`
	// Either `enabled` if an evaluation is currently in progress (that is, if the `evalOperation` parameter was `START`, `RESTART`, or `COMPLETE`) or `disabled` otherwise (that is, if the `evalOperation` parameter was `STOP` or `UPDATE`).
	EvalStatus pulumi.StringOutput `pulumi:"evalStatus"`
	// The set of rules being evaluated.
	EvaluatingRuleset pulumi.StringOutput `pulumi:"evaluatingRuleset"`
	// The date on which the evaluation period ends.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecEval registers a new resource with the given unique name, arguments, and options.
func NewAppSecEval(ctx *pulumi.Context,
	name string, args *AppSecEvalArgs, opts ...pulumi.ResourceOption) (*AppSecEval, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.EvalOperation == nil {
		return nil, errors.New("invalid value for required argument 'EvalOperation'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecEval
	err := ctx.RegisterResource("akamai:index/appSecEval:AppSecEval", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecEval gets an existing AppSecEval resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecEval(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecEvalState, opts ...pulumi.ResourceOption) (*AppSecEval, error) {
	var resource AppSecEval
	err := ctx.ReadResource("akamai:index/appSecEval:AppSecEval", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecEval resources.
type appSecEvalState struct {
	// The ID of the security configuration to use.
	ConfigId *int `pulumi:"configId"`
	// The set of rules currently in effect.
	CurrentRuleset *string `pulumi:"currentRuleset"`
	// The operation to perform: START, STOP, RESTART, UPDATE, or COMPLETE.
	EvalOperation *string `pulumi:"evalOperation"`
	// Either `enabled` if an evaluation is currently in progress (that is, if the `evalOperation` parameter was `START`, `RESTART`, or `COMPLETE`) or `disabled` otherwise (that is, if the `evalOperation` parameter was `STOP` or `UPDATE`).
	EvalStatus *string `pulumi:"evalStatus"`
	// The set of rules being evaluated.
	EvaluatingRuleset *string `pulumi:"evaluatingRuleset"`
	// The date on which the evaluation period ends.
	ExpirationDate *string `pulumi:"expirationDate"`
	// The ID of the security policy to use.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecEvalState struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntPtrInput
	// The set of rules currently in effect.
	CurrentRuleset pulumi.StringPtrInput
	// The operation to perform: START, STOP, RESTART, UPDATE, or COMPLETE.
	EvalOperation pulumi.StringPtrInput
	// Either `enabled` if an evaluation is currently in progress (that is, if the `evalOperation` parameter was `START`, `RESTART`, or `COMPLETE`) or `disabled` otherwise (that is, if the `evalOperation` parameter was `STOP` or `UPDATE`).
	EvalStatus pulumi.StringPtrInput
	// The set of rules being evaluated.
	EvaluatingRuleset pulumi.StringPtrInput
	// The date on which the evaluation period ends.
	ExpirationDate pulumi.StringPtrInput
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecEvalState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecEvalState)(nil)).Elem()
}

type appSecEvalArgs struct {
	// The ID of the security configuration to use.
	ConfigId int `pulumi:"configId"`
	// The operation to perform: START, STOP, RESTART, UPDATE, or COMPLETE.
	EvalOperation string `pulumi:"evalOperation"`
	// The ID of the security policy to use.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecEval resource.
type AppSecEvalArgs struct {
	// The ID of the security configuration to use.
	ConfigId pulumi.IntInput
	// The operation to perform: START, STOP, RESTART, UPDATE, or COMPLETE.
	EvalOperation pulumi.StringInput
	// The ID of the security policy to use.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecEvalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecEvalArgs)(nil)).Elem()
}

type AppSecEvalInput interface {
	pulumi.Input

	ToAppSecEvalOutput() AppSecEvalOutput
	ToAppSecEvalOutputWithContext(ctx context.Context) AppSecEvalOutput
}

func (*AppSecEval) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecEval)(nil))
}

func (i *AppSecEval) ToAppSecEvalOutput() AppSecEvalOutput {
	return i.ToAppSecEvalOutputWithContext(context.Background())
}

func (i *AppSecEval) ToAppSecEvalOutputWithContext(ctx context.Context) AppSecEvalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalOutput)
}

func (i *AppSecEval) ToAppSecEvalPtrOutput() AppSecEvalPtrOutput {
	return i.ToAppSecEvalPtrOutputWithContext(context.Background())
}

func (i *AppSecEval) ToAppSecEvalPtrOutputWithContext(ctx context.Context) AppSecEvalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalPtrOutput)
}

type AppSecEvalPtrInput interface {
	pulumi.Input

	ToAppSecEvalPtrOutput() AppSecEvalPtrOutput
	ToAppSecEvalPtrOutputWithContext(ctx context.Context) AppSecEvalPtrOutput
}

type appSecEvalPtrType AppSecEvalArgs

func (*appSecEvalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecEval)(nil))
}

func (i *appSecEvalPtrType) ToAppSecEvalPtrOutput() AppSecEvalPtrOutput {
	return i.ToAppSecEvalPtrOutputWithContext(context.Background())
}

func (i *appSecEvalPtrType) ToAppSecEvalPtrOutputWithContext(ctx context.Context) AppSecEvalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalPtrOutput)
}

// AppSecEvalArrayInput is an input type that accepts AppSecEvalArray and AppSecEvalArrayOutput values.
// You can construct a concrete instance of `AppSecEvalArrayInput` via:
//
//          AppSecEvalArray{ AppSecEvalArgs{...} }
type AppSecEvalArrayInput interface {
	pulumi.Input

	ToAppSecEvalArrayOutput() AppSecEvalArrayOutput
	ToAppSecEvalArrayOutputWithContext(context.Context) AppSecEvalArrayOutput
}

type AppSecEvalArray []AppSecEvalInput

func (AppSecEvalArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AppSecEval)(nil))
}

func (i AppSecEvalArray) ToAppSecEvalArrayOutput() AppSecEvalArrayOutput {
	return i.ToAppSecEvalArrayOutputWithContext(context.Background())
}

func (i AppSecEvalArray) ToAppSecEvalArrayOutputWithContext(ctx context.Context) AppSecEvalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalArrayOutput)
}

// AppSecEvalMapInput is an input type that accepts AppSecEvalMap and AppSecEvalMapOutput values.
// You can construct a concrete instance of `AppSecEvalMapInput` via:
//
//          AppSecEvalMap{ "key": AppSecEvalArgs{...} }
type AppSecEvalMapInput interface {
	pulumi.Input

	ToAppSecEvalMapOutput() AppSecEvalMapOutput
	ToAppSecEvalMapOutputWithContext(context.Context) AppSecEvalMapOutput
}

type AppSecEvalMap map[string]AppSecEvalInput

func (AppSecEvalMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AppSecEval)(nil))
}

func (i AppSecEvalMap) ToAppSecEvalMapOutput() AppSecEvalMapOutput {
	return i.ToAppSecEvalMapOutputWithContext(context.Background())
}

func (i AppSecEvalMap) ToAppSecEvalMapOutputWithContext(ctx context.Context) AppSecEvalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalMapOutput)
}

type AppSecEvalOutput struct {
	*pulumi.OutputState
}

func (AppSecEvalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecEval)(nil))
}

func (o AppSecEvalOutput) ToAppSecEvalOutput() AppSecEvalOutput {
	return o
}

func (o AppSecEvalOutput) ToAppSecEvalOutputWithContext(ctx context.Context) AppSecEvalOutput {
	return o
}

func (o AppSecEvalOutput) ToAppSecEvalPtrOutput() AppSecEvalPtrOutput {
	return o.ToAppSecEvalPtrOutputWithContext(context.Background())
}

func (o AppSecEvalOutput) ToAppSecEvalPtrOutputWithContext(ctx context.Context) AppSecEvalPtrOutput {
	return o.ApplyT(func(v AppSecEval) *AppSecEval {
		return &v
	}).(AppSecEvalPtrOutput)
}

type AppSecEvalPtrOutput struct {
	*pulumi.OutputState
}

func (AppSecEvalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecEval)(nil))
}

func (o AppSecEvalPtrOutput) ToAppSecEvalPtrOutput() AppSecEvalPtrOutput {
	return o
}

func (o AppSecEvalPtrOutput) ToAppSecEvalPtrOutputWithContext(ctx context.Context) AppSecEvalPtrOutput {
	return o
}

type AppSecEvalArrayOutput struct{ *pulumi.OutputState }

func (AppSecEvalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecEval)(nil))
}

func (o AppSecEvalArrayOutput) ToAppSecEvalArrayOutput() AppSecEvalArrayOutput {
	return o
}

func (o AppSecEvalArrayOutput) ToAppSecEvalArrayOutputWithContext(ctx context.Context) AppSecEvalArrayOutput {
	return o
}

func (o AppSecEvalArrayOutput) Index(i pulumi.IntInput) AppSecEvalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecEval {
		return vs[0].([]AppSecEval)[vs[1].(int)]
	}).(AppSecEvalOutput)
}

type AppSecEvalMapOutput struct{ *pulumi.OutputState }

func (AppSecEvalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecEval)(nil))
}

func (o AppSecEvalMapOutput) ToAppSecEvalMapOutput() AppSecEvalMapOutput {
	return o
}

func (o AppSecEvalMapOutput) ToAppSecEvalMapOutputWithContext(ctx context.Context) AppSecEvalMapOutput {
	return o
}

func (o AppSecEvalMapOutput) MapIndex(k pulumi.StringInput) AppSecEvalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecEval {
		return vs[0].(map[string]AppSecEval)[vs[1].(string)]
	}).(AppSecEvalOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSecEvalOutput{})
	pulumi.RegisterOutputType(AppSecEvalPtrOutput{})
	pulumi.RegisterOutputType(AppSecEvalArrayOutput{})
	pulumi.RegisterOutputType(AppSecEvalMapOutput{})
}