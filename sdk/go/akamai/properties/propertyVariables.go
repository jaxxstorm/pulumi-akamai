// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package properties

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: akamai.properties.PropertyVariables has been deprecated in favor of akamai.PropertyVariables
type PropertyVariables struct {
	pulumi.CustomResourceState

	// JSON variables representation
	Json pulumi.StringOutput `pulumi:"json"`
	// Deprecated: The setting "akamai_property_variables" has been deprecated.
	Variables PropertyVariablesVariableArrayOutput `pulumi:"variables"`
}

// NewPropertyVariables registers a new resource with the given unique name, arguments, and options.
func NewPropertyVariables(ctx *pulumi.Context,
	name string, args *PropertyVariablesArgs, opts ...pulumi.ResourceOption) (*PropertyVariables, error) {
	if args == nil {
		args = &PropertyVariablesArgs{}
	}

	var resource PropertyVariables
	err := ctx.RegisterResource("akamai:properties/propertyVariables:PropertyVariables", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPropertyVariables gets an existing PropertyVariables resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPropertyVariables(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PropertyVariablesState, opts ...pulumi.ResourceOption) (*PropertyVariables, error) {
	var resource PropertyVariables
	err := ctx.ReadResource("akamai:properties/propertyVariables:PropertyVariables", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PropertyVariables resources.
type propertyVariablesState struct {
	// JSON variables representation
	Json *string `pulumi:"json"`
	// Deprecated: The setting "akamai_property_variables" has been deprecated.
	Variables []PropertyVariablesVariable `pulumi:"variables"`
}

type PropertyVariablesState struct {
	// JSON variables representation
	Json pulumi.StringPtrInput
	// Deprecated: The setting "akamai_property_variables" has been deprecated.
	Variables PropertyVariablesVariableArrayInput
}

func (PropertyVariablesState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyVariablesState)(nil)).Elem()
}

type propertyVariablesArgs struct {
	// Deprecated: The setting "akamai_property_variables" has been deprecated.
	Variables []PropertyVariablesVariable `pulumi:"variables"`
}

// The set of arguments for constructing a PropertyVariables resource.
type PropertyVariablesArgs struct {
	// Deprecated: The setting "akamai_property_variables" has been deprecated.
	Variables PropertyVariablesVariableArrayInput
}

func (PropertyVariablesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyVariablesArgs)(nil)).Elem()
}

type PropertyVariablesInput interface {
	pulumi.Input

	ToPropertyVariablesOutput() PropertyVariablesOutput
	ToPropertyVariablesOutputWithContext(ctx context.Context) PropertyVariablesOutput
}

func (*PropertyVariables) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyVariables)(nil))
}

func (i *PropertyVariables) ToPropertyVariablesOutput() PropertyVariablesOutput {
	return i.ToPropertyVariablesOutputWithContext(context.Background())
}

func (i *PropertyVariables) ToPropertyVariablesOutputWithContext(ctx context.Context) PropertyVariablesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyVariablesOutput)
}

func (i *PropertyVariables) ToPropertyVariablesPtrOutput() PropertyVariablesPtrOutput {
	return i.ToPropertyVariablesPtrOutputWithContext(context.Background())
}

func (i *PropertyVariables) ToPropertyVariablesPtrOutputWithContext(ctx context.Context) PropertyVariablesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyVariablesPtrOutput)
}

type PropertyVariablesPtrInput interface {
	pulumi.Input

	ToPropertyVariablesPtrOutput() PropertyVariablesPtrOutput
	ToPropertyVariablesPtrOutputWithContext(ctx context.Context) PropertyVariablesPtrOutput
}

type propertyVariablesPtrType PropertyVariablesArgs

func (*propertyVariablesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyVariables)(nil))
}

func (i *propertyVariablesPtrType) ToPropertyVariablesPtrOutput() PropertyVariablesPtrOutput {
	return i.ToPropertyVariablesPtrOutputWithContext(context.Background())
}

func (i *propertyVariablesPtrType) ToPropertyVariablesPtrOutputWithContext(ctx context.Context) PropertyVariablesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyVariablesPtrOutput)
}

// PropertyVariablesArrayInput is an input type that accepts PropertyVariablesArray and PropertyVariablesArrayOutput values.
// You can construct a concrete instance of `PropertyVariablesArrayInput` via:
//
//          PropertyVariablesArray{ PropertyVariablesArgs{...} }
type PropertyVariablesArrayInput interface {
	pulumi.Input

	ToPropertyVariablesArrayOutput() PropertyVariablesArrayOutput
	ToPropertyVariablesArrayOutputWithContext(context.Context) PropertyVariablesArrayOutput
}

type PropertyVariablesArray []PropertyVariablesInput

func (PropertyVariablesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PropertyVariables)(nil)).Elem()
}

func (i PropertyVariablesArray) ToPropertyVariablesArrayOutput() PropertyVariablesArrayOutput {
	return i.ToPropertyVariablesArrayOutputWithContext(context.Background())
}

func (i PropertyVariablesArray) ToPropertyVariablesArrayOutputWithContext(ctx context.Context) PropertyVariablesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyVariablesArrayOutput)
}

// PropertyVariablesMapInput is an input type that accepts PropertyVariablesMap and PropertyVariablesMapOutput values.
// You can construct a concrete instance of `PropertyVariablesMapInput` via:
//
//          PropertyVariablesMap{ "key": PropertyVariablesArgs{...} }
type PropertyVariablesMapInput interface {
	pulumi.Input

	ToPropertyVariablesMapOutput() PropertyVariablesMapOutput
	ToPropertyVariablesMapOutputWithContext(context.Context) PropertyVariablesMapOutput
}

type PropertyVariablesMap map[string]PropertyVariablesInput

func (PropertyVariablesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PropertyVariables)(nil)).Elem()
}

func (i PropertyVariablesMap) ToPropertyVariablesMapOutput() PropertyVariablesMapOutput {
	return i.ToPropertyVariablesMapOutputWithContext(context.Background())
}

func (i PropertyVariablesMap) ToPropertyVariablesMapOutputWithContext(ctx context.Context) PropertyVariablesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PropertyVariablesMapOutput)
}

type PropertyVariablesOutput struct{ *pulumi.OutputState }

func (PropertyVariablesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyVariables)(nil))
}

func (o PropertyVariablesOutput) ToPropertyVariablesOutput() PropertyVariablesOutput {
	return o
}

func (o PropertyVariablesOutput) ToPropertyVariablesOutputWithContext(ctx context.Context) PropertyVariablesOutput {
	return o
}

func (o PropertyVariablesOutput) ToPropertyVariablesPtrOutput() PropertyVariablesPtrOutput {
	return o.ToPropertyVariablesPtrOutputWithContext(context.Background())
}

func (o PropertyVariablesOutput) ToPropertyVariablesPtrOutputWithContext(ctx context.Context) PropertyVariablesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PropertyVariables) *PropertyVariables {
		return &v
	}).(PropertyVariablesPtrOutput)
}

type PropertyVariablesPtrOutput struct{ *pulumi.OutputState }

func (PropertyVariablesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PropertyVariables)(nil))
}

func (o PropertyVariablesPtrOutput) ToPropertyVariablesPtrOutput() PropertyVariablesPtrOutput {
	return o
}

func (o PropertyVariablesPtrOutput) ToPropertyVariablesPtrOutputWithContext(ctx context.Context) PropertyVariablesPtrOutput {
	return o
}

func (o PropertyVariablesPtrOutput) Elem() PropertyVariablesOutput {
	return o.ApplyT(func(v *PropertyVariables) PropertyVariables {
		if v != nil {
			return *v
		}
		var ret PropertyVariables
		return ret
	}).(PropertyVariablesOutput)
}

type PropertyVariablesArrayOutput struct{ *pulumi.OutputState }

func (PropertyVariablesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PropertyVariables)(nil))
}

func (o PropertyVariablesArrayOutput) ToPropertyVariablesArrayOutput() PropertyVariablesArrayOutput {
	return o
}

func (o PropertyVariablesArrayOutput) ToPropertyVariablesArrayOutputWithContext(ctx context.Context) PropertyVariablesArrayOutput {
	return o
}

func (o PropertyVariablesArrayOutput) Index(i pulumi.IntInput) PropertyVariablesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PropertyVariables {
		return vs[0].([]PropertyVariables)[vs[1].(int)]
	}).(PropertyVariablesOutput)
}

type PropertyVariablesMapOutput struct{ *pulumi.OutputState }

func (PropertyVariablesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PropertyVariables)(nil))
}

func (o PropertyVariablesMapOutput) ToPropertyVariablesMapOutput() PropertyVariablesMapOutput {
	return o
}

func (o PropertyVariablesMapOutput) ToPropertyVariablesMapOutputWithContext(ctx context.Context) PropertyVariablesMapOutput {
	return o
}

func (o PropertyVariablesMapOutput) MapIndex(k pulumi.StringInput) PropertyVariablesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PropertyVariables {
		return vs[0].(map[string]PropertyVariables)[vs[1].(string)]
	}).(PropertyVariablesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyVariablesInput)(nil)).Elem(), &PropertyVariables{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyVariablesPtrInput)(nil)).Elem(), &PropertyVariables{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyVariablesArrayInput)(nil)).Elem(), PropertyVariablesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PropertyVariablesMapInput)(nil)).Elem(), PropertyVariablesMap{})
	pulumi.RegisterOutputType(PropertyVariablesOutput{})
	pulumi.RegisterOutputType(PropertyVariablesPtrOutput{})
	pulumi.RegisterOutputType(PropertyVariablesArrayOutput{})
	pulumi.RegisterOutputType(PropertyVariablesMapOutput{})
}
