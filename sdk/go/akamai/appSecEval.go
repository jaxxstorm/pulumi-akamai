// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Issues an evaluation mode command (`Start`, `Stop`, `Restart`, `Update`, or `Complete`) to a security configuration.
// Evaluation mode is used for testing and fine-tuning your Kona Rule Set rules and configuration settings.
// In evaluation mode rules are triggered by events, but the only thing those rules do is record the actions they *would* have taken had the event occurred on the production network.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/eval](https://techdocs.akamai.com/application-security/reference/post-policy-eval)
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			configuration, err := akamai.LookupAppSecConfiguration(ctx, &GetAppSecConfigurationArgs{
//				Name: pulumi.StringRef("Documentation"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			evalOperation, err := akamai.NewAppSecEval(ctx, "evalOperation", &akamai.AppSecEvalArgs{
//				ConfigId:         pulumi.Int(configuration.ConfigId),
//				SecurityPolicyId: pulumi.String("gms1_134637"),
//				EvalOperation:    pulumi.String("START"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("evalModeEvaluatingRuleset", evalOperation.EvaluatingRuleset)
//			ctx.Export("evalModeExpirationDate", evalOperation.ExpirationDate)
//			ctx.Export("evalModeCurrentRuleset", evalOperation.CurrentRuleset)
//			ctx.Export("evalModeStatus", evalOperation.EvalStatus)
//			return nil
//		})
//	}
//
// ```
// ## Output Options
//
// The following options can be used to determine the information returned, and how that returned information is formatted:
//
// - `evaluatingRuleset`. Versioning information for the Kona Rule Set being evaluated.
// - `expirationDate`. Date when the evaluation period ends.
// - `currentRuleset`. Versioning information for the Kona Rule Set currently in use on the production network.
// - `evalStatus`. If **true**, an evaluation is currently in progress; if **false**, evaluation is either paused or is not running.
type AppSecEval struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// Versioning information for the Kona Rule Set currently in use in production
	CurrentRuleset pulumi.StringOutput `pulumi:"currentRuleset"`
	// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
	EvalMode pulumi.StringPtrOutput `pulumi:"evalMode"`
	// . Evaluation mode operation. Allowed values are:
	// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
	// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
	// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
	// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
	// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
	EvalOperation pulumi.StringOutput `pulumi:"evalOperation"`
	// Whether an evaluation is currently in progress
	EvalStatus pulumi.StringOutput `pulumi:"evalStatus"`
	// Versioning information for the Kona Rule Set being evaluated
	EvaluatingRuleset pulumi.StringOutput `pulumi:"evaluatingRuleset"`
	// Date when the evaluation period ends
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// . Unique identifier of the security policy associated with the evaluation process.
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
	// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
	ConfigId *int `pulumi:"configId"`
	// Versioning information for the Kona Rule Set currently in use in production
	CurrentRuleset *string `pulumi:"currentRuleset"`
	// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
	EvalMode *string `pulumi:"evalMode"`
	// . Evaluation mode operation. Allowed values are:
	// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
	// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
	// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
	// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
	// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
	EvalOperation *string `pulumi:"evalOperation"`
	// Whether an evaluation is currently in progress
	EvalStatus *string `pulumi:"evalStatus"`
	// Versioning information for the Kona Rule Set being evaluated
	EvaluatingRuleset *string `pulumi:"evaluatingRuleset"`
	// Date when the evaluation period ends
	ExpirationDate *string `pulumi:"expirationDate"`
	// . Unique identifier of the security policy associated with the evaluation process.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
}

type AppSecEvalState struct {
	// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
	ConfigId pulumi.IntPtrInput
	// Versioning information for the Kona Rule Set currently in use in production
	CurrentRuleset pulumi.StringPtrInput
	// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
	EvalMode pulumi.StringPtrInput
	// . Evaluation mode operation. Allowed values are:
	// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
	// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
	// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
	// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
	// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
	EvalOperation pulumi.StringPtrInput
	// Whether an evaluation is currently in progress
	EvalStatus pulumi.StringPtrInput
	// Versioning information for the Kona Rule Set being evaluated
	EvaluatingRuleset pulumi.StringPtrInput
	// Date when the evaluation period ends
	ExpirationDate pulumi.StringPtrInput
	// . Unique identifier of the security policy associated with the evaluation process.
	SecurityPolicyId pulumi.StringPtrInput
}

func (AppSecEvalState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecEvalState)(nil)).Elem()
}

type appSecEvalArgs struct {
	// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
	ConfigId int `pulumi:"configId"`
	// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
	EvalMode *string `pulumi:"evalMode"`
	// . Evaluation mode operation. Allowed values are:
	// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
	// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
	// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
	// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
	// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
	EvalOperation string `pulumi:"evalOperation"`
	// . Unique identifier of the security policy associated with the evaluation process.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
}

// The set of arguments for constructing a AppSecEval resource.
type AppSecEvalArgs struct {
	// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
	ConfigId pulumi.IntInput
	// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
	EvalMode pulumi.StringPtrInput
	// . Evaluation mode operation. Allowed values are:
	// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
	// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
	// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
	// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
	// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
	EvalOperation pulumi.StringInput
	// . Unique identifier of the security policy associated with the evaluation process.
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
	return reflect.TypeOf((**AppSecEval)(nil)).Elem()
}

func (i *AppSecEval) ToAppSecEvalOutput() AppSecEvalOutput {
	return i.ToAppSecEvalOutputWithContext(context.Background())
}

func (i *AppSecEval) ToAppSecEvalOutputWithContext(ctx context.Context) AppSecEvalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalOutput)
}

// AppSecEvalArrayInput is an input type that accepts AppSecEvalArray and AppSecEvalArrayOutput values.
// You can construct a concrete instance of `AppSecEvalArrayInput` via:
//
//	AppSecEvalArray{ AppSecEvalArgs{...} }
type AppSecEvalArrayInput interface {
	pulumi.Input

	ToAppSecEvalArrayOutput() AppSecEvalArrayOutput
	ToAppSecEvalArrayOutputWithContext(context.Context) AppSecEvalArrayOutput
}

type AppSecEvalArray []AppSecEvalInput

func (AppSecEvalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecEval)(nil)).Elem()
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
//	AppSecEvalMap{ "key": AppSecEvalArgs{...} }
type AppSecEvalMapInput interface {
	pulumi.Input

	ToAppSecEvalMapOutput() AppSecEvalMapOutput
	ToAppSecEvalMapOutputWithContext(context.Context) AppSecEvalMapOutput
}

type AppSecEvalMap map[string]AppSecEvalInput

func (AppSecEvalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecEval)(nil)).Elem()
}

func (i AppSecEvalMap) ToAppSecEvalMapOutput() AppSecEvalMapOutput {
	return i.ToAppSecEvalMapOutputWithContext(context.Background())
}

func (i AppSecEvalMap) ToAppSecEvalMapOutputWithContext(ctx context.Context) AppSecEvalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecEvalMapOutput)
}

type AppSecEvalOutput struct{ *pulumi.OutputState }

func (AppSecEvalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecEval)(nil)).Elem()
}

func (o AppSecEvalOutput) ToAppSecEvalOutput() AppSecEvalOutput {
	return o
}

func (o AppSecEvalOutput) ToAppSecEvalOutputWithContext(ctx context.Context) AppSecEvalOutput {
	return o
}

// . Unique identifier of the security configuration where evaluation mode will take place (or is currently taking place).
func (o AppSecEvalOutput) ConfigId() pulumi.IntOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.IntOutput { return v.ConfigId }).(pulumi.IntOutput)
}

// Versioning information for the Kona Rule Set currently in use in production
func (o AppSecEvalOutput) CurrentRuleset() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.StringOutput { return v.CurrentRuleset }).(pulumi.StringOutput)
}

// . Set to **ASE_AUTO** to have your Kona Rule Set rules automatically updated during the evaluation period; set to **ASE_MANUAL** if you want to manually update your evaluation rules. Note that this option is only available to organizations running the Adaptive Security Engine (ASE) beta. For more information about ASE, please contact your Akamai representative.
func (o AppSecEvalOutput) EvalMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.StringPtrOutput { return v.EvalMode }).(pulumi.StringPtrOutput)
}

// . Evaluation mode operation. Allowed values are:
// - **START**. Starts evaluation mode. By default, evaluation mode runs for four weeks.
// - **STOP**, Pauses evaluation mode without upgrading the Kona Rule Set on your production network.
// - **RESTART**. Resumes an evaluation trial that was paused by using the **STOP** command.
// - **UPDATE**. Upgrades the Kona Rule Set rules in the evaluation ruleset to their latest versions.
// - **COMPLETE**. Concludes the evaluation period (even if the four-week trial mode is not over) and automatically upgrades the Kona Rule Set on your production network to the same rule set you just finished evaluating.
func (o AppSecEvalOutput) EvalOperation() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.StringOutput { return v.EvalOperation }).(pulumi.StringOutput)
}

// Whether an evaluation is currently in progress
func (o AppSecEvalOutput) EvalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.StringOutput { return v.EvalStatus }).(pulumi.StringOutput)
}

// Versioning information for the Kona Rule Set being evaluated
func (o AppSecEvalOutput) EvaluatingRuleset() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.StringOutput { return v.EvaluatingRuleset }).(pulumi.StringOutput)
}

// Date when the evaluation period ends
func (o AppSecEvalOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

// . Unique identifier of the security policy associated with the evaluation process.
func (o AppSecEvalOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AppSecEval) pulumi.StringOutput { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

type AppSecEvalArrayOutput struct{ *pulumi.OutputState }

func (AppSecEvalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecEval)(nil)).Elem()
}

func (o AppSecEvalArrayOutput) ToAppSecEvalArrayOutput() AppSecEvalArrayOutput {
	return o
}

func (o AppSecEvalArrayOutput) ToAppSecEvalArrayOutputWithContext(ctx context.Context) AppSecEvalArrayOutput {
	return o
}

func (o AppSecEvalArrayOutput) Index(i pulumi.IntInput) AppSecEvalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppSecEval {
		return vs[0].([]*AppSecEval)[vs[1].(int)]
	}).(AppSecEvalOutput)
}

type AppSecEvalMapOutput struct{ *pulumi.OutputState }

func (AppSecEvalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecEval)(nil)).Elem()
}

func (o AppSecEvalMapOutput) ToAppSecEvalMapOutput() AppSecEvalMapOutput {
	return o
}

func (o AppSecEvalMapOutput) ToAppSecEvalMapOutputWithContext(ctx context.Context) AppSecEvalMapOutput {
	return o
}

func (o AppSecEvalMapOutput) MapIndex(k pulumi.StringInput) AppSecEvalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppSecEval {
		return vs[0].(map[string]*AppSecEval)[vs[1].(string)]
	}).(AppSecEvalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalInput)(nil)).Elem(), &AppSecEval{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalArrayInput)(nil)).Elem(), AppSecEvalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecEvalMapInput)(nil)).Elem(), AppSecEvalMap{})
	pulumi.RegisterOutputType(AppSecEvalOutput{})
	pulumi.RegisterOutputType(AppSecEvalArrayOutput{})
	pulumi.RegisterOutputType(AppSecEvalMapOutput{})
}
