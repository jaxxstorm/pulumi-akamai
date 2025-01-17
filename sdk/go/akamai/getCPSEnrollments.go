// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `getCPSEnrollments` data source to return data for all of a specific contract's enrollments.
//
// ## Basic usage
//
// This example shows how to set up a user:
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
//			testEnrollmentsList, err := akamai.GetCPSEnrollments(ctx, &GetCPSEnrollmentsArgs{
//				ContractId: _var.Contract_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("dvOutput", testEnrollmentsList)
//			return nil
//		})
//	}
//
// ```
//
// ## Attributes reference
//
// This data source returns these attributes:
//
// * `enrollments`
//   - `enrollmentId`
//   - `commonName` - The fully qualified domain name (FQDN) used for the certificate.
//   - `sans` - Additional common names in a Subject Alternative Names (SAN) list.
//   - `secureNetwork` - The type of deployment network used. `standard-tls` deploys your certificate to Akamai's standard secure network, but it isn't PCI compliant. `enhanced-tls` deploys your certificate to Akamai's more secure network with PCI compliance capability.
//   - `sniOnly` - Whether you enabled SNI-only extension for the enrollment. Server Name Indication (SNI) is an extension of the Transport Layer Security (TLS) networking protocol. It allows a server to present multiple certificates on the same IP address. All modern web browsers support the SNI extension. If you have the same SAN on two or more certificates with the SNI-only option set, Akamai may serve traffic using any certificate which matches the requested SNI hostname.
//   - `adminContact` - Contact information for the certificate administrator at your company.
//   - `certificateChainType` - Certificate trust chain type.
//   - `csr` - When you create an enrollment, you also generate a certificate signing request (CSR) using CPS. CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
//   - `countryCode` - The country code for the country where your organization is located.
//   - `city` - The city where your organization resides.
//   - `organization` - The name of your company or organization.
//   - `organizationalUnit` - Your organizational unit.
//   - `state` - Your state or province.
//   - `enableMultiStackedCertificates` - If present, an ECDSA certificate is enabled in addition to an RSA certificate. CPS automatically performs all certificate operations on both certificates, and uses the best certificate for each client connection to your secure properties.
//   - `networkConfiguration` - The network information and TLS Metadata you want CPS to use to push the completed certificate to the network.
//   - `clientMutualAuthentication` - If present, shows the configuration for client mutual authentication. Specifies the trust chain that is used to verify client certificates and some configuration options.
//   - `sendCaListToClient` - If present, the server is enabled to send the certificate authority (CA) list to the client.
//   - `ocspEnabled` - If present, the Online Certificate Status Protocol (OCSP) stapling is enabled for client certificates.
//   - `setId` - The identifier of the set of trust chains, created in [Trust Chain Manager](https://techdocs.akamai.com/trust-chain-mgr/docs/welcome-trust-chain-manager).
//   - `disallowedTlsVersions` - The TLS protocol version that is not trusted. CPS uses the TLS protocols that Akamai currently supports as a best practice.
//   - `cloneDnsNames` - If present, CPS directs traffic using all the SANs listed in the SANs parameter when the enrollment was created.
//   - `geography` - A list of where you can deploy the certificate. Either `core` to specify worldwide deployment (including China and Russia), `china+core` to specify worldwide deployment and China, or `russia+core` to specify worldwide deployment and Russia.
//   - `mustHaveCiphers` - If present, shows ciphers included for enrollment when deployed on the network. The default is `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
//   - `ocspStapling` - If present, the enrollment is using OCSP stapling. OCSP stapling improves performance by including a valid OCSP response in every TLS handshake. This option allows the visitors on your site to query the Online Certificate Status Protocol (OCSP) server at regular intervals to obtain a signed time-stamped OCSP response. Possible values are `on`, `off`, or `not-set`.
//   - `preferredCiphers` - If present, shows the ciphers that you prefer to include for the enrollment while deploying it on the network. The default is `ak-akamai-default` when its not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
//   - `quicEnabled` - If present, uses the QUIC transport layer network protocol.
//   - `signatureAlgorithm` - If present, shows the Secure Hash Algorithm (SHA) function, either `SHA-1` or `SHA-256`.
//   - `techContact` - The technical contact within Akamai. This is the person you work closest with at Akamai and who can verify the certificate request. The CA calls this contact if there are any issues with the certificate and they can't reach the `adminContact`.
//   - `organization` - The name of the organization in Akamai where your technical contact works.
//   - `name` - The name of the technical contact at Akamai.
//   - `phone` - The phone number of the technical contact at Akamai.
//   - `addressLineOne` - The address for the technical contact at Akamai.
//   - `addressLineTwo` - The address for the technical contact at Akamai.
//   - `city` - The address for the technical contact at Akamai.
//   - `region` - The region for the technical contact at Akamai.
//   - `postalCode` - The postal code for the technical contact at Akamai.
//   - `countryCode` - The country code for the technical contact at Akamai.
//   - `certificateType` - Populates automatically with the `san` certificate type and is preserved in the `state` file.
//   - `validationType` - Populates automatically with the `dv` validation type and is preserved in the `state` file.
//   - `registrationAuthority` - Populates automatically with the `lets-encrypt` certificate type and is preserved in the `state` file.
//   - `pendingChanges` - If `true`, there are changes currently pending in CPS. To view pending changes, use the `dataAkamaiCpsEnrollment` data source.
func GetCPSEnrollments(ctx *pulumi.Context, args *GetCPSEnrollmentsArgs, opts ...pulumi.InvokeOption) (*GetCPSEnrollmentsResult, error) {
	var rv GetCPSEnrollmentsResult
	err := ctx.Invoke("akamai:index/getCPSEnrollments:getCPSEnrollments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCPSEnrollments.
type GetCPSEnrollmentsArgs struct {
	// A contract's ID, optionally with the `ctr_` prefix.
	ContractId string `pulumi:"contractId"`
}

// A collection of values returned by getCPSEnrollments.
type GetCPSEnrollmentsResult struct {
	ContractId  string                        `pulumi:"contractId"`
	Enrollments []GetCPSEnrollmentsEnrollment `pulumi:"enrollments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetCPSEnrollmentsOutput(ctx *pulumi.Context, args GetCPSEnrollmentsOutputArgs, opts ...pulumi.InvokeOption) GetCPSEnrollmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCPSEnrollmentsResult, error) {
			args := v.(GetCPSEnrollmentsArgs)
			r, err := GetCPSEnrollments(ctx, &args, opts...)
			var s GetCPSEnrollmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCPSEnrollmentsResultOutput)
}

// A collection of arguments for invoking getCPSEnrollments.
type GetCPSEnrollmentsOutputArgs struct {
	// A contract's ID, optionally with the `ctr_` prefix.
	ContractId pulumi.StringInput `pulumi:"contractId"`
}

func (GetCPSEnrollmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCPSEnrollmentsArgs)(nil)).Elem()
}

// A collection of values returned by getCPSEnrollments.
type GetCPSEnrollmentsResultOutput struct{ *pulumi.OutputState }

func (GetCPSEnrollmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCPSEnrollmentsResult)(nil)).Elem()
}

func (o GetCPSEnrollmentsResultOutput) ToGetCPSEnrollmentsResultOutput() GetCPSEnrollmentsResultOutput {
	return o
}

func (o GetCPSEnrollmentsResultOutput) ToGetCPSEnrollmentsResultOutputWithContext(ctx context.Context) GetCPSEnrollmentsResultOutput {
	return o
}

func (o GetCPSEnrollmentsResultOutput) ContractId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCPSEnrollmentsResult) string { return v.ContractId }).(pulumi.StringOutput)
}

func (o GetCPSEnrollmentsResultOutput) Enrollments() GetCPSEnrollmentsEnrollmentArrayOutput {
	return o.ApplyT(func(v GetCPSEnrollmentsResult) []GetCPSEnrollmentsEnrollment { return v.Enrollments }).(GetCPSEnrollmentsEnrollmentArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCPSEnrollmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCPSEnrollmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCPSEnrollmentsResultOutput{})
}
