// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Scopes**: Security policy
//
// Renames an existing security policy. Note that you can only change the name of the policy: once issued, the security policy ID can't be modified.
//
// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}](https://developer.akamai.com/api/cloud_security/application_security/v1.html#putsecuritypolicy)
type AppSecSecurityPolicyRename struct {
	pulumi.CustomResourceState

	// . Unique identifier of the security configuration associated with the security policy being renamed.
	ConfigId pulumi.IntOutput `pulumi:"configId"`
	// . Unique identifier of the security policy being renamed.
	SecurityPolicyId pulumi.StringOutput `pulumi:"securityPolicyId"`
	// . New name to be given to the security policy.
	SecurityPolicyName pulumi.StringOutput `pulumi:"securityPolicyName"`
}

// NewAppSecSecurityPolicyRename registers a new resource with the given unique name, arguments, and options.
func NewAppSecSecurityPolicyRename(ctx *pulumi.Context,
	name string, args *AppSecSecurityPolicyRenameArgs, opts ...pulumi.ResourceOption) (*AppSecSecurityPolicyRename, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.SecurityPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyId'")
	}
	if args.SecurityPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'SecurityPolicyName'")
	}
	var resource AppSecSecurityPolicyRename
	err := ctx.RegisterResource("akamai:index/appSecSecurityPolicyRename:AppSecSecurityPolicyRename", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppSecSecurityPolicyRename gets an existing AppSecSecurityPolicyRename resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppSecSecurityPolicyRename(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppSecSecurityPolicyRenameState, opts ...pulumi.ResourceOption) (*AppSecSecurityPolicyRename, error) {
	var resource AppSecSecurityPolicyRename
	err := ctx.ReadResource("akamai:index/appSecSecurityPolicyRename:AppSecSecurityPolicyRename", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppSecSecurityPolicyRename resources.
type appSecSecurityPolicyRenameState struct {
	// . Unique identifier of the security configuration associated with the security policy being renamed.
	ConfigId *int `pulumi:"configId"`
	// . Unique identifier of the security policy being renamed.
	SecurityPolicyId *string `pulumi:"securityPolicyId"`
	// . New name to be given to the security policy.
	SecurityPolicyName *string `pulumi:"securityPolicyName"`
}

type AppSecSecurityPolicyRenameState struct {
	// . Unique identifier of the security configuration associated with the security policy being renamed.
	ConfigId pulumi.IntPtrInput
	// . Unique identifier of the security policy being renamed.
	SecurityPolicyId pulumi.StringPtrInput
	// . New name to be given to the security policy.
	SecurityPolicyName pulumi.StringPtrInput
}

func (AppSecSecurityPolicyRenameState) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSecurityPolicyRenameState)(nil)).Elem()
}

type appSecSecurityPolicyRenameArgs struct {
	// . Unique identifier of the security configuration associated with the security policy being renamed.
	ConfigId int `pulumi:"configId"`
	// . Unique identifier of the security policy being renamed.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	// . New name to be given to the security policy.
	SecurityPolicyName string `pulumi:"securityPolicyName"`
}

// The set of arguments for constructing a AppSecSecurityPolicyRename resource.
type AppSecSecurityPolicyRenameArgs struct {
	// . Unique identifier of the security configuration associated with the security policy being renamed.
	ConfigId pulumi.IntInput
	// . Unique identifier of the security policy being renamed.
	SecurityPolicyId pulumi.StringInput
	// . New name to be given to the security policy.
	SecurityPolicyName pulumi.StringInput
}

func (AppSecSecurityPolicyRenameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appSecSecurityPolicyRenameArgs)(nil)).Elem()
}

type AppSecSecurityPolicyRenameInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyRenameOutput() AppSecSecurityPolicyRenameOutput
	ToAppSecSecurityPolicyRenameOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenameOutput
}

func (*AppSecSecurityPolicyRename) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecSecurityPolicyRename)(nil))
}

func (i *AppSecSecurityPolicyRename) ToAppSecSecurityPolicyRenameOutput() AppSecSecurityPolicyRenameOutput {
	return i.ToAppSecSecurityPolicyRenameOutputWithContext(context.Background())
}

func (i *AppSecSecurityPolicyRename) ToAppSecSecurityPolicyRenameOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyRenameOutput)
}

func (i *AppSecSecurityPolicyRename) ToAppSecSecurityPolicyRenamePtrOutput() AppSecSecurityPolicyRenamePtrOutput {
	return i.ToAppSecSecurityPolicyRenamePtrOutputWithContext(context.Background())
}

func (i *AppSecSecurityPolicyRename) ToAppSecSecurityPolicyRenamePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyRenamePtrOutput)
}

type AppSecSecurityPolicyRenamePtrInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyRenamePtrOutput() AppSecSecurityPolicyRenamePtrOutput
	ToAppSecSecurityPolicyRenamePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenamePtrOutput
}

type appSecSecurityPolicyRenamePtrType AppSecSecurityPolicyRenameArgs

func (*appSecSecurityPolicyRenamePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSecurityPolicyRename)(nil))
}

func (i *appSecSecurityPolicyRenamePtrType) ToAppSecSecurityPolicyRenamePtrOutput() AppSecSecurityPolicyRenamePtrOutput {
	return i.ToAppSecSecurityPolicyRenamePtrOutputWithContext(context.Background())
}

func (i *appSecSecurityPolicyRenamePtrType) ToAppSecSecurityPolicyRenamePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyRenamePtrOutput)
}

// AppSecSecurityPolicyRenameArrayInput is an input type that accepts AppSecSecurityPolicyRenameArray and AppSecSecurityPolicyRenameArrayOutput values.
// You can construct a concrete instance of `AppSecSecurityPolicyRenameArrayInput` via:
//
//          AppSecSecurityPolicyRenameArray{ AppSecSecurityPolicyRenameArgs{...} }
type AppSecSecurityPolicyRenameArrayInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyRenameArrayOutput() AppSecSecurityPolicyRenameArrayOutput
	ToAppSecSecurityPolicyRenameArrayOutputWithContext(context.Context) AppSecSecurityPolicyRenameArrayOutput
}

type AppSecSecurityPolicyRenameArray []AppSecSecurityPolicyRenameInput

func (AppSecSecurityPolicyRenameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppSecSecurityPolicyRename)(nil)).Elem()
}

func (i AppSecSecurityPolicyRenameArray) ToAppSecSecurityPolicyRenameArrayOutput() AppSecSecurityPolicyRenameArrayOutput {
	return i.ToAppSecSecurityPolicyRenameArrayOutputWithContext(context.Background())
}

func (i AppSecSecurityPolicyRenameArray) ToAppSecSecurityPolicyRenameArrayOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyRenameArrayOutput)
}

// AppSecSecurityPolicyRenameMapInput is an input type that accepts AppSecSecurityPolicyRenameMap and AppSecSecurityPolicyRenameMapOutput values.
// You can construct a concrete instance of `AppSecSecurityPolicyRenameMapInput` via:
//
//          AppSecSecurityPolicyRenameMap{ "key": AppSecSecurityPolicyRenameArgs{...} }
type AppSecSecurityPolicyRenameMapInput interface {
	pulumi.Input

	ToAppSecSecurityPolicyRenameMapOutput() AppSecSecurityPolicyRenameMapOutput
	ToAppSecSecurityPolicyRenameMapOutputWithContext(context.Context) AppSecSecurityPolicyRenameMapOutput
}

type AppSecSecurityPolicyRenameMap map[string]AppSecSecurityPolicyRenameInput

func (AppSecSecurityPolicyRenameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppSecSecurityPolicyRename)(nil)).Elem()
}

func (i AppSecSecurityPolicyRenameMap) ToAppSecSecurityPolicyRenameMapOutput() AppSecSecurityPolicyRenameMapOutput {
	return i.ToAppSecSecurityPolicyRenameMapOutputWithContext(context.Background())
}

func (i AppSecSecurityPolicyRenameMap) ToAppSecSecurityPolicyRenameMapOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSecSecurityPolicyRenameMapOutput)
}

type AppSecSecurityPolicyRenameOutput struct{ *pulumi.OutputState }

func (AppSecSecurityPolicyRenameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSecSecurityPolicyRename)(nil))
}

func (o AppSecSecurityPolicyRenameOutput) ToAppSecSecurityPolicyRenameOutput() AppSecSecurityPolicyRenameOutput {
	return o
}

func (o AppSecSecurityPolicyRenameOutput) ToAppSecSecurityPolicyRenameOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenameOutput {
	return o
}

func (o AppSecSecurityPolicyRenameOutput) ToAppSecSecurityPolicyRenamePtrOutput() AppSecSecurityPolicyRenamePtrOutput {
	return o.ToAppSecSecurityPolicyRenamePtrOutputWithContext(context.Background())
}

func (o AppSecSecurityPolicyRenameOutput) ToAppSecSecurityPolicyRenamePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSecSecurityPolicyRename) *AppSecSecurityPolicyRename {
		return &v
	}).(AppSecSecurityPolicyRenamePtrOutput)
}

type AppSecSecurityPolicyRenamePtrOutput struct{ *pulumi.OutputState }

func (AppSecSecurityPolicyRenamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSecSecurityPolicyRename)(nil))
}

func (o AppSecSecurityPolicyRenamePtrOutput) ToAppSecSecurityPolicyRenamePtrOutput() AppSecSecurityPolicyRenamePtrOutput {
	return o
}

func (o AppSecSecurityPolicyRenamePtrOutput) ToAppSecSecurityPolicyRenamePtrOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenamePtrOutput {
	return o
}

func (o AppSecSecurityPolicyRenamePtrOutput) Elem() AppSecSecurityPolicyRenameOutput {
	return o.ApplyT(func(v *AppSecSecurityPolicyRename) AppSecSecurityPolicyRename {
		if v != nil {
			return *v
		}
		var ret AppSecSecurityPolicyRename
		return ret
	}).(AppSecSecurityPolicyRenameOutput)
}

type AppSecSecurityPolicyRenameArrayOutput struct{ *pulumi.OutputState }

func (AppSecSecurityPolicyRenameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppSecSecurityPolicyRename)(nil))
}

func (o AppSecSecurityPolicyRenameArrayOutput) ToAppSecSecurityPolicyRenameArrayOutput() AppSecSecurityPolicyRenameArrayOutput {
	return o
}

func (o AppSecSecurityPolicyRenameArrayOutput) ToAppSecSecurityPolicyRenameArrayOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenameArrayOutput {
	return o
}

func (o AppSecSecurityPolicyRenameArrayOutput) Index(i pulumi.IntInput) AppSecSecurityPolicyRenameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppSecSecurityPolicyRename {
		return vs[0].([]AppSecSecurityPolicyRename)[vs[1].(int)]
	}).(AppSecSecurityPolicyRenameOutput)
}

type AppSecSecurityPolicyRenameMapOutput struct{ *pulumi.OutputState }

func (AppSecSecurityPolicyRenameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppSecSecurityPolicyRename)(nil))
}

func (o AppSecSecurityPolicyRenameMapOutput) ToAppSecSecurityPolicyRenameMapOutput() AppSecSecurityPolicyRenameMapOutput {
	return o
}

func (o AppSecSecurityPolicyRenameMapOutput) ToAppSecSecurityPolicyRenameMapOutputWithContext(ctx context.Context) AppSecSecurityPolicyRenameMapOutput {
	return o
}

func (o AppSecSecurityPolicyRenameMapOutput) MapIndex(k pulumi.StringInput) AppSecSecurityPolicyRenameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppSecSecurityPolicyRename {
		return vs[0].(map[string]AppSecSecurityPolicyRename)[vs[1].(string)]
	}).(AppSecSecurityPolicyRenameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSecurityPolicyRenameInput)(nil)).Elem(), &AppSecSecurityPolicyRename{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSecurityPolicyRenamePtrInput)(nil)).Elem(), &AppSecSecurityPolicyRename{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSecurityPolicyRenameArrayInput)(nil)).Elem(), AppSecSecurityPolicyRenameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppSecSecurityPolicyRenameMapInput)(nil)).Elem(), AppSecSecurityPolicyRenameMap{})
	pulumi.RegisterOutputType(AppSecSecurityPolicyRenameOutput{})
	pulumi.RegisterOutputType(AppSecSecurityPolicyRenamePtrOutput{})
	pulumi.RegisterOutputType(AppSecSecurityPolicyRenameArrayOutput{})
	pulumi.RegisterOutputType(AppSecSecurityPolicyRenameMapOutput{})
}
