// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `IamUser` resource represents a user on the Akamai platform.
//
// ## Basic usage
//
// This example shows how to set up a user:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-akamai/sdk/v3/go/akamai"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := akamai.NewIamUser(ctx, "exampleUser", &akamai.IamUserArgs{
// 			AuthGrantsJson: pulumi.String("[{\"groupId\":18451,\"roleId\":14},{\"groupId\":18453,\"roleId\":13}]"),
// 			Country:        pulumi.String("Grenada"),
// 			Email:          pulumi.String("jperez@example.com"),
// 			EnableTfa:      pulumi.Bool(false),
// 			FirstName:      pulumi.String("Juan"),
// 			LastName:       pulumi.String("Perez"),
// 			Phone:          pulumi.String("+1 206-555-0100"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Attributes reference
//
// This resource returns these attributes:
//
// * `sessionTimeout` - The number of seconds it takes for the user's session to time out if there hasn't been any activity.
// * `userName` - A user's `loginId`. Typically, a user's email address.
// * `isLocked` - The user's lock status.
// * `lastLogin` - ISO 8601 timestamp indicating when the user last logged in.
// * `passwordExpiredAfter` - The date a user's password expires.
// * `tfaConfigured` - Indicates whether two-factor authentication is configured.
// * `emailUpdatePending` - Indicates whether email update is pending.
// * `lock` - (Optional) Flag to block as user account.
type IamUser struct {
	pulumi.CustomResourceState

	// The user's street address.
	Address pulumi.StringOutput `pulumi:"address"`
	// A user's per-group role assignments, in JSON form.
	AuthGrantsJson pulumi.StringOutput `pulumi:"authGrantsJson"`
	// The user's city.
	City pulumi.StringPtrOutput `pulumi:"city"`
	// To help characterize the user, the value can be any that are available from the view-contact-types operation.
	ContactType pulumi.StringOutput `pulumi:"contactType"`
	// As part of the user's location, the value can be any that are available from the view-supported-countries operation.
	Country pulumi.StringOutput `pulumi:"country"`
	// The user's email address.
	Email pulumi.StringOutput `pulumi:"email"`
	// Indicates whether email update is pending
	EmailUpdatePending pulumi.BoolOutput `pulumi:"emailUpdatePending"`
	// Indicates whether two-factor authentication is allowed.
	EnableTfa pulumi.BoolOutput `pulumi:"enableTfa"`
	// The user's first name.
	FirstName pulumi.StringOutput `pulumi:"firstName"`
	// The user's lock status.
	//
	// Deprecated: The setting "is_locked" has been deprecated. Please use "lock" setting instead
	IsLocked pulumi.BoolOutput `pulumi:"isLocked"`
	// The user's position at your company
	JobTitle pulumi.StringPtrOutput `pulumi:"jobTitle"`
	// ISO 8601 timestamp indicating when the user last logged in
	LastLogin pulumi.StringOutput `pulumi:"lastLogin"`
	// The user's last name.
	LastName pulumi.StringOutput `pulumi:"lastName"`
	// Flag to block a user account
	Lock pulumi.BoolPtrOutput `pulumi:"lock"`
	// The user's mobile phone number.
	MobilePhone pulumi.StringPtrOutput `pulumi:"mobilePhone"`
	// The date a user's password expires
	PasswordExpiredAfter pulumi.StringOutput `pulumi:"passwordExpiredAfter"`
	// The user's main phone number.
	Phone pulumi.StringOutput `pulumi:"phone"`
	// The value can be any that are available from the view-languages operation
	PreferredLanguage pulumi.StringOutput `pulumi:"preferredLanguage"`
	// The user's secondary email address.
	SecondaryEmail pulumi.StringPtrOutput `pulumi:"secondaryEmail"`
	// The number of seconds it takes for the user's Control Center session to time out if there hasn't been any activity
	SessionTimeout pulumi.IntOutput `pulumi:"sessionTimeout"`
	// The user's state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Indicates whether two-factor authentication is configured
	TfaConfigured pulumi.BoolOutput `pulumi:"tfaConfigured"`
	// The user's time zone. The value can be any that are available from the view-time-zones operation
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// A user's `loginId`. Typically, a user's email address
	UserName pulumi.StringOutput `pulumi:"userName"`
	// The user's five-digit ZIP code.
	ZipCode pulumi.StringPtrOutput `pulumi:"zipCode"`
}

// NewIamUser registers a new resource with the given unique name, arguments, and options.
func NewIamUser(ctx *pulumi.Context,
	name string, args *IamUserArgs, opts ...pulumi.ResourceOption) (*IamUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthGrantsJson == nil {
		return nil, errors.New("invalid value for required argument 'AuthGrantsJson'")
	}
	if args.Country == nil {
		return nil, errors.New("invalid value for required argument 'Country'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.EnableTfa == nil {
		return nil, errors.New("invalid value for required argument 'EnableTfa'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.Phone == nil {
		return nil, errors.New("invalid value for required argument 'Phone'")
	}
	var resource IamUser
	err := ctx.RegisterResource("akamai:index/iamUser:IamUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamUser gets an existing IamUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamUserState, opts ...pulumi.ResourceOption) (*IamUser, error) {
	var resource IamUser
	err := ctx.ReadResource("akamai:index/iamUser:IamUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamUser resources.
type iamUserState struct {
	// The user's street address.
	Address *string `pulumi:"address"`
	// A user's per-group role assignments, in JSON form.
	AuthGrantsJson *string `pulumi:"authGrantsJson"`
	// The user's city.
	City *string `pulumi:"city"`
	// To help characterize the user, the value can be any that are available from the view-contact-types operation.
	ContactType *string `pulumi:"contactType"`
	// As part of the user's location, the value can be any that are available from the view-supported-countries operation.
	Country *string `pulumi:"country"`
	// The user's email address.
	Email *string `pulumi:"email"`
	// Indicates whether email update is pending
	EmailUpdatePending *bool `pulumi:"emailUpdatePending"`
	// Indicates whether two-factor authentication is allowed.
	EnableTfa *bool `pulumi:"enableTfa"`
	// The user's first name.
	FirstName *string `pulumi:"firstName"`
	// The user's lock status.
	//
	// Deprecated: The setting "is_locked" has been deprecated. Please use "lock" setting instead
	IsLocked *bool `pulumi:"isLocked"`
	// The user's position at your company
	JobTitle *string `pulumi:"jobTitle"`
	// ISO 8601 timestamp indicating when the user last logged in
	LastLogin *string `pulumi:"lastLogin"`
	// The user's last name.
	LastName *string `pulumi:"lastName"`
	// Flag to block a user account
	Lock *bool `pulumi:"lock"`
	// The user's mobile phone number.
	MobilePhone *string `pulumi:"mobilePhone"`
	// The date a user's password expires
	PasswordExpiredAfter *string `pulumi:"passwordExpiredAfter"`
	// The user's main phone number.
	Phone *string `pulumi:"phone"`
	// The value can be any that are available from the view-languages operation
	PreferredLanguage *string `pulumi:"preferredLanguage"`
	// The user's secondary email address.
	SecondaryEmail *string `pulumi:"secondaryEmail"`
	// The number of seconds it takes for the user's Control Center session to time out if there hasn't been any activity
	SessionTimeout *int `pulumi:"sessionTimeout"`
	// The user's state.
	State *string `pulumi:"state"`
	// Indicates whether two-factor authentication is configured
	TfaConfigured *bool `pulumi:"tfaConfigured"`
	// The user's time zone. The value can be any that are available from the view-time-zones operation
	TimeZone *string `pulumi:"timeZone"`
	// A user's `loginId`. Typically, a user's email address
	UserName *string `pulumi:"userName"`
	// The user's five-digit ZIP code.
	ZipCode *string `pulumi:"zipCode"`
}

type IamUserState struct {
	// The user's street address.
	Address pulumi.StringPtrInput
	// A user's per-group role assignments, in JSON form.
	AuthGrantsJson pulumi.StringPtrInput
	// The user's city.
	City pulumi.StringPtrInput
	// To help characterize the user, the value can be any that are available from the view-contact-types operation.
	ContactType pulumi.StringPtrInput
	// As part of the user's location, the value can be any that are available from the view-supported-countries operation.
	Country pulumi.StringPtrInput
	// The user's email address.
	Email pulumi.StringPtrInput
	// Indicates whether email update is pending
	EmailUpdatePending pulumi.BoolPtrInput
	// Indicates whether two-factor authentication is allowed.
	EnableTfa pulumi.BoolPtrInput
	// The user's first name.
	FirstName pulumi.StringPtrInput
	// The user's lock status.
	//
	// Deprecated: The setting "is_locked" has been deprecated. Please use "lock" setting instead
	IsLocked pulumi.BoolPtrInput
	// The user's position at your company
	JobTitle pulumi.StringPtrInput
	// ISO 8601 timestamp indicating when the user last logged in
	LastLogin pulumi.StringPtrInput
	// The user's last name.
	LastName pulumi.StringPtrInput
	// Flag to block a user account
	Lock pulumi.BoolPtrInput
	// The user's mobile phone number.
	MobilePhone pulumi.StringPtrInput
	// The date a user's password expires
	PasswordExpiredAfter pulumi.StringPtrInput
	// The user's main phone number.
	Phone pulumi.StringPtrInput
	// The value can be any that are available from the view-languages operation
	PreferredLanguage pulumi.StringPtrInput
	// The user's secondary email address.
	SecondaryEmail pulumi.StringPtrInput
	// The number of seconds it takes for the user's Control Center session to time out if there hasn't been any activity
	SessionTimeout pulumi.IntPtrInput
	// The user's state.
	State pulumi.StringPtrInput
	// Indicates whether two-factor authentication is configured
	TfaConfigured pulumi.BoolPtrInput
	// The user's time zone. The value can be any that are available from the view-time-zones operation
	TimeZone pulumi.StringPtrInput
	// A user's `loginId`. Typically, a user's email address
	UserName pulumi.StringPtrInput
	// The user's five-digit ZIP code.
	ZipCode pulumi.StringPtrInput
}

func (IamUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserState)(nil)).Elem()
}

type iamUserArgs struct {
	// The user's street address.
	Address *string `pulumi:"address"`
	// A user's per-group role assignments, in JSON form.
	AuthGrantsJson string `pulumi:"authGrantsJson"`
	// The user's city.
	City *string `pulumi:"city"`
	// To help characterize the user, the value can be any that are available from the view-contact-types operation.
	ContactType *string `pulumi:"contactType"`
	// As part of the user's location, the value can be any that are available from the view-supported-countries operation.
	Country string `pulumi:"country"`
	// The user's email address.
	Email string `pulumi:"email"`
	// Indicates whether two-factor authentication is allowed.
	EnableTfa bool `pulumi:"enableTfa"`
	// The user's first name.
	FirstName string `pulumi:"firstName"`
	// The user's position at your company
	JobTitle *string `pulumi:"jobTitle"`
	// The user's last name.
	LastName string `pulumi:"lastName"`
	// Flag to block a user account
	Lock *bool `pulumi:"lock"`
	// The user's mobile phone number.
	MobilePhone *string `pulumi:"mobilePhone"`
	// The user's main phone number.
	Phone string `pulumi:"phone"`
	// The value can be any that are available from the view-languages operation
	PreferredLanguage *string `pulumi:"preferredLanguage"`
	// The user's secondary email address.
	SecondaryEmail *string `pulumi:"secondaryEmail"`
	// The number of seconds it takes for the user's Control Center session to time out if there hasn't been any activity
	SessionTimeout *int `pulumi:"sessionTimeout"`
	// The user's state.
	State *string `pulumi:"state"`
	// The user's time zone. The value can be any that are available from the view-time-zones operation
	TimeZone *string `pulumi:"timeZone"`
	// The user's five-digit ZIP code.
	ZipCode *string `pulumi:"zipCode"`
}

// The set of arguments for constructing a IamUser resource.
type IamUserArgs struct {
	// The user's street address.
	Address pulumi.StringPtrInput
	// A user's per-group role assignments, in JSON form.
	AuthGrantsJson pulumi.StringInput
	// The user's city.
	City pulumi.StringPtrInput
	// To help characterize the user, the value can be any that are available from the view-contact-types operation.
	ContactType pulumi.StringPtrInput
	// As part of the user's location, the value can be any that are available from the view-supported-countries operation.
	Country pulumi.StringInput
	// The user's email address.
	Email pulumi.StringInput
	// Indicates whether two-factor authentication is allowed.
	EnableTfa pulumi.BoolInput
	// The user's first name.
	FirstName pulumi.StringInput
	// The user's position at your company
	JobTitle pulumi.StringPtrInput
	// The user's last name.
	LastName pulumi.StringInput
	// Flag to block a user account
	Lock pulumi.BoolPtrInput
	// The user's mobile phone number.
	MobilePhone pulumi.StringPtrInput
	// The user's main phone number.
	Phone pulumi.StringInput
	// The value can be any that are available from the view-languages operation
	PreferredLanguage pulumi.StringPtrInput
	// The user's secondary email address.
	SecondaryEmail pulumi.StringPtrInput
	// The number of seconds it takes for the user's Control Center session to time out if there hasn't been any activity
	SessionTimeout pulumi.IntPtrInput
	// The user's state.
	State pulumi.StringPtrInput
	// The user's time zone. The value can be any that are available from the view-time-zones operation
	TimeZone pulumi.StringPtrInput
	// The user's five-digit ZIP code.
	ZipCode pulumi.StringPtrInput
}

func (IamUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserArgs)(nil)).Elem()
}

type IamUserInput interface {
	pulumi.Input

	ToIamUserOutput() IamUserOutput
	ToIamUserOutputWithContext(ctx context.Context) IamUserOutput
}

func (*IamUser) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUser)(nil)).Elem()
}

func (i *IamUser) ToIamUserOutput() IamUserOutput {
	return i.ToIamUserOutputWithContext(context.Background())
}

func (i *IamUser) ToIamUserOutputWithContext(ctx context.Context) IamUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserOutput)
}

// IamUserArrayInput is an input type that accepts IamUserArray and IamUserArrayOutput values.
// You can construct a concrete instance of `IamUserArrayInput` via:
//
//          IamUserArray{ IamUserArgs{...} }
type IamUserArrayInput interface {
	pulumi.Input

	ToIamUserArrayOutput() IamUserArrayOutput
	ToIamUserArrayOutputWithContext(context.Context) IamUserArrayOutput
}

type IamUserArray []IamUserInput

func (IamUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamUser)(nil)).Elem()
}

func (i IamUserArray) ToIamUserArrayOutput() IamUserArrayOutput {
	return i.ToIamUserArrayOutputWithContext(context.Background())
}

func (i IamUserArray) ToIamUserArrayOutputWithContext(ctx context.Context) IamUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserArrayOutput)
}

// IamUserMapInput is an input type that accepts IamUserMap and IamUserMapOutput values.
// You can construct a concrete instance of `IamUserMapInput` via:
//
//          IamUserMap{ "key": IamUserArgs{...} }
type IamUserMapInput interface {
	pulumi.Input

	ToIamUserMapOutput() IamUserMapOutput
	ToIamUserMapOutputWithContext(context.Context) IamUserMapOutput
}

type IamUserMap map[string]IamUserInput

func (IamUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamUser)(nil)).Elem()
}

func (i IamUserMap) ToIamUserMapOutput() IamUserMapOutput {
	return i.ToIamUserMapOutputWithContext(context.Background())
}

func (i IamUserMap) ToIamUserMapOutputWithContext(ctx context.Context) IamUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserMapOutput)
}

type IamUserOutput struct{ *pulumi.OutputState }

func (IamUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUser)(nil)).Elem()
}

func (o IamUserOutput) ToIamUserOutput() IamUserOutput {
	return o
}

func (o IamUserOutput) ToIamUserOutputWithContext(ctx context.Context) IamUserOutput {
	return o
}

// The user's street address.
func (o IamUserOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// A user's per-group role assignments, in JSON form.
func (o IamUserOutput) AuthGrantsJson() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.AuthGrantsJson }).(pulumi.StringOutput)
}

// The user's city.
func (o IamUserOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringPtrOutput { return v.City }).(pulumi.StringPtrOutput)
}

// To help characterize the user, the value can be any that are available from the view-contact-types operation.
func (o IamUserOutput) ContactType() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.ContactType }).(pulumi.StringOutput)
}

// As part of the user's location, the value can be any that are available from the view-supported-countries operation.
func (o IamUserOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// The user's email address.
func (o IamUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Indicates whether email update is pending
func (o IamUserOutput) EmailUpdatePending() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamUser) pulumi.BoolOutput { return v.EmailUpdatePending }).(pulumi.BoolOutput)
}

// Indicates whether two-factor authentication is allowed.
func (o IamUserOutput) EnableTfa() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamUser) pulumi.BoolOutput { return v.EnableTfa }).(pulumi.BoolOutput)
}

// The user's first name.
func (o IamUserOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.FirstName }).(pulumi.StringOutput)
}

// The user's lock status.
//
// Deprecated: The setting "is_locked" has been deprecated. Please use "lock" setting instead
func (o IamUserOutput) IsLocked() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamUser) pulumi.BoolOutput { return v.IsLocked }).(pulumi.BoolOutput)
}

// The user's position at your company
func (o IamUserOutput) JobTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringPtrOutput { return v.JobTitle }).(pulumi.StringPtrOutput)
}

// ISO 8601 timestamp indicating when the user last logged in
func (o IamUserOutput) LastLogin() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.LastLogin }).(pulumi.StringOutput)
}

// The user's last name.
func (o IamUserOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.LastName }).(pulumi.StringOutput)
}

// Flag to block a user account
func (o IamUserOutput) Lock() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IamUser) pulumi.BoolPtrOutput { return v.Lock }).(pulumi.BoolPtrOutput)
}

// The user's mobile phone number.
func (o IamUserOutput) MobilePhone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringPtrOutput { return v.MobilePhone }).(pulumi.StringPtrOutput)
}

// The date a user's password expires
func (o IamUserOutput) PasswordExpiredAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.PasswordExpiredAfter }).(pulumi.StringOutput)
}

// The user's main phone number.
func (o IamUserOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.Phone }).(pulumi.StringOutput)
}

// The value can be any that are available from the view-languages operation
func (o IamUserOutput) PreferredLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.PreferredLanguage }).(pulumi.StringOutput)
}

// The user's secondary email address.
func (o IamUserOutput) SecondaryEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringPtrOutput { return v.SecondaryEmail }).(pulumi.StringPtrOutput)
}

// The number of seconds it takes for the user's Control Center session to time out if there hasn't been any activity
func (o IamUserOutput) SessionTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *IamUser) pulumi.IntOutput { return v.SessionTimeout }).(pulumi.IntOutput)
}

// The user's state.
func (o IamUserOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Indicates whether two-factor authentication is configured
func (o IamUserOutput) TfaConfigured() pulumi.BoolOutput {
	return o.ApplyT(func(v *IamUser) pulumi.BoolOutput { return v.TfaConfigured }).(pulumi.BoolOutput)
}

// The user's time zone. The value can be any that are available from the view-time-zones operation
func (o IamUserOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// A user's `loginId`. Typically, a user's email address
func (o IamUserOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

// The user's five-digit ZIP code.
func (o IamUserOutput) ZipCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamUser) pulumi.StringPtrOutput { return v.ZipCode }).(pulumi.StringPtrOutput)
}

type IamUserArrayOutput struct{ *pulumi.OutputState }

func (IamUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamUser)(nil)).Elem()
}

func (o IamUserArrayOutput) ToIamUserArrayOutput() IamUserArrayOutput {
	return o
}

func (o IamUserArrayOutput) ToIamUserArrayOutputWithContext(ctx context.Context) IamUserArrayOutput {
	return o
}

func (o IamUserArrayOutput) Index(i pulumi.IntInput) IamUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamUser {
		return vs[0].([]*IamUser)[vs[1].(int)]
	}).(IamUserOutput)
}

type IamUserMapOutput struct{ *pulumi.OutputState }

func (IamUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamUser)(nil)).Elem()
}

func (o IamUserMapOutput) ToIamUserMapOutput() IamUserMapOutput {
	return o
}

func (o IamUserMapOutput) ToIamUserMapOutputWithContext(ctx context.Context) IamUserMapOutput {
	return o
}

func (o IamUserMapOutput) MapIndex(k pulumi.StringInput) IamUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamUser {
		return vs[0].(map[string]*IamUser)[vs[1].(string)]
	}).(IamUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserInput)(nil)).Elem(), &IamUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserArrayInput)(nil)).Elem(), IamUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserMapInput)(nil)).Elem(), IamUserMap{})
	pulumi.RegisterOutputType(IamUserOutput{})
	pulumi.RegisterOutputType(IamUserArrayOutput{})
	pulumi.RegisterOutputType(IamUserMapOutput{})
}
