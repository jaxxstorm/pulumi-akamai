// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package akamai

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the `CpsDvEnrollment` resource to create an enrollment with all the information about your certificate life cycle, from the time you request it, through removal or automatic renewal. You can treat an enrollment as a core container for all the operations you perform within CPS.
//
// You can use this resource with `DnsRecord` or other third-party DNS provider to create DNS records, and `CpsDvValidation` to complete the certificate validation.
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
// 		example, err := akamai.NewCpsDvEnrollment(ctx, "example", &akamai.CpsDvEnrollmentArgs{
// 			ContractId:                         pulumi.String("ctr_1-AB123"),
// 			AcknowledgePreVerificationWarnings: pulumi.Bool(true),
// 			CommonName:                         pulumi.String("cps-test.example.net"),
// 			Sans: pulumi.StringArray{
// 				pulumi.String("san1.cps-test.example.net"),
// 				pulumi.String("san2.cps-test.example.net"),
// 			},
// 			SecureNetwork: pulumi.String("enhanced-tls"),
// 			SniOnly:       pulumi.Bool(true),
// 			AdminContact: &CpsDvEnrollmentAdminContactArgs{
// 				FirstName:      pulumi.String("x1"),
// 				LastName:       pulumi.String("x2"),
// 				Phone:          pulumi.String("123123123"),
// 				Email:          pulumi.String("x1x2@example.net"),
// 				AddressLineOne: pulumi.String("150 Broadway"),
// 				City:           pulumi.String("Cambridge"),
// 				CountryCode:    pulumi.String("US"),
// 				Organization:   pulumi.String("Akamai"),
// 				PostalCode:     pulumi.String("02142"),
// 				Region:         pulumi.String("MA"),
// 				Title:          pulumi.String("Administrator"),
// 			},
// 			TechContact: &CpsDvEnrollmentTechContactArgs{
// 				FirstName:      pulumi.String("x3"),
// 				LastName:       pulumi.String("x4"),
// 				Phone:          pulumi.String("123123123"),
// 				Email:          pulumi.String("x3x4@akamai.com"),
// 				AddressLineOne: pulumi.String("150 Broadway"),
// 				City:           pulumi.String("Cambridge"),
// 				CountryCode:    pulumi.String("US"),
// 				Organization:   pulumi.String("Akamai"),
// 				PostalCode:     pulumi.String("02142"),
// 				Region:         pulumi.String("MA"),
// 				Title:          pulumi.String("Administrator"),
// 			},
// 			CertificateChainType: pulumi.String("default"),
// 			Csr: &CpsDvEnrollmentCsrArgs{
// 				CountryCode:        pulumi.String("US"),
// 				City:               pulumi.String("cambridge"),
// 				Organization:       pulumi.String("Akamai"),
// 				OrganizationalUnit: pulumi.String("Dev"),
// 				State:              pulumi.String("MA"),
// 			},
// 			EnableMultiStackedCertificates: pulumi.Bool(false),
// 			NetworkConfiguration: &CpsDvEnrollmentNetworkConfigurationArgs{
// 				DisallowedTlsVersions: pulumi.StringArray{
// 					pulumi.String("TLSv1"),
// 					pulumi.String("TLSv1_1"),
// 				},
// 				CloneDnsNames:    pulumi.Bool(false),
// 				Geography:        pulumi.String("core"),
// 				OcspStapling:     pulumi.String("on"),
// 				PreferredCiphers: pulumi.String("ak-akamai-2020q1"),
// 				MustHaveCiphers:  pulumi.String("ak-akamai-2020q1"),
// 				QuicEnabled:      pulumi.Bool(false),
// 			},
// 			SignatureAlgorithm: pulumi.String("SHA-256"),
// 			Organization: &CpsDvEnrollmentOrganizationArgs{
// 				Name:           pulumi.String("Akamai"),
// 				Phone:          pulumi.String("123123123"),
// 				AddressLineOne: pulumi.String("150 Broadway"),
// 				City:           pulumi.String("Cambridge"),
// 				CountryCode:    pulumi.String("US"),
// 				PostalCode:     pulumi.String("02142"),
// 				Region:         pulumi.String("MA"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("dnsChallenges", example.DnsChallenges)
// 		ctx.Export("httpChallenges", example.HttpChallenges)
// 		ctx.Export("enrollmentId", example.ID())
// 		return nil
// 	})
// }
// ```
// ## Argument reference
//
// The following arguments are supported:
//
// * `contractId` - (Required) A contract's ID, optionally with the `ctr_` prefix.
// * `commonName` - (Required) The fully qualified domain name (FQDN) for which you plan to use your certificate. The domain name you specify here must be owned or have legal rights to use the domain by the company you specify as `organization`. The company that owns the domain name must be a legally incorporated entity and be active and in good standing.
// * `sans` - (Optional) Additional common names to create a Subject Alternative Names (SAN) list.
// * `secureNetwork` - (Required) The type of deployment network you want to use. `standard-tls` deploys your certificate to Akamai's standard secure network, but it isn't PCI compliant. `enhanced-tls` deploys your certificate to Akamai's more secure network with PCI compliance capability.
// * `sniOnly` - (Required) Whether you want to enable SNI-only extension for the enrollment. Server Name Indication (SNI) is an extension of the Transport Layer Security (TLS) networking protocol. It allows a server to present multiple certificates on the same IP address. All modern web browsers support the SNI extension. If you have the same SAN on two or more certificates with the SNI-only option set, Akamai may serve traffic using any certificate which matches the requested SNI hostname. You should avoid multiple certificates with overlapping SAN names when using SNI-only. You can't change this setting once an enrollment is created.
// * `acknowledgePreVerificationWarnings` - (Optional) Whether you want to automatically acknowledge the validation warnings of the current job state and proceed with the execution of a change.
// * `adminContact` - (Required) Contact information for the certificate administrator at your company.
//
//   Requires these additional arguments:
//
//       * `firstName` - (Required) The first name of the certificate administrator at your company.
//       * `lastName` - (Required) The last name of the certificate administrator at your company.
//       * `title` - (Optional) The title of the certificate administrator at your company.
//       * `organization` - (Required) The name of your organization.
//       * `email` - (Required) The email address of the administrator who you want to use as a contact at your company.
//       * `phone` - (Required) The phone number of your organization.
//       * `addressLineOne` - (Required) The address of your organization.
//       * `addressLineTwo` - (Optional) The address of your organization.
//       * `city` - (Required) The city where your organization resides.
//       * `region` - (Required) The region of your organization, typically a state or province.
//       * `postalCode` - (Required) The postal code of your organization.
//       * `countryCode` - (Required) The code for the counrty where your organization resides.
// * `certificateChainType` - (Optional) Certificate trust chain type.
// * `csr` - (Required) 	When you create an enrollment, you also generate a certificate signing request (CSR) using CPS. CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
//
//   Requires these additional arguments:
//
//       * `countryCode` - (Required) The country code for the country where your organization is located.
//       * `city` - (Required) The city where your organization resides.
//       * `organization` - (Required The name of your company or organization. Enter the name as it appears in all legal documents and as it appears in the legal entity filing.
//       * `organizationalUnit` - (Required) Your organizational unit.
//       * `state` - (Required) 	Your state or province.
// * `enableMultiStackedCertificates` - (Optional) Whether to enable an ECDSA certificate in addition to an RSA certificate. CPS automatically performs all certificate operations on both certificates, and uses the best certificate for each client connection to your secure properties. If you are pinning the certificates, you need to pin both the RSA and the ECDSA certificate.
// * `networkConfiguration` - (Required) The network information and TLS Metadata you want CPS to use to push the completed certificate to the network.
//
//   Requires these additional arguments:
//
//       * `clientMutualAuthentication` - (Optional) The configuration for client mutual authentication. Specifies the trust chain that is used to verify client certificates and some configuration options.
//
//         Requires these additional arguments:
//
//          * `sendCaListToClient` - (Optional) Whether you want to enable the server to send the certificate authority (CA) list to the client.
//          * `ocspEnabled` - (Optional) Whether you want to enable the Online Certificate Status Protocol (OCSP) stapling for client certificates.
//          * `setId` - (Optional) The identifier of the set of trust chains, created in the [Trust Chain Manager](https://developer.akamai.com/api/web_performance/trust_chain_manager/v1.html).
//       * `disallowedTlsVersions` - (Optional) The TLS protocol version to disallow. CPS uses the TLS protocols that Akamai currently supports as a best practice.
//       * `cloneDnsNames` - (Optional) Whether CPS should direct traffic using all the SANs you listed in the SANs parameter when you created your enrollment.
//       * `geography` - (Required) Lists where you can deploy the certificate. Either `core` to specify worldwide deployment (including China and Russia), `china+core` to specify worldwide deployment and China, or `russia+core` to specify worldwide deployment and Russia. You can only use the setting to include China and Russia if your Akamai contract specifies your ability to do so and you have approval from the Chinese and Russian government.
//       * `mustHaveCiphers` - (Optional) The ciphers to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
//       * `ocspStapling` - (Optional) Whether to use OCSP stapling for the enrollment, either `on`, `off` or `not-set`. OCSP Stapling improves performance by including a valid OCSP response in every TLS handshake. This option allows the visitors on your site to query the Online Certificate Status Protocol (OCSP) server at regular intervals to obtain a signed time-stamped OCSP response. This response must be signed by the CA, not the server, therefore ensuring security. Disable OSCP Stapling if you want visitors to your site to contact the CA directly for an OSCP response. OCSP allows you to obtain the revocation status of a certificate.
//       * `preferredCiphers` - (Optional) Ciphers that you preferably want to include for the enrollment while deploying it on the network. Defaults to `ak-akamai-default` when it is not set. For more information on cipher profiles, see [Akamai community](https://community.akamai.com/customers/s/article/SSL-TLS-Cipher-Profiles-for-Akamai-Secure-CDNrxdxm).
//       * `quicEnabled` - (Optional) Whether to use the QUIC transport layer network protocol.
// * `signatureAlgorithm` - (Required) The Secure Hash Algorithm (SHA) function, either `SHA-1` or `SHA-256`.
// * `techContact` - (Required) The technical contact within Akamai. This is the person you work closest with at Akamai and who can verify the certificate request. The CA calls this contact if there are any issues with the certificate and they can't reach the `adminContact`.
//
//   Requires these additional arguments:
//
//       * `firstName` - (Required) The first name of the technical contact at Akamai.
//       * `lastName` - (Required) The last name of the technical contact at Akamai.
//       * `title` - (Optional) The title of the technical contact at Akamai.
//       * `organization` - (Required) The name of the organization in Akamai where your technical contact works.
//       * `email` - (Required) The email address of the technical contact at Akamai, accessible at the `akamai.com` domain.
//       * `phone` - (Required) The phone number of the technical contact at Akamai.
//       * `addressLineOne` - (Required) The address for the technical contact at Akamai.
//       * `addressLineTwo` - (Optional) The address for the technical contact at Akamai.
//       * `city` - (Required) The address for the technical contact at Akamai.
//       * `region` - (Required) The region for the technical contact at Akamai.
//       * `postalCode` - (Required) The postal code for the technical contact at Akamai.
//       * `countryCode` - (Required) The country code for the technical contact at Akamai.
// * `organization` - (Required) Your organization information.
//
//   Requires these additional arguments:
//
//       * `name` - (Required) The name of your organization.
//       * `phone` - (Required) The phone number of the administrator who you want to use as a contact at your company.
//       * `addressLineOne` - (Required) The address of your organization.
//       * `addressLineTwo` - (Optional) The address of your organization.
//       * `city` - (Required) The city where your organization resides.
//       * `region` - (Required) The region of your organization, typically a state or province.
//       * `postalCode` - (Required) The postal code of your organization.
//       * `countryCode` - (Required) The code for the country where your organization resides.
//
// ## Attributes reference
//
// The resource returns these attributes:
//
// * `registrationAuthority` - (Required) This value populates automatically with the `lets-encrypt` certificate type and is preserved in the `state` file.
// * `certificateType` - (Required) This value populates automatically with the `san` certificate type and is preserved in the `state` file.
// * `validationType` - (Required) This value populates automatically with the `dv` validation type and is preserved in the `state` file.
// * `id` - The unique identifier for this enrollment.
// * `dnsChallenges` - The validation challenge for the domains listed in the certificate. To successfully perform the validation, only one challenge for each domain must be complete, either `dnsChallenges` or `httpChallenges`.
//
//   Returns these additional attributes:
//
//       * `domain` - The domain to validate.
//       * `fullPath` - The URL where Akamai publishes `responseBody` for Let's Encrypt to validate.
//       * `responseBody` - The data Let's Encrypt expects to find served at `fullPath` URL.
// * `httpChallenges` - The validation challenge for the domains listed in the certificate. To successfully perform the validation, only one challenge for each domain must be complete, either `dnsChallenges` or `httpChallenges`.
//
//   Returns these additional attributes:
//
//       * `domain` - The domain to validate.
//       * `fullPath` - The URL where Akamai publishes `responseBody` for Let's Encrypt to validate.
//       * `responseBody` - The data Let's Encrypt expects to find served at `fullPath` URL.
//
// ## Import
//
// Basic Usagehcl resource "akamai_cps_dv_enrollment" "example" { # (resource arguments) } You can import your Akamai DV enrollment using a comma-delimited string of the enrollment ID and
//
//  contract ID, optionally with the `ctr_` prefix. You have to enter the IDs in this order`enrollment_id,contract_id` For example
//
// ```sh
//  $ pulumi import akamai:index/cpsDvEnrollment:CpsDvEnrollment example 12345,1-AB123
// ```
type CpsDvEnrollment struct {
	pulumi.CustomResourceState

	AcknowledgePreVerificationWarnings pulumi.BoolPtrOutput                      `pulumi:"acknowledgePreVerificationWarnings"`
	AdminContact                       CpsDvEnrollmentAdminContactOutput         `pulumi:"adminContact"`
	CertificateChainType               pulumi.StringPtrOutput                    `pulumi:"certificateChainType"`
	CertificateType                    pulumi.StringOutput                       `pulumi:"certificateType"`
	CommonName                         pulumi.StringOutput                       `pulumi:"commonName"`
	ContractId                         pulumi.StringOutput                       `pulumi:"contractId"`
	Csr                                CpsDvEnrollmentCsrOutput                  `pulumi:"csr"`
	DnsChallenges                      CpsDvEnrollmentDnsChallengeArrayOutput    `pulumi:"dnsChallenges"`
	EnableMultiStackedCertificates     pulumi.BoolPtrOutput                      `pulumi:"enableMultiStackedCertificates"`
	HttpChallenges                     CpsDvEnrollmentHttpChallengeArrayOutput   `pulumi:"httpChallenges"`
	NetworkConfiguration               CpsDvEnrollmentNetworkConfigurationOutput `pulumi:"networkConfiguration"`
	Organization                       CpsDvEnrollmentOrganizationOutput         `pulumi:"organization"`
	RegistrationAuthority              pulumi.StringOutput                       `pulumi:"registrationAuthority"`
	Sans                               pulumi.StringArrayOutput                  `pulumi:"sans"`
	SecureNetwork                      pulumi.StringOutput                       `pulumi:"secureNetwork"`
	SignatureAlgorithm                 pulumi.StringOutput                       `pulumi:"signatureAlgorithm"`
	SniOnly                            pulumi.BoolOutput                         `pulumi:"sniOnly"`
	TechContact                        CpsDvEnrollmentTechContactOutput          `pulumi:"techContact"`
	ValidationType                     pulumi.StringOutput                       `pulumi:"validationType"`
}

// NewCpsDvEnrollment registers a new resource with the given unique name, arguments, and options.
func NewCpsDvEnrollment(ctx *pulumi.Context,
	name string, args *CpsDvEnrollmentArgs, opts ...pulumi.ResourceOption) (*CpsDvEnrollment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminContact == nil {
		return nil, errors.New("invalid value for required argument 'AdminContact'")
	}
	if args.CommonName == nil {
		return nil, errors.New("invalid value for required argument 'CommonName'")
	}
	if args.ContractId == nil {
		return nil, errors.New("invalid value for required argument 'ContractId'")
	}
	if args.Csr == nil {
		return nil, errors.New("invalid value for required argument 'Csr'")
	}
	if args.NetworkConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConfiguration'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.SecureNetwork == nil {
		return nil, errors.New("invalid value for required argument 'SecureNetwork'")
	}
	if args.SignatureAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'SignatureAlgorithm'")
	}
	if args.SniOnly == nil {
		return nil, errors.New("invalid value for required argument 'SniOnly'")
	}
	if args.TechContact == nil {
		return nil, errors.New("invalid value for required argument 'TechContact'")
	}
	var resource CpsDvEnrollment
	err := ctx.RegisterResource("akamai:index/cpsDvEnrollment:CpsDvEnrollment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCpsDvEnrollment gets an existing CpsDvEnrollment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCpsDvEnrollment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CpsDvEnrollmentState, opts ...pulumi.ResourceOption) (*CpsDvEnrollment, error) {
	var resource CpsDvEnrollment
	err := ctx.ReadResource("akamai:index/cpsDvEnrollment:CpsDvEnrollment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CpsDvEnrollment resources.
type cpsDvEnrollmentState struct {
	AcknowledgePreVerificationWarnings *bool                                `pulumi:"acknowledgePreVerificationWarnings"`
	AdminContact                       *CpsDvEnrollmentAdminContact         `pulumi:"adminContact"`
	CertificateChainType               *string                              `pulumi:"certificateChainType"`
	CertificateType                    *string                              `pulumi:"certificateType"`
	CommonName                         *string                              `pulumi:"commonName"`
	ContractId                         *string                              `pulumi:"contractId"`
	Csr                                *CpsDvEnrollmentCsr                  `pulumi:"csr"`
	DnsChallenges                      []CpsDvEnrollmentDnsChallenge        `pulumi:"dnsChallenges"`
	EnableMultiStackedCertificates     *bool                                `pulumi:"enableMultiStackedCertificates"`
	HttpChallenges                     []CpsDvEnrollmentHttpChallenge       `pulumi:"httpChallenges"`
	NetworkConfiguration               *CpsDvEnrollmentNetworkConfiguration `pulumi:"networkConfiguration"`
	Organization                       *CpsDvEnrollmentOrganization         `pulumi:"organization"`
	RegistrationAuthority              *string                              `pulumi:"registrationAuthority"`
	Sans                               []string                             `pulumi:"sans"`
	SecureNetwork                      *string                              `pulumi:"secureNetwork"`
	SignatureAlgorithm                 *string                              `pulumi:"signatureAlgorithm"`
	SniOnly                            *bool                                `pulumi:"sniOnly"`
	TechContact                        *CpsDvEnrollmentTechContact          `pulumi:"techContact"`
	ValidationType                     *string                              `pulumi:"validationType"`
}

type CpsDvEnrollmentState struct {
	AcknowledgePreVerificationWarnings pulumi.BoolPtrInput
	AdminContact                       CpsDvEnrollmentAdminContactPtrInput
	CertificateChainType               pulumi.StringPtrInput
	CertificateType                    pulumi.StringPtrInput
	CommonName                         pulumi.StringPtrInput
	ContractId                         pulumi.StringPtrInput
	Csr                                CpsDvEnrollmentCsrPtrInput
	DnsChallenges                      CpsDvEnrollmentDnsChallengeArrayInput
	EnableMultiStackedCertificates     pulumi.BoolPtrInput
	HttpChallenges                     CpsDvEnrollmentHttpChallengeArrayInput
	NetworkConfiguration               CpsDvEnrollmentNetworkConfigurationPtrInput
	Organization                       CpsDvEnrollmentOrganizationPtrInput
	RegistrationAuthority              pulumi.StringPtrInput
	Sans                               pulumi.StringArrayInput
	SecureNetwork                      pulumi.StringPtrInput
	SignatureAlgorithm                 pulumi.StringPtrInput
	SniOnly                            pulumi.BoolPtrInput
	TechContact                        CpsDvEnrollmentTechContactPtrInput
	ValidationType                     pulumi.StringPtrInput
}

func (CpsDvEnrollmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*cpsDvEnrollmentState)(nil)).Elem()
}

type cpsDvEnrollmentArgs struct {
	AcknowledgePreVerificationWarnings *bool                               `pulumi:"acknowledgePreVerificationWarnings"`
	AdminContact                       CpsDvEnrollmentAdminContact         `pulumi:"adminContact"`
	CertificateChainType               *string                             `pulumi:"certificateChainType"`
	CommonName                         string                              `pulumi:"commonName"`
	ContractId                         string                              `pulumi:"contractId"`
	Csr                                CpsDvEnrollmentCsr                  `pulumi:"csr"`
	EnableMultiStackedCertificates     *bool                               `pulumi:"enableMultiStackedCertificates"`
	NetworkConfiguration               CpsDvEnrollmentNetworkConfiguration `pulumi:"networkConfiguration"`
	Organization                       CpsDvEnrollmentOrganization         `pulumi:"organization"`
	Sans                               []string                            `pulumi:"sans"`
	SecureNetwork                      string                              `pulumi:"secureNetwork"`
	SignatureAlgorithm                 string                              `pulumi:"signatureAlgorithm"`
	SniOnly                            bool                                `pulumi:"sniOnly"`
	TechContact                        CpsDvEnrollmentTechContact          `pulumi:"techContact"`
}

// The set of arguments for constructing a CpsDvEnrollment resource.
type CpsDvEnrollmentArgs struct {
	AcknowledgePreVerificationWarnings pulumi.BoolPtrInput
	AdminContact                       CpsDvEnrollmentAdminContactInput
	CertificateChainType               pulumi.StringPtrInput
	CommonName                         pulumi.StringInput
	ContractId                         pulumi.StringInput
	Csr                                CpsDvEnrollmentCsrInput
	EnableMultiStackedCertificates     pulumi.BoolPtrInput
	NetworkConfiguration               CpsDvEnrollmentNetworkConfigurationInput
	Organization                       CpsDvEnrollmentOrganizationInput
	Sans                               pulumi.StringArrayInput
	SecureNetwork                      pulumi.StringInput
	SignatureAlgorithm                 pulumi.StringInput
	SniOnly                            pulumi.BoolInput
	TechContact                        CpsDvEnrollmentTechContactInput
}

func (CpsDvEnrollmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cpsDvEnrollmentArgs)(nil)).Elem()
}

type CpsDvEnrollmentInput interface {
	pulumi.Input

	ToCpsDvEnrollmentOutput() CpsDvEnrollmentOutput
	ToCpsDvEnrollmentOutputWithContext(ctx context.Context) CpsDvEnrollmentOutput
}

func (*CpsDvEnrollment) ElementType() reflect.Type {
	return reflect.TypeOf((*CpsDvEnrollment)(nil))
}

func (i *CpsDvEnrollment) ToCpsDvEnrollmentOutput() CpsDvEnrollmentOutput {
	return i.ToCpsDvEnrollmentOutputWithContext(context.Background())
}

func (i *CpsDvEnrollment) ToCpsDvEnrollmentOutputWithContext(ctx context.Context) CpsDvEnrollmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvEnrollmentOutput)
}

func (i *CpsDvEnrollment) ToCpsDvEnrollmentPtrOutput() CpsDvEnrollmentPtrOutput {
	return i.ToCpsDvEnrollmentPtrOutputWithContext(context.Background())
}

func (i *CpsDvEnrollment) ToCpsDvEnrollmentPtrOutputWithContext(ctx context.Context) CpsDvEnrollmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvEnrollmentPtrOutput)
}

type CpsDvEnrollmentPtrInput interface {
	pulumi.Input

	ToCpsDvEnrollmentPtrOutput() CpsDvEnrollmentPtrOutput
	ToCpsDvEnrollmentPtrOutputWithContext(ctx context.Context) CpsDvEnrollmentPtrOutput
}

type cpsDvEnrollmentPtrType CpsDvEnrollmentArgs

func (*cpsDvEnrollmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CpsDvEnrollment)(nil))
}

func (i *cpsDvEnrollmentPtrType) ToCpsDvEnrollmentPtrOutput() CpsDvEnrollmentPtrOutput {
	return i.ToCpsDvEnrollmentPtrOutputWithContext(context.Background())
}

func (i *cpsDvEnrollmentPtrType) ToCpsDvEnrollmentPtrOutputWithContext(ctx context.Context) CpsDvEnrollmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvEnrollmentPtrOutput)
}

// CpsDvEnrollmentArrayInput is an input type that accepts CpsDvEnrollmentArray and CpsDvEnrollmentArrayOutput values.
// You can construct a concrete instance of `CpsDvEnrollmentArrayInput` via:
//
//          CpsDvEnrollmentArray{ CpsDvEnrollmentArgs{...} }
type CpsDvEnrollmentArrayInput interface {
	pulumi.Input

	ToCpsDvEnrollmentArrayOutput() CpsDvEnrollmentArrayOutput
	ToCpsDvEnrollmentArrayOutputWithContext(context.Context) CpsDvEnrollmentArrayOutput
}

type CpsDvEnrollmentArray []CpsDvEnrollmentInput

func (CpsDvEnrollmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CpsDvEnrollment)(nil)).Elem()
}

func (i CpsDvEnrollmentArray) ToCpsDvEnrollmentArrayOutput() CpsDvEnrollmentArrayOutput {
	return i.ToCpsDvEnrollmentArrayOutputWithContext(context.Background())
}

func (i CpsDvEnrollmentArray) ToCpsDvEnrollmentArrayOutputWithContext(ctx context.Context) CpsDvEnrollmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvEnrollmentArrayOutput)
}

// CpsDvEnrollmentMapInput is an input type that accepts CpsDvEnrollmentMap and CpsDvEnrollmentMapOutput values.
// You can construct a concrete instance of `CpsDvEnrollmentMapInput` via:
//
//          CpsDvEnrollmentMap{ "key": CpsDvEnrollmentArgs{...} }
type CpsDvEnrollmentMapInput interface {
	pulumi.Input

	ToCpsDvEnrollmentMapOutput() CpsDvEnrollmentMapOutput
	ToCpsDvEnrollmentMapOutputWithContext(context.Context) CpsDvEnrollmentMapOutput
}

type CpsDvEnrollmentMap map[string]CpsDvEnrollmentInput

func (CpsDvEnrollmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CpsDvEnrollment)(nil)).Elem()
}

func (i CpsDvEnrollmentMap) ToCpsDvEnrollmentMapOutput() CpsDvEnrollmentMapOutput {
	return i.ToCpsDvEnrollmentMapOutputWithContext(context.Background())
}

func (i CpsDvEnrollmentMap) ToCpsDvEnrollmentMapOutputWithContext(ctx context.Context) CpsDvEnrollmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CpsDvEnrollmentMapOutput)
}

type CpsDvEnrollmentOutput struct{ *pulumi.OutputState }

func (CpsDvEnrollmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CpsDvEnrollment)(nil))
}

func (o CpsDvEnrollmentOutput) ToCpsDvEnrollmentOutput() CpsDvEnrollmentOutput {
	return o
}

func (o CpsDvEnrollmentOutput) ToCpsDvEnrollmentOutputWithContext(ctx context.Context) CpsDvEnrollmentOutput {
	return o
}

func (o CpsDvEnrollmentOutput) ToCpsDvEnrollmentPtrOutput() CpsDvEnrollmentPtrOutput {
	return o.ToCpsDvEnrollmentPtrOutputWithContext(context.Background())
}

func (o CpsDvEnrollmentOutput) ToCpsDvEnrollmentPtrOutputWithContext(ctx context.Context) CpsDvEnrollmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CpsDvEnrollment) *CpsDvEnrollment {
		return &v
	}).(CpsDvEnrollmentPtrOutput)
}

type CpsDvEnrollmentPtrOutput struct{ *pulumi.OutputState }

func (CpsDvEnrollmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CpsDvEnrollment)(nil))
}

func (o CpsDvEnrollmentPtrOutput) ToCpsDvEnrollmentPtrOutput() CpsDvEnrollmentPtrOutput {
	return o
}

func (o CpsDvEnrollmentPtrOutput) ToCpsDvEnrollmentPtrOutputWithContext(ctx context.Context) CpsDvEnrollmentPtrOutput {
	return o
}

func (o CpsDvEnrollmentPtrOutput) Elem() CpsDvEnrollmentOutput {
	return o.ApplyT(func(v *CpsDvEnrollment) CpsDvEnrollment {
		if v != nil {
			return *v
		}
		var ret CpsDvEnrollment
		return ret
	}).(CpsDvEnrollmentOutput)
}

type CpsDvEnrollmentArrayOutput struct{ *pulumi.OutputState }

func (CpsDvEnrollmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CpsDvEnrollment)(nil))
}

func (o CpsDvEnrollmentArrayOutput) ToCpsDvEnrollmentArrayOutput() CpsDvEnrollmentArrayOutput {
	return o
}

func (o CpsDvEnrollmentArrayOutput) ToCpsDvEnrollmentArrayOutputWithContext(ctx context.Context) CpsDvEnrollmentArrayOutput {
	return o
}

func (o CpsDvEnrollmentArrayOutput) Index(i pulumi.IntInput) CpsDvEnrollmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CpsDvEnrollment {
		return vs[0].([]CpsDvEnrollment)[vs[1].(int)]
	}).(CpsDvEnrollmentOutput)
}

type CpsDvEnrollmentMapOutput struct{ *pulumi.OutputState }

func (CpsDvEnrollmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CpsDvEnrollment)(nil))
}

func (o CpsDvEnrollmentMapOutput) ToCpsDvEnrollmentMapOutput() CpsDvEnrollmentMapOutput {
	return o
}

func (o CpsDvEnrollmentMapOutput) ToCpsDvEnrollmentMapOutputWithContext(ctx context.Context) CpsDvEnrollmentMapOutput {
	return o
}

func (o CpsDvEnrollmentMapOutput) MapIndex(k pulumi.StringInput) CpsDvEnrollmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CpsDvEnrollment {
		return vs[0].(map[string]CpsDvEnrollment)[vs[1].(string)]
	}).(CpsDvEnrollmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CpsDvEnrollmentInput)(nil)).Elem(), &CpsDvEnrollment{})
	pulumi.RegisterInputType(reflect.TypeOf((*CpsDvEnrollmentPtrInput)(nil)).Elem(), &CpsDvEnrollment{})
	pulumi.RegisterInputType(reflect.TypeOf((*CpsDvEnrollmentArrayInput)(nil)).Elem(), CpsDvEnrollmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CpsDvEnrollmentMapInput)(nil)).Elem(), CpsDvEnrollmentMap{})
	pulumi.RegisterOutputType(CpsDvEnrollmentOutput{})
	pulumi.RegisterOutputType(CpsDvEnrollmentPtrOutput{})
	pulumi.RegisterOutputType(CpsDvEnrollmentArrayOutput{})
	pulumi.RegisterOutputType(CpsDvEnrollmentMapOutput{})
}
