// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Once you complete the Let’s Encrypt challenges, optionally use the `CpsDvValidation` resource to send the acknowledgement to CPS and inform it that tokens are ready for validation. You can also wait for CPS to check for the tokens, which it does on a regular schedule. Next, CPS automatically deploys the certificate on Staging, and eventually on the Production network.
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
// 		_, err := akamai.NewCpsDvValidation(ctx, "example", &akamai.CpsDvValidationArgs{
// 			EnrollmentId: pulumi.Any(akamai_cps_dv_enrollment.Example.Id),
// 			Sans:         akamai_cps_dv_enrollment.Example.Sans,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Argument reference
//
// The following arguments are supported:
//
// * `enrollmentId` (Required) - Unique identifier for the DV certificate enrollment.
// * `sans` - (Optional) The Subject Alternative Names (SAN) list for tracking changes on related enrollments. Whenever any SAN changes, the Akamai provider recreates this resource and sends another acknowledgement request to CPS.
//
// ## Attributes reference
//
// * `status` - The status of certificate validation.
type CpsDvValidation struct {
	pulumi.CustomResourceState

	EnrollmentId pulumi.IntOutput         `pulumi:"enrollmentId"`
	Sans         pulumi.StringArrayOutput `pulumi:"sans"`
	Status       pulumi.StringOutput      `pulumi:"status"`
}

// NewCpsDvValidation registers a new resource with the given unique name, arguments, and options.
func NewCpsDvValidation(ctx *pulumi.Context,
	name string, args *CpsDvValidationArgs, opts ...pulumi.ResourceOption) (*CpsDvValidation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnrollmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnrollmentId'")
	}
	var resource CpsDvValidation
	err := ctx.RegisterResource("akamai:index/cpsDvValidation:CpsDvValidation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCpsDvValidation gets an existing CpsDvValidation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCpsDvValidation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CpsDvValidationState, opts ...pulumi.ResourceOption) (*CpsDvValidation, error) {
	var resource CpsDvValidation
	err := ctx.ReadResource("akamai:index/cpsDvValidation:CpsDvValidation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CpsDvValidation resources.
type cpsDvValidationState struct {
	EnrollmentId *int     `pulumi:"enrollmentId"`
	Sans         []string `pulumi:"sans"`
	Status       *string  `pulumi:"status"`
}

type CpsDvValidationState struct {
	EnrollmentId pulumi.IntPtrInput
	Sans         pulumi.StringArrayInput
	Status       pulumi.StringPtrInput
}

func (CpsDvValidationState) ElementType() reflect.Type {
	return reflect.TypeOf((*cpsDvValidationState)(nil)).Elem()
}

type cpsDvValidationArgs struct {
	EnrollmentId int      `pulumi:"enrollmentId"`
	Sans         []string `pulumi:"sans"`
}

// The set of arguments for constructing a CpsDvValidation resource.
type CpsDvValidationArgs struct {
	EnrollmentId pulumi.IntInput
	Sans         pulumi.StringArrayInput
}

func (CpsDvValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cpsDvValidationArgs)(nil)).Elem()
}

type CpsDvValidationInput interface {
	pulumi.Input

	ToCpsDvValidationOutput() CpsDvValidationOutput
	ToCpsDvValidationOutputWithContext(ctx context.Context) CpsDvValidationOutput
}

func (*CpsDvValidation) ElementType() reflect.Type {
	return reflect.TypeOf((*CpsDvValidation)(nil))
}

func (i *CpsDvValidation) ToCpsDvValidationOutput() CpsDvValidationOutput {
	return i.ToCpsDvValidationOutputWithContext(context.Background())
}

func (i *CpsDvValidation) ToCpsDvValidationOutputWithContext(ctx context.Context) CpsDvValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvValidationOutput)
}

func (i *CpsDvValidation) ToCpsDvValidationPtrOutput() CpsDvValidationPtrOutput {
	return i.ToCpsDvValidationPtrOutputWithContext(context.Background())
}

func (i *CpsDvValidation) ToCpsDvValidationPtrOutputWithContext(ctx context.Context) CpsDvValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvValidationPtrOutput)
}

type CpsDvValidationPtrInput interface {
	pulumi.Input

	ToCpsDvValidationPtrOutput() CpsDvValidationPtrOutput
	ToCpsDvValidationPtrOutputWithContext(ctx context.Context) CpsDvValidationPtrOutput
}

type cpsDvValidationPtrType CpsDvValidationArgs

func (*cpsDvValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CpsDvValidation)(nil))
}

func (i *cpsDvValidationPtrType) ToCpsDvValidationPtrOutput() CpsDvValidationPtrOutput {
	return i.ToCpsDvValidationPtrOutputWithContext(context.Background())
}

func (i *cpsDvValidationPtrType) ToCpsDvValidationPtrOutputWithContext(ctx context.Context) CpsDvValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvValidationPtrOutput)
}

// CpsDvValidationArrayInput is an input type that accepts CpsDvValidationArray and CpsDvValidationArrayOutput values.
// You can construct a concrete instance of `CpsDvValidationArrayInput` via:
//
//          CpsDvValidationArray{ CpsDvValidationArgs{...} }
type CpsDvValidationArrayInput interface {
	pulumi.Input

	ToCpsDvValidationArrayOutput() CpsDvValidationArrayOutput
	ToCpsDvValidationArrayOutputWithContext(context.Context) CpsDvValidationArrayOutput
}

type CpsDvValidationArray []CpsDvValidationInput

func (CpsDvValidationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CpsDvValidation)(nil))
}

func (i CpsDvValidationArray) ToCpsDvValidationArrayOutput() CpsDvValidationArrayOutput {
	return i.ToCpsDvValidationArrayOutputWithContext(context.Background())
}

func (i CpsDvValidationArray) ToCpsDvValidationArrayOutputWithContext(ctx context.Context) CpsDvValidationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvValidationArrayOutput)
}

// CpsDvValidationMapInput is an input type that accepts CpsDvValidationMap and CpsDvValidationMapOutput values.
// You can construct a concrete instance of `CpsDvValidationMapInput` via:
//
//          CpsDvValidationMap{ "key": CpsDvValidationArgs{...} }
type CpsDvValidationMapInput interface {
	pulumi.Input

	ToCpsDvValidationMapOutput() CpsDvValidationMapOutput
	ToCpsDvValidationMapOutputWithContext(context.Context) CpsDvValidationMapOutput
}

type CpsDvValidationMap map[string]CpsDvValidationInput

func (CpsDvValidationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CpsDvValidation)(nil))
}

func (i CpsDvValidationMap) ToCpsDvValidationMapOutput() CpsDvValidationMapOutput {
	return i.ToCpsDvValidationMapOutputWithContext(context.Background())
}

func (i CpsDvValidationMap) ToCpsDvValidationMapOutputWithContext(ctx context.Context) CpsDvValidationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvValidationMapOutput)
}

type CpsDvValidationOutput struct {
	*pulumi.OutputState
}

func (CpsDvValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CpsDvValidation)(nil))
}

func (o CpsDvValidationOutput) ToCpsDvValidationOutput() CpsDvValidationOutput {
	return o
}

func (o CpsDvValidationOutput) ToCpsDvValidationOutputWithContext(ctx context.Context) CpsDvValidationOutput {
	return o
}

func (o CpsDvValidationOutput) ToCpsDvValidationPtrOutput() CpsDvValidationPtrOutput {
	return o.ToCpsDvValidationPtrOutputWithContext(context.Background())
}

func (o CpsDvValidationOutput) ToCpsDvValidationPtrOutputWithContext(ctx context.Context) CpsDvValidationPtrOutput {
	return o.ApplyT(func(v CpsDvValidation) *CpsDvValidation {
		return &v
	}).(CpsDvValidationPtrOutput)
}

type CpsDvValidationPtrOutput struct {
	*pulumi.OutputState
}

func (CpsDvValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CpsDvValidation)(nil))
}

func (o CpsDvValidationPtrOutput) ToCpsDvValidationPtrOutput() CpsDvValidationPtrOutput {
	return o
}

func (o CpsDvValidationPtrOutput) ToCpsDvValidationPtrOutputWithContext(ctx context.Context) CpsDvValidationPtrOutput {
	return o
}

type CpsDvValidationArrayOutput struct{ *pulumi.OutputState }

func (CpsDvValidationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CpsDvValidation)(nil))
}

func (o CpsDvValidationArrayOutput) ToCpsDvValidationArrayOutput() CpsDvValidationArrayOutput {
	return o
}

func (o CpsDvValidationArrayOutput) ToCpsDvValidationArrayOutputWithContext(ctx context.Context) CpsDvValidationArrayOutput {
	return o
}

func (o CpsDvValidationArrayOutput) Index(i pulumi.IntInput) CpsDvValidationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CpsDvValidation {
		return vs[0].([]CpsDvValidation)[vs[1].(int)]
	}).(CpsDvValidationOutput)
}

type CpsDvValidationMapOutput struct{ *pulumi.OutputState }

func (CpsDvValidationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CpsDvValidation)(nil))
}

func (o CpsDvValidationMapOutput) ToCpsDvValidationMapOutput() CpsDvValidationMapOutput {
	return o
}

func (o CpsDvValidationMapOutput) ToCpsDvValidationMapOutputWithContext(ctx context.Context) CpsDvValidationMapOutput {
	return o
}

func (o CpsDvValidationMapOutput) MapIndex(k pulumi.StringInput) CpsDvValidationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CpsDvValidation {
		return vs[0].(map[string]CpsDvValidation)[vs[1].(string)]
	}).(CpsDvValidationOutput)
}

func init() {
	pulumi.RegisterOutputType(CpsDvValidationOutput{})
	pulumi.RegisterOutputType(CpsDvValidationPtrOutput{})
	pulumi.RegisterOutputType(CpsDvValidationArrayOutput{})
	pulumi.RegisterOutputType(CpsDvValidationMapOutput{})
}
