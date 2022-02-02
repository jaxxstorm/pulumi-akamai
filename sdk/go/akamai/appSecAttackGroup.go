// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Attack group
//
// Modify an attack group's action, conditions, and exceptions. Attack groups are collections of Kona Rule Set rules used to streamline the management of website protections.
//
// **Related API Endpoints**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/attack-groups/{attackGroupId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroup) *and* [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/attack-groups/{attackGroupId}/condition-exception](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception)
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-akamai/sdk/v2/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Documentation"
// 		configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = akamai.NewAppSecAttackGroup(ctx, "attackGroup", &akamai.AppSecAttackGroupArgs{
// 			ConfigId:           pulumi.Int(configuration.ConfigId),
// 			SecurityPolicyId:   pulumi.String("gms1_134637"),
// 			AttackGroup:        pulumi.String("SQL"),
// 			AttackGroupAction:  pulumi.String("deny"),
// 			ConditionException: readFileOrPanic(fmt.Sprintf("%v%v", path.Module, "/condition_exception.json")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppSecAttackGroup struct {
	pulumi.CustomResourceState

	// . Unique name of the attack group being modified.
	AttackGroup pulumi.StringOutput `pulumi:"attackGroup"`
	// . Action taken any time the attack group is triggered. Allowed values are:
	// - **alert**. Record information about the request.
	// - **deny**. Block the request,
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action.
	AttackGroupAction pulumi.StringOutput `pulumi:"attackGroupAction"`
	// . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
	ConditionException pulumi.StringPtrOutput `pulumi:"conditionException"`
	// . Unique identifier of the security configuration associated with the attack group being modified.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the attack group being modified.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
}

// NewAppSecAttackGroup registers a new resource with the given unique name, arguments, and options.
func NewAppSecAttackGroup(ctx *pulumi.Context,
	name string, args *AppSecAttackGroupArgs, opts ...pulumi.ResourceOption) (*AppSecAttackGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttackGroup == nil {
		return nil, errors.New("invalid value for required argument 'AttackGroup'")
	}
	if args.AttackGroupAction == nil {
		return nil, errors.New("invalid value for required argument 'AttackGroupAction'")
	}
	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	var resource AppSecAttackGroup
	err := ctx.RegisterResource("akamai:index/appSecAttackGroup:AppSecAttackGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecAttackGroup gets an existing AppSecAttackGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecAttackGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecAttackGroupState, opts ...pulumi.ResourceOption) (*AppSecAttackGroup, error) {
	var resource AppSecAttackGroup
	err := ctx.ReadResource("akamai:index/appSecAttackGroup:AppSecAttackGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecAttackGroup resources.
type appSecAttackGroupState struct {
	// . Unique name of the attack group being modified.
	AttackGroup *string `pulumi:"attackGroup"`
	// . Action taken any time the attack group is triggered. Allowed values are:
	// - **alert**. Record information about the request.
	// - **deny**. Block the request,
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action.
	AttackGroupAction *string `pulumi:"attackGroupAction"`
	// . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
	ConditionException *string `pulumi:"conditionException"`
	// . Unique identifier of the security configuration associated with the attack group being modified.
	ConfigId *int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the attack group being modified.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecAttackGroupState struct {
	// . Unique name of the attack group being modified.
	AttackGroup pulumi.StringPtrInput
	// . Action taken any time the attack group is triggered. Allowed values are:
	// - **alert**. Record information about the request.
	// - **deny**. Block the request,
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action.
	AttackGroupAction pulumi.StringPtrInput
	// . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
	ConditionException pulumi.StringPtrInput
	// . Unique identifier of the security configuration associated with the attack group being modified.
	ConfigId pulumi.IntPtrInput
	// . Unique identifier of the security policy associated with the attack group being modified.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecAttackGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecAttackGroupState)(nil)).Elem()
}

type appSecAttackGroupArgs struct {
	// . Unique name of the attack group being modified.
	AttackGroup string `pulumi:"attackGroup"`
	// . Action taken any time the attack group is triggered. Allowed values are:
	// - **alert**. Record information about the request.
	// - **deny**. Block the request,
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action.
	AttackGroupAction string `pulumi:"attackGroupAction"`
	// . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
	ConditionException *string `pulumi:"conditionException"`
	// . Unique identifier of the security configuration associated with the attack group being modified.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy associated with the attack group being modified.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecAttackGroup resource.
type AppSecAttackGroupArgs struct {
	// . Unique name of the attack group being modified.
	AttackGroup pulumi.StringInput
	// . Action taken any time the attack group is triggered. Allowed values are:
	// - **alert**. Record information about the request.
	// - **deny**. Block the request,
	// - **deny_custom_{custom_deny_id}**. Take the action specified by the custom deny.
	// - **none**. Take no action.
	AttackGroupAction pulumi.StringInput
	// . Path to a JSON file containing the conditions and exceptions to be assigned to the attack group. You can view a sample JSON file in the [Modify the exceptions of an attack group](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putattackgroupconditionexception) section of the Application Security API documentation.
	ConditionException pulumi.StringPtrInput
	// . Unique identifier of the security configuration associated with the attack group being modified.
	ConfigId pulumi.IntInput
	// . Unique identifier of the security policy associated with the attack group being modified.
	SecurityPolicyId pulumi.StringInput
}

func (AppSecAttackGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecAttackGroupArgs)(nil)).Elem()
}

type AppSecAttackGroupInput interface {
	pulumi.Input

	ToAppSecAttackGroupOutput() AppSecAttackGroupOutput
	ToAppSecAttackGroupOutputWithContext(ctx context.Context) AppSecAttackGroupOutput
}

func (*AppSecAttackGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecAttackGroup)(nil)).Elem()
}

func (i *AppSecAttackGroup) ToAppSecAttackGroupOutput() AppSecAttackGroupOutput {
	return i.ToAppSecAttackGroupOutputWithContext(context.Background())
}

func (i *AppSecAttackGroup) ToAppSecAttackGroupOutputWithContext(ctx context.Context) AppSecAttackGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecAttackGroupOutput)
}

// AppSecAttackGroupArrayInput is an input type that accepts AppSecAttackGroupArray and AppSecAttackGroupArrayOutput values.
// You can construct a concrete instance of `AppSecAttackGroupArrayInput` via:
//
//          AppSecAttackGroupArray{ AppSecAttackGroupArgs{...} }
type AppSecAttackGroupArrayInput interface {
	pulumi.Input

	ToAppSecAttackGroupArrayOutput() AppSecAttackGroupArrayOutput
	ToAppSecAttackGroupArrayOutputWithContext(context.Context) AppSecAttackGroupArrayOutput
}

type AppSecAttackGroupArray []AppSecAttackGroupInput

func (AppSecAttackGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecAttackGroup)(nil)).Elem()
}

func (i AppSecAttackGroupArray) ToAppSecAttackGroupArrayOutput() AppSecAttackGroupArrayOutput {
	return i.ToAppSecAttackGroupArrayOutputWithContext(context.Background())
}

func (i AppSecAttackGroupArray) ToAppSecAttackGroupArrayOutputWithContext(ctx context.Context) AppSecAttackGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecAttackGroupArrayOutput)
}

// AppSecAttackGroupMapInput is an input type that accepts AppSecAttackGroupMap and AppSecAttackGroupMapOutput values.
// You can construct a concrete instance of `AppSecAttackGroupMapInput` via:
//
//          AppSecAttackGroupMap{ "key": AppSecAttackGroupArgs{...} }
type AppSecAttackGroupMapInput interface {
	pulumi.Input

	ToAppSecAttackGroupMapOutput() AppSecAttackGroupMapOutput
	ToAppSecAttackGroupMapOutputWithContext(context.Context) AppSecAttackGroupMapOutput
}

type AppSecAttackGroupMap map[string]AppSecAttackGroupInput

func (AppSecAttackGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecAttackGroup)(nil)).Elem()
}

func (i AppSecAttackGroupMap) ToAppSecAttackGroupMapOutput() AppSecAttackGroupMapOutput {
	return i.ToAppSecAttackGroupMapOutputWithContext(context.Background())
}

func (i AppSecAttackGroupMap) ToAppSecAttackGroupMapOutputWithContext(ctx context.Context) AppSecAttackGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecAttackGroupMapOutput)
}

type AppSecAttackGroupOutput struct{ *pulumi.OutputState }

func (AppSecAttackGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecAttackGroup)(nil)).Elem()
}

func (o AppSecAttackGroupOutput) ToAppSecAttackGroupOutput() AppSecAttackGroupOutput {
	return o
}

func (o AppSecAttackGroupOutput) ToAppSecAttackGroupOutputWithContext(ctx context.Context) AppSecAttackGroupOutput {
	return o
}

type AppSecAttackGroupArrayOutput struct{ *pulumi.OutputState }

func (AppSecAttackGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecAttackGroup)(nil)).Elem()
}

func (o AppSecAttackGroupArrayOutput) ToAppSecAttackGroupArrayOutput() AppSecAttackGroupArrayOutput {
	return o
}

func (o AppSecAttackGroupArrayOutput) ToAppSecAttackGroupArrayOutputWithContext(ctx context.Context) AppSecAttackGroupArrayOutput {
	return o
}

func (o AppSecAttackGroupArrayOutput) Index(i pulumi.IntInput) AppSecAttackGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecAttackGroup {
		return vs[0].([]*AppSecAttackGroup)[vs[1].(int)]
	}).(AppSecAttackGroupOutput)
}

type AppSecAttackGroupMapOutput struct{ *pulumi.OutputState }

func (AppSecAttackGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecAttackGroup)(nil)).Elem()
}

func (o AppSecAttackGroupMapOutput) ToAppSecAttackGroupMapOutput() AppSecAttackGroupMapOutput {
	return o
}

func (o AppSecAttackGroupMapOutput) ToAppSecAttackGroupMapOutputWithContext(ctx context.Context) AppSecAttackGroupMapOutput {
	return o
}

func (o AppSecAttackGroupMapOutput) MapIndex(k pulumi.StringInput) AppSecAttackGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecAttackGroup {
		return vs[0].(map[string]*AppSecAttackGroup)[vs[1].(string)]
	}).(AppSecAttackGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecAttackGroupInput)(nil)).Elem(), &AppSecAttackGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecAttackGroupArrayInput)(nil)).Elem(), AppSecAttackGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecAttackGroupMapInput)(nil)).Elem(), AppSecAttackGroupMap{})
	pulumi.RegisterOutputType(AppSecAttackGroupOutput{})
	pulumi.RegisterOutputType(AppSecAttackGroupArrayOutput{})
	pulumi.RegisterOutputType(AppSecAttackGroupMapOutput{})
}
