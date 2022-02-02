// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use the `akamai.CpsDvEnrollment` resource to create an enrollment with all the information about your certificate life cycle, from the time you request it, through removal or automatic renewal. You can treat an enrollment as a core container for all the operations you perform within CPS.
 *
 * You can use this resource with `akamai.DnsRecord` or other third-party DNS provider to create DNS records, and `akamai.CpsDvValidation` to complete the certificate validation.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as akamai from "@pulumi/akamai";
 *
 * const example = new akamai.CpsDvEnrollment("example", {
 *     contractId: "ctr_1-AB123",
 *     acknowledgePreVerificationWarnings: true,
 *     commonName: "cps-test.example.net",
 *     sans: [
 *         "san1.cps-test.example.net",
 *         "san2.cps-test.example.net",
 *     ],
 *     secureNetwork: "enhanced-tls",
 *     sniOnly: true,
 *     adminContact: {
 *         firstName: "x1",
 *         lastName: "x2",
 *         phone: "123123123",
 *         email: "x1x2@example.net",
 *         addressLineOne: "150 Broadway",
 *         city: "Cambridge",
 *         countryCode: "US",
 *         organization: "Akamai",
 *         postalCode: "02142",
 *         region: "MA",
 *         title: "Administrator",
 *     },
 *     techContact: {
 *         firstName: "x3",
 *         lastName: "x4",
 *         phone: "123123123",
 *         email: "x3x4@akamai.com",
 *         addressLineOne: "150 Broadway",
 *         city: "Cambridge",
 *         countryCode: "US",
 *         organization: "Akamai",
 *         postalCode: "02142",
 *         region: "MA",
 *         title: "Administrator",
 *     },
 *     certificateChainType: "default",
 *     csr: {
 *         countryCode: "US",
 *         city: "cambridge",
 *         organization: "Akamai",
 *         organizationalUnit: "Dev",
 *         state: "MA",
 *     },
 *     enableMultiStackedCertificates: false,
 *     networkConfiguration: {
 *         disallowedTlsVersions: [
 *             "TLSv1",
 *             "TLSv1_1",
 *         ],
 *         cloneDnsNames: false,
 *         geography: "core",
 *         ocspStapling: "on",
 *         preferredCiphers: "ak-akamai-2020q1",
 *         mustHaveCiphers: "ak-akamai-2020q1",
 *         quicEnabled: false,
 *     },
 *     signatureAlgorithm: "SHA-256",
 *     organization: {
 *         name: "Akamai",
 *         phone: "123123123",
 *         addressLineOne: "150 Broadway",
 *         city: "Cambridge",
 *         countryCode: "US",
 *         postalCode: "02142",
 *         region: "MA",
 *     },
 * });
 * export const dnsChallenges = example.dnsChallenges;
 * export const httpChallenges = example.httpChallenges;
 * export const enrollmentId = example.id;
 * ```
 * ## Attributes reference
 *
 * The resource returns these attributes:
 *
 * * `registrationAuthority` - (Required) This value populates automatically with the `lets-encrypt` certificate type and is preserved in the `state` file.
 * * `certificateType` - (Required) This value populates automatically with the `san` certificate type and is preserved in the `state` file.
 * * `validationType` - (Required) This value populates automatically with the `dv` validation type and is preserved in the `state` file.
 * * `id` - The unique identifier for this enrollment.
 * * `dnsChallenges` - The validation challenge for the domains listed in the certificate. To successfully perform the validation, only one challenge for each domain must be complete, either `dnsChallenges` or `httpChallenges`.
 *   
 *   Returns these additional attributes:
 *   
 *       * `domain` - The domain to validate.
 *       * `fullPath` - The URL where Akamai publishes `responseBody` for Let's Encrypt to validate.
 *       * `responseBody` - The data Let's Encrypt expects to find served at `fullPath` URL.
 * * `httpChallenges` - The validation challenge for the domains listed in the certificate. To successfully perform the validation, only one challenge for each domain must be complete, either `dnsChallenges` or `httpChallenges`.
 *   
 *   Returns these additional attributes:
 *   
 *       * `domain` - The domain to validate.
 *       * `fullPath` - The URL where Akamai publishes `responseBody` for Let's Encrypt to validate.
 *       * `responseBody` - The data Let's Encrypt expects to find served at `fullPath` URL.
 *
 * ## Import
 *
 * Basic Usagehcl resource "akamai_cps_dv_enrollment" "example" { # (resource arguments) } You can import your Akamai DV enrollment using a comma-delimited string of the enrollment ID and
 *
 *  contract ID, optionally with the `ctr_` prefix. You have to enter the IDs in this order`enrollment_id,contract_id` For example
 *
 * ```sh
 *  $ pulumi import akamai:index/cpsDvEnrollment:CpsDvEnrollment example 12345,1-AB123
 * ```
 */
export class CpsDvEnrollment extends pulumi.CustomResource {
    /**
     * Get an existing CpsDvEnrollment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CpsDvEnrollmentState, opts?: pulumi.CustomResourceOptions): CpsDvEnrollment {
        return new CpsDvEnrollment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'akamai:index/cpsDvEnrollment:CpsDvEnrollment';

    /**
     * Returns true if the given object is an instance of CpsDvEnrollment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CpsDvEnrollment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CpsDvEnrollment.__pulumiType;
    }

    /**
     * Whether you want to automatically acknowledge the validation warnings of the current job state and proceed with the execution of a change.
     */
    public readonly acknowledgePreVerificationWarnings!: pulumi.Output<boolean | undefined>;
    /**
     * Contact information for the certificate administrator at your company.
     */
    public readonly adminContact!: pulumi.Output<outputs.CpsDvEnrollmentAdminContact>;
    /**
     * Certificate trust chain type.
     */
    public readonly certificateChainType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly certificateType!: pulumi.Output<string>;
    /**
     * - (Required) The fully qualified domain name (FQDN) for which you plan to use your certificate. The domain name you specify here must be owned or have legal rights to use the domain by the company you specify as `organization`. The company that owns the domain name must be a legally incorporated entity and be active and in good standing.
     */
    public readonly commonName!: pulumi.Output<string>;
    /**
     * - (Required) A contract's ID, optionally with the `ctr_` prefix.
     */
    public readonly contractId!: pulumi.Output<string>;
    /**
     * When you create an enrollment, you also generate a certificate signing request (CSR) using CPS. CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
     */
    public readonly csr!: pulumi.Output<outputs.CpsDvEnrollmentCsr>;
    public /*out*/ readonly dnsChallenges!: pulumi.Output<outputs.CpsDvEnrollmentDnsChallenge[]>;
    /**
     * Whether to enable an ECDSA certificate in addition to an RSA certificate. CPS automatically performs all certificate operations on both certificates, and uses the best certificate for each client connection to your secure properties. If you are pinning the certificates, you need to pin both the RSA and the ECDSA certificate.
     */
    public readonly enableMultiStackedCertificates!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly httpChallenges!: pulumi.Output<outputs.CpsDvEnrollmentHttpChallenge[]>;
    /**
     * The network information and TLS Metadata you want CPS to use to push the completed certificate to the network.
     */
    public readonly networkConfiguration!: pulumi.Output<outputs.CpsDvEnrollmentNetworkConfiguration>;
    /**
     * Your organization information.
     */
    public readonly organization!: pulumi.Output<outputs.CpsDvEnrollmentOrganization>;
    public /*out*/ readonly registrationAuthority!: pulumi.Output<string>;
    /**
     * Additional common names to create a Subject Alternative Names (SAN) list.
     */
    public readonly sans!: pulumi.Output<string[] | undefined>;
    /**
     * The type of deployment network you want to use. `standard-tls` deploys your certificate to Akamai's standard secure network, but it isn't PCI compliant. `enhanced-tls` deploys your certificate to Akamai's more secure network with PCI compliance capability.
     */
    public readonly secureNetwork!: pulumi.Output<string>;
    /**
     * The Secure Hash Algorithm (SHA) function, either `SHA-1` or `SHA-256`.
     */
    public readonly signatureAlgorithm!: pulumi.Output<string>;
    /**
     * Whether you want to enable SNI-only extension for the enrollment. Server Name Indication (SNI) is an extension of the Transport Layer Security (TLS) networking protocol. It allows a server to present multiple certificates on the same IP address. All modern web browsers support the SNI extension. If you have the same SAN on two or more certificates with the SNI-only option set, Akamai may serve traffic using any certificate which matches the requested SNI hostname. You should avoid multiple certificates with overlapping SAN names when using SNI-only. You can't change this setting once an enrollment is created.
     */
    public readonly sniOnly!: pulumi.Output<boolean>;
    /**
     * The technical contact within Akamai. This is the person you work closest with at Akamai and who can verify the certificate request. The CA calls this contact if there are any issues with the certificate and they can't reach the `adminContact`.
     */
    public readonly techContact!: pulumi.Output<outputs.CpsDvEnrollmentTechContact>;
    public /*out*/ readonly validationType!: pulumi.Output<string>;

    /**
     * Create a CpsDvEnrollment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CpsDvEnrollmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CpsDvEnrollmentArgs | CpsDvEnrollmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CpsDvEnrollmentState | undefined;
            resourceInputs["acknowledgePreVerificationWarnings"] = state ? state.acknowledgePreVerificationWarnings : undefined;
            resourceInputs["adminContact"] = state ? state.adminContact : undefined;
            resourceInputs["certificateChainType"] = state ? state.certificateChainType : undefined;
            resourceInputs["certificateType"] = state ? state.certificateType : undefined;
            resourceInputs["commonName"] = state ? state.commonName : undefined;
            resourceInputs["contractId"] = state ? state.contractId : undefined;
            resourceInputs["csr"] = state ? state.csr : undefined;
            resourceInputs["dnsChallenges"] = state ? state.dnsChallenges : undefined;
            resourceInputs["enableMultiStackedCertificates"] = state ? state.enableMultiStackedCertificates : undefined;
            resourceInputs["httpChallenges"] = state ? state.httpChallenges : undefined;
            resourceInputs["networkConfiguration"] = state ? state.networkConfiguration : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["registrationAuthority"] = state ? state.registrationAuthority : undefined;
            resourceInputs["sans"] = state ? state.sans : undefined;
            resourceInputs["secureNetwork"] = state ? state.secureNetwork : undefined;
            resourceInputs["signatureAlgorithm"] = state ? state.signatureAlgorithm : undefined;
            resourceInputs["sniOnly"] = state ? state.sniOnly : undefined;
            resourceInputs["techContact"] = state ? state.techContact : undefined;
            resourceInputs["validationType"] = state ? state.validationType : undefined;
        } else {
            const args = argsOrState as CpsDvEnrollmentArgs | undefined;
            if ((!args || args.adminContact === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminContact'");
            }
            if ((!args || args.commonName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'commonName'");
            }
            if ((!args || args.contractId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contractId'");
            }
            if ((!args || args.csr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'csr'");
            }
            if ((!args || args.networkConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkConfiguration'");
            }
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            if ((!args || args.secureNetwork === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secureNetwork'");
            }
            if ((!args || args.signatureAlgorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signatureAlgorithm'");
            }
            if ((!args || args.sniOnly === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sniOnly'");
            }
            if ((!args || args.techContact === undefined) && !opts.urn) {
                throw new Error("Missing required property 'techContact'");
            }
            resourceInputs["acknowledgePreVerificationWarnings"] = args ? args.acknowledgePreVerificationWarnings : undefined;
            resourceInputs["adminContact"] = args ? args.adminContact : undefined;
            resourceInputs["certificateChainType"] = args ? args.certificateChainType : undefined;
            resourceInputs["commonName"] = args ? args.commonName : undefined;
            resourceInputs["contractId"] = args ? args.contractId : undefined;
            resourceInputs["csr"] = args ? args.csr : undefined;
            resourceInputs["enableMultiStackedCertificates"] = args ? args.enableMultiStackedCertificates : undefined;
            resourceInputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["sans"] = args ? args.sans : undefined;
            resourceInputs["secureNetwork"] = args ? args.secureNetwork : undefined;
            resourceInputs["signatureAlgorithm"] = args ? args.signatureAlgorithm : undefined;
            resourceInputs["sniOnly"] = args ? args.sniOnly : undefined;
            resourceInputs["techContact"] = args ? args.techContact : undefined;
            resourceInputs["certificateType"] = undefined /*out*/;
            resourceInputs["dnsChallenges"] = undefined /*out*/;
            resourceInputs["httpChallenges"] = undefined /*out*/;
            resourceInputs["registrationAuthority"] = undefined /*out*/;
            resourceInputs["validationType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CpsDvEnrollment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CpsDvEnrollment resources.
 */
export interface CpsDvEnrollmentState {
    /**
     * Whether you want to automatically acknowledge the validation warnings of the current job state and proceed with the execution of a change.
     */
    acknowledgePreVerificationWarnings?: pulumi.Input<boolean>;
    /**
     * Contact information for the certificate administrator at your company.
     */
    adminContact?: pulumi.Input<inputs.CpsDvEnrollmentAdminContact>;
    /**
     * Certificate trust chain type.
     */
    certificateChainType?: pulumi.Input<string>;
    certificateType?: pulumi.Input<string>;
    /**
     * - (Required) The fully qualified domain name (FQDN) for which you plan to use your certificate. The domain name you specify here must be owned or have legal rights to use the domain by the company you specify as `organization`. The company that owns the domain name must be a legally incorporated entity and be active and in good standing.
     */
    commonName?: pulumi.Input<string>;
    /**
     * - (Required) A contract's ID, optionally with the `ctr_` prefix.
     */
    contractId?: pulumi.Input<string>;
    /**
     * When you create an enrollment, you also generate a certificate signing request (CSR) using CPS. CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
     */
    csr?: pulumi.Input<inputs.CpsDvEnrollmentCsr>;
    dnsChallenges?: pulumi.Input<pulumi.Input<inputs.CpsDvEnrollmentDnsChallenge>[]>;
    /**
     * Whether to enable an ECDSA certificate in addition to an RSA certificate. CPS automatically performs all certificate operations on both certificates, and uses the best certificate for each client connection to your secure properties. If you are pinning the certificates, you need to pin both the RSA and the ECDSA certificate.
     */
    enableMultiStackedCertificates?: pulumi.Input<boolean>;
    httpChallenges?: pulumi.Input<pulumi.Input<inputs.CpsDvEnrollmentHttpChallenge>[]>;
    /**
     * The network information and TLS Metadata you want CPS to use to push the completed certificate to the network.
     */
    networkConfiguration?: pulumi.Input<inputs.CpsDvEnrollmentNetworkConfiguration>;
    /**
     * Your organization information.
     */
    organization?: pulumi.Input<inputs.CpsDvEnrollmentOrganization>;
    registrationAuthority?: pulumi.Input<string>;
    /**
     * Additional common names to create a Subject Alternative Names (SAN) list.
     */
    sans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of deployment network you want to use. `standard-tls` deploys your certificate to Akamai's standard secure network, but it isn't PCI compliant. `enhanced-tls` deploys your certificate to Akamai's more secure network with PCI compliance capability.
     */
    secureNetwork?: pulumi.Input<string>;
    /**
     * The Secure Hash Algorithm (SHA) function, either `SHA-1` or `SHA-256`.
     */
    signatureAlgorithm?: pulumi.Input<string>;
    /**
     * Whether you want to enable SNI-only extension for the enrollment. Server Name Indication (SNI) is an extension of the Transport Layer Security (TLS) networking protocol. It allows a server to present multiple certificates on the same IP address. All modern web browsers support the SNI extension. If you have the same SAN on two or more certificates with the SNI-only option set, Akamai may serve traffic using any certificate which matches the requested SNI hostname. You should avoid multiple certificates with overlapping SAN names when using SNI-only. You can't change this setting once an enrollment is created.
     */
    sniOnly?: pulumi.Input<boolean>;
    /**
     * The technical contact within Akamai. This is the person you work closest with at Akamai and who can verify the certificate request. The CA calls this contact if there are any issues with the certificate and they can't reach the `adminContact`.
     */
    techContact?: pulumi.Input<inputs.CpsDvEnrollmentTechContact>;
    validationType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CpsDvEnrollment resource.
 */
export interface CpsDvEnrollmentArgs {
    /**
     * Whether you want to automatically acknowledge the validation warnings of the current job state and proceed with the execution of a change.
     */
    acknowledgePreVerificationWarnings?: pulumi.Input<boolean>;
    /**
     * Contact information for the certificate administrator at your company.
     */
    adminContact: pulumi.Input<inputs.CpsDvEnrollmentAdminContact>;
    /**
     * Certificate trust chain type.
     */
    certificateChainType?: pulumi.Input<string>;
    /**
     * - (Required) The fully qualified domain name (FQDN) for which you plan to use your certificate. The domain name you specify here must be owned or have legal rights to use the domain by the company you specify as `organization`. The company that owns the domain name must be a legally incorporated entity and be active and in good standing.
     */
    commonName: pulumi.Input<string>;
    /**
     * - (Required) A contract's ID, optionally with the `ctr_` prefix.
     */
    contractId: pulumi.Input<string>;
    /**
     * When you create an enrollment, you also generate a certificate signing request (CSR) using CPS. CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
     */
    csr: pulumi.Input<inputs.CpsDvEnrollmentCsr>;
    /**
     * Whether to enable an ECDSA certificate in addition to an RSA certificate. CPS automatically performs all certificate operations on both certificates, and uses the best certificate for each client connection to your secure properties. If you are pinning the certificates, you need to pin both the RSA and the ECDSA certificate.
     */
    enableMultiStackedCertificates?: pulumi.Input<boolean>;
    /**
     * The network information and TLS Metadata you want CPS to use to push the completed certificate to the network.
     */
    networkConfiguration: pulumi.Input<inputs.CpsDvEnrollmentNetworkConfiguration>;
    /**
     * Your organization information.
     */
    organization: pulumi.Input<inputs.CpsDvEnrollmentOrganization>;
    /**
     * Additional common names to create a Subject Alternative Names (SAN) list.
     */
    sans?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of deployment network you want to use. `standard-tls` deploys your certificate to Akamai's standard secure network, but it isn't PCI compliant. `enhanced-tls` deploys your certificate to Akamai's more secure network with PCI compliance capability.
     */
    secureNetwork: pulumi.Input<string>;
    /**
     * The Secure Hash Algorithm (SHA) function, either `SHA-1` or `SHA-256`.
     */
    signatureAlgorithm: pulumi.Input<string>;
    /**
     * Whether you want to enable SNI-only extension for the enrollment. Server Name Indication (SNI) is an extension of the Transport Layer Security (TLS) networking protocol. It allows a server to present multiple certificates on the same IP address. All modern web browsers support the SNI extension. If you have the same SAN on two or more certificates with the SNI-only option set, Akamai may serve traffic using any certificate which matches the requested SNI hostname. You should avoid multiple certificates with overlapping SAN names when using SNI-only. You can't change this setting once an enrollment is created.
     */
    sniOnly: pulumi.Input<boolean>;
    /**
     * The technical contact within Akamai. This is the person you work closest with at Akamai and who can verify the certificate request. The CA calls this contact if there are any issues with the certificate and they can't reach the `adminContact`.
     */
    techContact: pulumi.Input<inputs.CpsDvEnrollmentTechContact>;
}
