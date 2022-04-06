// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package edgedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `DnsRecord` resource to configure a DNS record that can integrate with your existing DNS infrastructure.
//
// ## Example Usage
//
// Here are examples of an A record and a CNAME record.
// ### An A record example
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
// 		_, err := akamai.NewDnsRecord(ctx, "origin", &akamai.DnsRecordArgs{
// 			Active:     pulumi.Bool(true),
// 			Recordtype: pulumi.String("A"),
// 			Targets: pulumi.StringArray{
// 				pulumi.String("192.0.2.42"),
// 			},
// 			Ttl:  pulumi.Int(30),
// 			Zone: pulumi.String("origin.org"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### CNAME Record Example
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
// 		_, err := akamai.NewDnsRecord(ctx, "www", &akamai.DnsRecordArgs{
// 			Active:     pulumi.Bool(true),
// 			Recordtype: pulumi.String("CNAME"),
// 			Targets:    pulumi.StringArray("origin.example.org.edgesuite.net"),
// 			Ttl:        pulumi.Int(600),
// 			Zone:       pulumi.String("example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Argument reference [argument-reference]
//
// This resource supports these arguments for all record types:
//
// * `name` - (Required) The DNS record name. This is the node this DNS record is associated with. Also known as an owner name.
// * `zone` - (Required) The domain zone, including any nested subdomains.
// * `recordType` - (Required) The DNS record type.
// * `ttl` - (Required) The time to live (TTL) is a 32-bit signed integer for the time the resource record is cached. <br /> A value of `0` means that the resource record is not cached. It's only used for the transaction in progress and may be useful for extremely volatile data.
//
// ## Additional arguments by record type
//
// This section lists additional required and optional arguments for specific record types.
//
// ### A record
//
// An A record requires this argument:
//
// * `target` - One or more IPv4 addresses, for example, 192.0.2.0.
//
// ### AAAA record
//
// An AAAA record requires this argument:
//
// * `target` - One or more IPv6 addresses, for example, 2001:0db8::ff00:0042:8329.
//
// ### AFSDB record
//
// An AFSDB record requires these arguments:
//
// * `target` - The domain name of the host having a server for the cell named by the owner name of the resource record.
// * `subtype` - An integer between `0` and `65535` that indicates the type of service provided by the host.
//
// ### AKAMAICDN record
//
// An AKAMAICDN record requires this argument:
//
// * `target` - A DNS name representing the selected edge hostname and domain.
//
// ### AKAMAITLC record
//
// No additional arguments are needed for AKAMAITLC records. This resource returns these computed attributes for this record type:
//
// * `dnsName` - A valid DNS name.
// * `answerType` - The answer type.
//
// ### CAA record
//
// A certificate authority authorization (CAA) record requires this argument:
//
// * `target` - One or more certificate authority authorizations. Each authorization contains three attributes: flags, property tag, and property value.
//
// Example:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### CERT record
//
// A CERT record requires these arguments:
//
// * `typeValue` - A numeric certificate type value.
// * `typeMnemonic` - A mnemonic certificate type value.
// * `keytag` - A value computed for the key embedded in the certificate.
// * `algorithm` - The cryptographic algorithm used to create the signature.
// * `certificate` - Certificate data.
//
// > **Note:** When entering the certificate type, you can enter `typeValue`, `typeMnemonic`, or  both arguments. If you use both, `typeMnemonic` takes precedence.
//
// ### CNAME record
//
// A CNAME record requires this argument:
//
// * ` target  `- A domain name that specifies the canonical or primary name for the owner. The owner name is an alias.
//
// ### DNSKEY record
//
// A DNSKEY record requires these arguments:
//
// * `flags`
// * `protocol` - Set to `3`. If the value isn't `3`, the DNSKEY resource record is treated as invalid during signature verification.
// * `algorithm` - The public key's cryptographic algorithm. This algorithm determines the format of the public key field.
// * `key` - A Base64 encoded value representing the public key. The format used depends on the `algorithm`.
//
// ### DS record
//
// A DS record requires these arguments:
//
// * `keytag` - The key tag of the DNSKEY record that the DS record refers to, in network byte order.
// * `algorithm` - The algorithm number of the DNSKEY resource record referred to by the DS record.
// * `digestType` - Identifies the algorithm used to construct the digest.
// * `digest` - A base 16 encoded DS record includes a digest of the DNSKEY record it refers to. The digest is conifgured the canonical form of the DNSKEY record's fully qualified owner name with the DNSKEY RDATA, and then applying the digest algorithm.
//
// ### HINFO record
//
// A HINFO record requires these arguments:
//
// * `hardware` - The type of hardware the host uses. A machine name or CPU type may be up to 40 characters long and include uppercase letters, digits, hyphens, and slashes. The entry needs to start and to end with an uppercase letter.
// * `software` - The type of software the host uses. A system name may be up to 40 characters long and include uppercase letters, digits, hyphens, and slashes. The entry needs to start with an uppercase letter and end with an uppercase letter or a digit.
//
// ### HTTPS Record
//
// The following fields are required:
//
// * `svcPriority` - Service priority associated with endpoint. Value mist be between 0 and 65535. A piority of 0 enables alias mode.
// * `svcParams` - Space separated list of endpoint parameters. Not allowed if service priority is 0.
// * `targetName` - Domain name of the service endpoint.
//
// ### LOC record
//
// A LOC record requires this argument:
//
// * `target` - A geographical location associated with a domain name.
//
// ### MX record
//
// An MX record supports these arguments:
//
// * `target` - (Required) One or more domain names that specify a host willing to act as a mail exchange for the owner name.
// * `priority` - (Optional) The preference value given to this MX record in relation to all other MX records. When a mailer needs to send mail to a certain DNS domain, it first contacts a DNS server for that domain and retrieves all the MX records. It then contacts the mailer with the lowest preference value. This value is ignored if an embedded priority exists in the target.
// * `priorityIncrement` - (Optional) An auto priority increment when multiple targets are provided with no embedded priority.
//
// See Working with MX records in the DNS Getting Started Guide for more information.
//
// ### NAPTR record
//
// An NAPTR record requires these arguments:
//
// * `order` - A 16-bit unsigned integer specifying the order in which the NAPTR records need to be processed to ensure the correct ordering of rules. Low numbers are processed before high numbers. Once a NAPTR is found whose rule matches the target, the client shouldn't consider any NAPTRs with a higher value for order (except as noted below for the flagsnapter field).
// * `preference` - A 16-bit unsigned integer that specifies the order in which NAPTR records with equal order values are processed. Low numbers are processed before high numbers.
// * `flagsnaptr` - A character string containing flags that control how fields in the record are rewritten and interpreted. Flags are single alphanumeric characters.
// * `service` - Specifies the services available down this rewrite path.
// * `regexp` - A regular expression string containing a substitution expression. This substitution expression is applied to the original client string in order to construct the next domain name to lookup.
// * `replacement` - Depending on the value of the `flags` attribute, the next NAME to query for NAPTR, SRV, or address records. Enter a fully qualified domain name as the value.
//
// ### NS record
//
// An NS record requires these arguments:
//
// * `target` - One or more domain names that specify authoritative hosts for the specified class and domain.
//
// ### NSEC3 record
//
// An NSEC3 record requires these arguments:
//
// * `algorithm` - The cryptographic hash algorithm used to construct the hash-value.
// * `flags` - Eight one-bit flags you can use to indicate different processing. All undefined flags must be zero.
// * `iterations` - The number of additional times the hash function has been performed.
// * `salt` - The base 16 encoded salt value, which is appended to the original owner name before hashing. Used to defend against pre-calculated dictionary attacks.
// * `nextHashedOwnerName` - Base 32 encoded. The next hashed owner name in hash order. This value is in binary format. Given the ordered set of all hashed owner names, the Next Hashed Owner Name field contains the hash of an owner name that immediately follows the owner name of the given NSEC3 RR.
// * `typeBitmaps` - The resource record set types that exist at the original owner name of the NSEC3 RR.
//
// ### NSEC3PARAM record
//
// An NSEC3PARAM record requires these arguments:
//
// * `algorithm` - The cryptographic hash algorithm used to construct the hash-value.
// * `flags` - Eight one-bit flags that can be used to indicate different processing. All undefined flags must be zero.
// * `iterations` - The number of additional times the hash function has been performed.
// * `salt` - The base 16 encoded salt value, which is appended to the original owner name before hashing in order to defend against pre-calculated dictionary attacks.
//
// ### PTR record
//
// A PTR record requires this argument:
//
// * `target` - A domain name that points to some location in the domain name space.
//
// ### RP record
//
// An RP record requires these arguments:
//
// * `mailbox` - A domain name that specifies the mailbox for the responsible person.
// * `txt` - A domain name for which TXT resource records exist.
//
// ### RRSIG record
//
// An RRSIG record requires these arguments:
//
// * `typeCovered` - The resource record set type covered by this signature.
// * `algorithm` - Identifies the cryptographic algorithm used to create the signature.
// * `originalTtl` - The TTL of the covered record set as it appears in the authoritative zone.
// * `expiration` - The end point of this signature's validity. The signature can`t be used for authentication past this point in time.
// * `inception` - The start point of this signature's validity. The signature can`t be used for authentication prior to this point in time.
// * `keytag` - The Key Tag field contains the key tag value of the DNSKEY RR that validates this signature, in network byte order.
// * `signer` - The owner of the DNSKEY resource record who validates this signature.
// * `signature` - The base 64 encoded cryptographic signature that covers the RRSIG RDATA and covered record set. Format depends on the TSIG algorithm in use.
// * `labels` - The Labels field specifies the number of labels in the original RRSIG RR owner name. The significance of this field is that a validator uses it to determine whether the answer was synthesized from a wildcard. If so, it can be used to determine what owner name was used in generating the signature.
//
// ### SPF record
//
// An SPF record requires this argument:
//
// * `target` - Indicates which hosts are, and are not, authorized to use a domain name for the “HELO” and “MAIL FROM” identities.
//
// ### SRV record
//
// An SRV record requires these arguments:
//
// * `target` - The domain name of the target host.
// * `priority` - A 16-bit integer that specifies the preference given to this resource record among others at the same owner. Lower values are preferred.
// * `weight` - A server selection mechanism that specifies a relative weight for entries with the same priority. Larger weights are given a proportionately higher probability of being selected. The range of this number is 0–65535, a 16-bit unsigned integer in network byte order. Domain administrators should use Weight 0 when there isn't any server selection to do, to make the RR easier to read for humans. In the presence of records containing weights greater than 0, records with weight 0 should have a very small chance of being selected.
// * `port` - The port on this target of this service. The range of this number is 0–65535, a 16-bit unsigned integer in network byte order.
//
// ### SSHFP record
//
// An SSHFP record requires these arguments:
//
// * `algorithm` - Describes the algorithm of the public key. The following values are assigned: `0` is reserved, `1` is for RSA, `2` is for DSS, and `3` is for ECDSA.
// * `fingerprintType` - Describes the message-digest algorithm used to calculate the fingerprint of the public key. The following values are assigned: 0 = reserved, 1 = SHA-1, 2 = SHA-256.
// * `fingerprint` - The base 16 encoded fingerprint as calculated over the public key blob. The message-digest algorithm is presumed to produce an opaque octet string output, which is placed as-is in the RDATA fingerprint field.
//
// ### SOA record
//
// An SOA record requires these arguments:
//
// * `nameServer` - The domain name of the name server that was the original or primary source of data for this zone.
// * `emailAddress` - A domain name that specifies the mailbox of this person responsible for this zone.
// * `serial` - The unsigned version number between 0 and 214748364 of the original copy of the zone.
// * `refresh` - A time interval between 0 and 214748364 before the zone should be refreshed.
// * `retry` - A time interval between 0 and 214748364 that should elapse before a failed refresh should be retried.
// * `expiry` - A time value between 0 and 214748364 that specifies the upper limit on the time interval that can elapse before the zone is no longer authoritative.
// * `nxdomainTtl` - The unsigned minimum TTL between 0 and 214748364 that should be exported with any resource record from this zone.
//
// ### SVCB record
//
// An SVCB record requires these arguments:
//
// * `svcPriority` - Service priority associated with endpoint. Value mist be between 0 and 65535. A piority of 0 enables alias mode.
// * `svcParams` - Space separated list of endpoint parameters. Not allowed if service priority is 0.
// * `targetName` - Domain name of the service endpoint.
//
// ### TLSA record
//
// A TLSA record requires these arguments:
//
// * `usage` - Specifies the association used to match the certificate presented in the TLS handshake.
// * `selector` - Specifies the part of the TLS certificate presented by the server that is matched against the association data.
// * `matchType` - Specifies how the certificate association is presented.
// * `certificate` - Specifies the certificate association data to be matched.
//
// ### TXT record
//
// A TXT record requires this argument:
//
// * `target` - One or more character strings. TXT resource records hold descriptive text. The semantics of the text depends on the domain where it is found.
//
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
