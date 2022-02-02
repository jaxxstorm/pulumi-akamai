// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package edgedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: akamai.edgedns.DnsRecord has been deprecated in favor of akamai.DnsRecord
type DnsRecord struct {
	pulumi.CustomResourceState

	Active              pulumi.BoolPtrOutput     `pulumi:"active"`
	Algorithm           pulumi.IntPtrOutput      `pulumi:"algorithm"`
	AnswerType          pulumi.StringOutput      `pulumi:"answerType"`
	Certificate         pulumi.StringPtrOutput   `pulumi:"certificate"`
	Digest              pulumi.StringPtrOutput   `pulumi:"digest"`
	DigestType          pulumi.IntPtrOutput      `pulumi:"digestType"`
	DnsName             pulumi.StringOutput      `pulumi:"dnsName"`
	EmailAddress        pulumi.StringPtrOutput   `pulumi:"emailAddress"`
	Expiration          pulumi.StringPtrOutput   `pulumi:"expiration"`
	Expiry              pulumi.IntPtrOutput      `pulumi:"expiry"`
	Fingerprint         pulumi.StringPtrOutput   `pulumi:"fingerprint"`
	FingerprintType     pulumi.IntPtrOutput      `pulumi:"fingerprintType"`
	Flags               pulumi.IntPtrOutput      `pulumi:"flags"`
	Flagsnaptr          pulumi.StringPtrOutput   `pulumi:"flagsnaptr"`
	Hardware            pulumi.StringPtrOutput   `pulumi:"hardware"`
	Inception           pulumi.StringPtrOutput   `pulumi:"inception"`
	Iterations          pulumi.IntPtrOutput      `pulumi:"iterations"`
	Key                 pulumi.StringPtrOutput   `pulumi:"key"`
	Keytag              pulumi.IntPtrOutput      `pulumi:"keytag"`
	Labels              pulumi.IntPtrOutput      `pulumi:"labels"`
	Mailbox             pulumi.StringPtrOutput   `pulumi:"mailbox"`
	MatchType           pulumi.IntPtrOutput      `pulumi:"matchType"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	NameServer          pulumi.StringPtrOutput   `pulumi:"nameServer"`
	NextHashedOwnerName pulumi.StringPtrOutput   `pulumi:"nextHashedOwnerName"`
	NxdomainTtl         pulumi.IntPtrOutput      `pulumi:"nxdomainTtl"`
	Order               pulumi.IntPtrOutput      `pulumi:"order"`
	OriginalTtl         pulumi.IntPtrOutput      `pulumi:"originalTtl"`
	Port                pulumi.IntPtrOutput      `pulumi:"port"`
	Preference          pulumi.IntPtrOutput      `pulumi:"preference"`
	Priority            pulumi.IntPtrOutput      `pulumi:"priority"`
	PriorityIncrement   pulumi.IntPtrOutput      `pulumi:"priorityIncrement"`
	Protocol            pulumi.IntPtrOutput      `pulumi:"protocol"`
	RecordSha           pulumi.StringOutput      `pulumi:"recordSha"`
	Recordtype          pulumi.StringOutput      `pulumi:"recordtype"`
	Refresh             pulumi.IntPtrOutput      `pulumi:"refresh"`
	Regexp              pulumi.StringPtrOutput   `pulumi:"regexp"`
	Replacement         pulumi.StringPtrOutput   `pulumi:"replacement"`
	Retry               pulumi.IntPtrOutput      `pulumi:"retry"`
	Salt                pulumi.StringPtrOutput   `pulumi:"salt"`
	Selector            pulumi.IntPtrOutput      `pulumi:"selector"`
	Serial              pulumi.IntOutput         `pulumi:"serial"`
	Service             pulumi.StringPtrOutput   `pulumi:"service"`
	Signature           pulumi.StringPtrOutput   `pulumi:"signature"`
	Signer              pulumi.StringPtrOutput   `pulumi:"signer"`
	Software            pulumi.StringPtrOutput   `pulumi:"software"`
	Subtype             pulumi.IntPtrOutput      `pulumi:"subtype"`
	SvcParams           pulumi.StringPtrOutput   `pulumi:"svcParams"`
	SvcPriority         pulumi.IntPtrOutput      `pulumi:"svcPriority"`
	TargetName          pulumi.StringPtrOutput   `pulumi:"targetName"`
	Targets             pulumi.StringArrayOutput `pulumi:"targets"`
	Ttl                 pulumi.IntOutput         `pulumi:"ttl"`
	Txt                 pulumi.StringPtrOutput   `pulumi:"txt"`
	TypeBitmaps         pulumi.StringPtrOutput   `pulumi:"typeBitmaps"`
	TypeCovered         pulumi.StringPtrOutput   `pulumi:"typeCovered"`
	TypeMnemonic        pulumi.StringPtrOutput   `pulumi:"typeMnemonic"`
	TypeValue           pulumi.IntPtrOutput      `pulumi:"typeValue"`
	Usage               pulumi.IntPtrOutput      `pulumi:"usage"`
	Weight              pulumi.IntPtrOutput      `pulumi:"weight"`
	Zone                pulumi.StringOutput      `pulumi:"zone"`
}

// NewDnsRecord registers a new resource with the given unique name, arguments, and options.
func NewDnsRecord(ctx *pulumi.Context,
	name string, args *DnsRecordArgs, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Recordtype == nil {
		return nil, errors.New("invalid value for required argument 'Recordtype'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	var resource DnsRecord
	err := ctx.RegisterResource("akamai:edgedns/dnsRecord:DnsRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsRecord gets an existing DnsRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsRecordState, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	var resource DnsRecord
	err := ctx.ReadResource("akamai:edgedns/dnsRecord:DnsRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsRecord resources.
type dnsRecordState struct {
	Active              *bool    `pulumi:"active"`
	Algorithm           *int     `pulumi:"algorithm"`
	AnswerType          *string  `pulumi:"answerType"`
	Certificate         *string  `pulumi:"certificate"`
	Digest              *string  `pulumi:"digest"`
	DigestType          *int     `pulumi:"digestType"`
	DnsName             *string  `pulumi:"dnsName"`
	EmailAddress        *string  `pulumi:"emailAddress"`
	Expiration          *string  `pulumi:"expiration"`
	Expiry              *int     `pulumi:"expiry"`
	Fingerprint         *string  `pulumi:"fingerprint"`
	FingerprintType     *int     `pulumi:"fingerprintType"`
	Flags               *int     `pulumi:"flags"`
	Flagsnaptr          *string  `pulumi:"flagsnaptr"`
	Hardware            *string  `pulumi:"hardware"`
	Inception           *string  `pulumi:"inception"`
	Iterations          *int     `pulumi:"iterations"`
	Key                 *string  `pulumi:"key"`
	Keytag              *int     `pulumi:"keytag"`
	Labels              *int     `pulumi:"labels"`
	Mailbox             *string  `pulumi:"mailbox"`
	MatchType           *int     `pulumi:"matchType"`
	Name                *string  `pulumi:"name"`
	NameServer          *string  `pulumi:"nameServer"`
	NextHashedOwnerName *string  `pulumi:"nextHashedOwnerName"`
	NxdomainTtl         *int     `pulumi:"nxdomainTtl"`
	Order               *int     `pulumi:"order"`
	OriginalTtl         *int     `pulumi:"originalTtl"`
	Port                *int     `pulumi:"port"`
	Preference          *int     `pulumi:"preference"`
	Priority            *int     `pulumi:"priority"`
	PriorityIncrement   *int     `pulumi:"priorityIncrement"`
	Protocol            *int     `pulumi:"protocol"`
	RecordSha           *string  `pulumi:"recordSha"`
	Recordtype          *string  `pulumi:"recordtype"`
	Refresh             *int     `pulumi:"refresh"`
	Regexp              *string  `pulumi:"regexp"`
	Replacement         *string  `pulumi:"replacement"`
	Retry               *int     `pulumi:"retry"`
	Salt                *string  `pulumi:"salt"`
	Selector            *int     `pulumi:"selector"`
	Serial              *int     `pulumi:"serial"`
	Service             *string  `pulumi:"service"`
	Signature           *string  `pulumi:"signature"`
	Signer              *string  `pulumi:"signer"`
	Software            *string  `pulumi:"software"`
	Subtype             *int     `pulumi:"subtype"`
	SvcParams           *string  `pulumi:"svcParams"`
	SvcPriority         *int     `pulumi:"svcPriority"`
	TargetName          *string  `pulumi:"targetName"`
	Targets             []string `pulumi:"targets"`
	Ttl                 *int     `pulumi:"ttl"`
	Txt                 *string  `pulumi:"txt"`
	TypeBitmaps         *string  `pulumi:"typeBitmaps"`
	TypeCovered         *string  `pulumi:"typeCovered"`
	TypeMnemonic        *string  `pulumi:"typeMnemonic"`
	TypeValue           *int     `pulumi:"typeValue"`
	Usage               *int     `pulumi:"usage"`
	Weight              *int     `pulumi:"weight"`
	Zone                *string  `pulumi:"zone"`
}

type DnsRecordState struct {
	Active              pulumi.BoolPtrInput
	Algorithm           pulumi.IntPtrInput
	AnswerType          pulumi.StringPtrInput
	Certificate         pulumi.StringPtrInput
	Digest              pulumi.StringPtrInput
	DigestType          pulumi.IntPtrInput
	DnsName             pulumi.StringPtrInput
	EmailAddress        pulumi.StringPtrInput
	Expiration          pulumi.StringPtrInput
	Expiry              pulumi.IntPtrInput
	Fingerprint         pulumi.StringPtrInput
	FingerprintType     pulumi.IntPtrInput
	Flags               pulumi.IntPtrInput
	Flagsnaptr          pulumi.StringPtrInput
	Hardware            pulumi.StringPtrInput
	Inception           pulumi.StringPtrInput
	Iterations          pulumi.IntPtrInput
	Key                 pulumi.StringPtrInput
	Keytag              pulumi.IntPtrInput
	Labels              pulumi.IntPtrInput
	Mailbox             pulumi.StringPtrInput
	MatchType           pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	NameServer          pulumi.StringPtrInput
	NextHashedOwnerName pulumi.StringPtrInput
	NxdomainTtl         pulumi.IntPtrInput
	Order               pulumi.IntPtrInput
	OriginalTtl         pulumi.IntPtrInput
	Port                pulumi.IntPtrInput
	Preference          pulumi.IntPtrInput
	Priority            pulumi.IntPtrInput
	PriorityIncrement   pulumi.IntPtrInput
	Protocol            pulumi.IntPtrInput
	RecordSha           pulumi.StringPtrInput
	Recordtype          pulumi.StringPtrInput
	Refresh             pulumi.IntPtrInput
	Regexp              pulumi.StringPtrInput
	Replacement         pulumi.StringPtrInput
	Retry               pulumi.IntPtrInput
	Salt                pulumi.StringPtrInput
	Selector            pulumi.IntPtrInput
	Serial              pulumi.IntPtrInput
	Service             pulumi.StringPtrInput
	Signature           pulumi.StringPtrInput
	Signer              pulumi.StringPtrInput
	Software            pulumi.StringPtrInput
	Subtype             pulumi.IntPtrInput
	SvcParams           pulumi.StringPtrInput
	SvcPriority         pulumi.IntPtrInput
	TargetName          pulumi.StringPtrInput
	Targets             pulumi.StringArrayInput
	Ttl                 pulumi.IntPtrInput
	Txt                 pulumi.StringPtrInput
	TypeBitmaps         pulumi.StringPtrInput
	TypeCovered         pulumi.StringPtrInput
	TypeMnemonic        pulumi.StringPtrInput
	TypeValue           pulumi.IntPtrInput
	Usage               pulumi.IntPtrInput
	Weight              pulumi.IntPtrInput
	Zone                pulumi.StringPtrInput
}

func (DnsRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordState)(nil)).Elem()
}

type dnsRecordArgs struct {
	Active              *bool    `pulumi:"active"`
	Algorithm           *int     `pulumi:"algorithm"`
	Certificate         *string  `pulumi:"certificate"`
	Digest              *string  `pulumi:"digest"`
	DigestType          *int     `pulumi:"digestType"`
	EmailAddress        *string  `pulumi:"emailAddress"`
	Expiration          *string  `pulumi:"expiration"`
	Expiry              *int     `pulumi:"expiry"`
	Fingerprint         *string  `pulumi:"fingerprint"`
	FingerprintType     *int     `pulumi:"fingerprintType"`
	Flags               *int     `pulumi:"flags"`
	Flagsnaptr          *string  `pulumi:"flagsnaptr"`
	Hardware            *string  `pulumi:"hardware"`
	Inception           *string  `pulumi:"inception"`
	Iterations          *int     `pulumi:"iterations"`
	Key                 *string  `pulumi:"key"`
	Keytag              *int     `pulumi:"keytag"`
	Labels              *int     `pulumi:"labels"`
	Mailbox             *string  `pulumi:"mailbox"`
	MatchType           *int     `pulumi:"matchType"`
	Name                *string  `pulumi:"name"`
	NameServer          *string  `pulumi:"nameServer"`
	NextHashedOwnerName *string  `pulumi:"nextHashedOwnerName"`
	NxdomainTtl         *int     `pulumi:"nxdomainTtl"`
	Order               *int     `pulumi:"order"`
	OriginalTtl         *int     `pulumi:"originalTtl"`
	Port                *int     `pulumi:"port"`
	Preference          *int     `pulumi:"preference"`
	Priority            *int     `pulumi:"priority"`
	PriorityIncrement   *int     `pulumi:"priorityIncrement"`
	Protocol            *int     `pulumi:"protocol"`
	Recordtype          string   `pulumi:"recordtype"`
	Refresh             *int     `pulumi:"refresh"`
	Regexp              *string  `pulumi:"regexp"`
	Replacement         *string  `pulumi:"replacement"`
	Retry               *int     `pulumi:"retry"`
	Salt                *string  `pulumi:"salt"`
	Selector            *int     `pulumi:"selector"`
	Service             *string  `pulumi:"service"`
	Signature           *string  `pulumi:"signature"`
	Signer              *string  `pulumi:"signer"`
	Software            *string  `pulumi:"software"`
	Subtype             *int     `pulumi:"subtype"`
	SvcParams           *string  `pulumi:"svcParams"`
	SvcPriority         *int     `pulumi:"svcPriority"`
	TargetName          *string  `pulumi:"targetName"`
	Targets             []string `pulumi:"targets"`
	Ttl                 int      `pulumi:"ttl"`
	Txt                 *string  `pulumi:"txt"`
	TypeBitmaps         *string  `pulumi:"typeBitmaps"`
	TypeCovered         *string  `pulumi:"typeCovered"`
	TypeMnemonic        *string  `pulumi:"typeMnemonic"`
	TypeValue           *int     `pulumi:"typeValue"`
	Usage               *int     `pulumi:"usage"`
	Weight              *int     `pulumi:"weight"`
	Zone                string   `pulumi:"zone"`
}

// The set of arguments for constructing a DnsRecord resource.
type DnsRecordArgs struct {
	Active              pulumi.BoolPtrInput
	Algorithm           pulumi.IntPtrInput
	Certificate         pulumi.StringPtrInput
	Digest              pulumi.StringPtrInput
	DigestType          pulumi.IntPtrInput
	EmailAddress        pulumi.StringPtrInput
	Expiration          pulumi.StringPtrInput
	Expiry              pulumi.IntPtrInput
	Fingerprint         pulumi.StringPtrInput
	FingerprintType     pulumi.IntPtrInput
	Flags               pulumi.IntPtrInput
	Flagsnaptr          pulumi.StringPtrInput
	Hardware            pulumi.StringPtrInput
	Inception           pulumi.StringPtrInput
	Iterations          pulumi.IntPtrInput
	Key                 pulumi.StringPtrInput
	Keytag              pulumi.IntPtrInput
	Labels              pulumi.IntPtrInput
	Mailbox             pulumi.StringPtrInput
	MatchType           pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	NameServer          pulumi.StringPtrInput
	NextHashedOwnerName pulumi.StringPtrInput
	NxdomainTtl         pulumi.IntPtrInput
	Order               pulumi.IntPtrInput
	OriginalTtl         pulumi.IntPtrInput
	Port                pulumi.IntPtrInput
	Preference          pulumi.IntPtrInput
	Priority            pulumi.IntPtrInput
	PriorityIncrement   pulumi.IntPtrInput
	Protocol            pulumi.IntPtrInput
	Recordtype          pulumi.StringInput
	Refresh             pulumi.IntPtrInput
	Regexp              pulumi.StringPtrInput
	Replacement         pulumi.StringPtrInput
	Retry               pulumi.IntPtrInput
	Salt                pulumi.StringPtrInput
	Selector            pulumi.IntPtrInput
	Service             pulumi.StringPtrInput
	Signature           pulumi.StringPtrInput
	Signer              pulumi.StringPtrInput
	Software            pulumi.StringPtrInput
	Subtype             pulumi.IntPtrInput
	SvcParams           pulumi.StringPtrInput
	SvcPriority         pulumi.IntPtrInput
	TargetName          pulumi.StringPtrInput
	Targets             pulumi.StringArrayInput
	Ttl                 pulumi.IntInput
	Txt                 pulumi.StringPtrInput
	TypeBitmaps         pulumi.StringPtrInput
	TypeCovered         pulumi.StringPtrInput
	TypeMnemonic        pulumi.StringPtrInput
	TypeValue           pulumi.IntPtrInput
	Usage               pulumi.IntPtrInput
	Weight              pulumi.IntPtrInput
	Zone                pulumi.StringInput
}

func (DnsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordArgs)(nil)).Elem()
}

type DnsRecordInput interface {
	pulumi.Input

	ToDnsRecordOutput() DnsRecordOutput
	ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput
}

func (*DnsRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (i *DnsRecord) ToDnsRecordOutput() DnsRecordOutput {
	return i.ToDnsRecordOutputWithContext(context.Background())
}

func (i *DnsRecord) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordOutput)
}

// DnsRecordArrayInput is an input type that accepts DnsRecordArray and DnsRecordArrayOutput values.
// You can construct a concrete instance of `DnsRecordArrayInput` via:
//
//          DnsRecordArray{ DnsRecordArgs{...} }
type DnsRecordArrayInput interface {
	pulumi.Input

	ToDnsRecordArrayOutput() DnsRecordArrayOutput
	ToDnsRecordArrayOutputWithContext(context.Context) DnsRecordArrayOutput
}

type DnsRecordArray []DnsRecordInput

func (DnsRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordArray) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return i.ToDnsRecordArrayOutputWithContext(context.Background())
}

func (i DnsRecordArray) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordArrayOutput)
}

// DnsRecordMapInput is an input type that accepts DnsRecordMap and DnsRecordMapOutput values.
// You can construct a concrete instance of `DnsRecordMapInput` via:
//
//          DnsRecordMap{ "key": DnsRecordArgs{...} }
type DnsRecordMapInput interface {
	pulumi.Input

	ToDnsRecordMapOutput() DnsRecordMapOutput
	ToDnsRecordMapOutputWithContext(context.Context) DnsRecordMapOutput
}

type DnsRecordMap map[string]DnsRecordInput

func (DnsRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordMap) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return i.ToDnsRecordMapOutputWithContext(context.Background())
}

func (i DnsRecordMap) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordMapOutput)
}

type DnsRecordOutput struct{ *pulumi.OutputState }

func (DnsRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (o DnsRecordOutput) ToDnsRecordOutput() DnsRecordOutput {
	return o
}

func (o DnsRecordOutput) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return o
}

type DnsRecordArrayOutput struct{ *pulumi.OutputState }

func (DnsRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) Index(i pulumi.IntInput) DnsRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].([]*DnsRecord)[vs[1].(int)]
	}).(DnsRecordOutput)
}

type DnsRecordMapOutput struct{ *pulumi.OutputState }

func (DnsRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) MapIndex(k pulumi.StringInput) DnsRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].(map[string]*DnsRecord)[vs[1].(string)]
	}).(DnsRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordInput)(nil)).Elem(), &DnsRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordArrayInput)(nil)).Elem(), DnsRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordMapInput)(nil)).Elem(), DnsRecordMap{})
	pulumi.RegisterOutputType(DnsRecordOutput{})
	pulumi.RegisterOutputType(DnsRecordArrayOutput{})
	pulumi.RegisterOutputType(DnsRecordMapOutput{})
}
