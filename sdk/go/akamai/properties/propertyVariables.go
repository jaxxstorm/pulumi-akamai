// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package properties

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `properties.PropertyVariables` allows you to implement dynamic functionality. You can perform conditional logic based on the variable’s value, and catch any unforeseen errors that execute on the edge at runtime.
//
// Typical uses for variables include:
//
// * Simplify configurations by reducing the number of rules and behaviors.
// * Improve self serviceability by replacing or extending advanced metadata.
// * Automate redirects, forward path rewrites, HTTP header and cookie manipulation.
// * Move origin functionality to the edge.
//
// ## Example Usage
// ### Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/go/akamai/properties"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := properties.NewPropertyVariables(ctx, "origin", &properties.PropertyVariablesArgs{
// 			Variables: properties.PropertyVariablesVariableArray{
// 				&properties.PropertyVariablesVariableArgs{
// 					Variables: properties.PropertyVariablesVariableVariableArray{
// 						&properties.PropertyVariablesVariableVariableArgs{
// 							Description: pulumi.String("Origin Hostname"),
// 							Hidden:      pulumi.Bool(true),
// 							Name:        pulumi.String("PMUSER_ORIGIN"),
// 							Sensitive:   pulumi.Bool(true),
// 							Value:       pulumi.String("origin.example.org"),
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PropertyVariables struct {
	pulumi.CustomResourceState

	// JSON variables representation
	Json      pulumi.StringOutput                  `pulumi:"json"`
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
	Json      *string                     `pulumi:"json"`
	Variables []PropertyVariablesVariable `pulumi:"variables"`
}

type PropertyVariablesState struct {
	// JSON variables representation
	Json      pulumi.StringPtrInput
	Variables PropertyVariablesVariableArrayInput
}

func (PropertyVariablesState) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyVariablesState)(nil)).Elem()
}

type propertyVariablesArgs struct {
	Variables []PropertyVariablesVariable `pulumi:"variables"`
}

// The set of arguments for constructing a PropertyVariables resource.
type PropertyVariablesArgs struct {
	Variables PropertyVariablesVariableArrayInput
}

func (PropertyVariablesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*propertyVariablesArgs)(nil)).Elem()
}
