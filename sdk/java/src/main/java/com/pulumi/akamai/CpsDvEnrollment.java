// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai;

import com.pulumi.akamai.CpsDvEnrollmentArgs;
import com.pulumi.akamai.Utilities;
import com.pulumi.akamai.inputs.CpsDvEnrollmentState;
import com.pulumi.akamai.outputs.CpsDvEnrollmentAdminContact;
import com.pulumi.akamai.outputs.CpsDvEnrollmentCsr;
import com.pulumi.akamai.outputs.CpsDvEnrollmentDnsChallenge;
import com.pulumi.akamai.outputs.CpsDvEnrollmentHttpChallenge;
import com.pulumi.akamai.outputs.CpsDvEnrollmentNetworkConfiguration;
import com.pulumi.akamai.outputs.CpsDvEnrollmentOrganization;
import com.pulumi.akamai.outputs.CpsDvEnrollmentTechContact;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Use the `akamai.CpsDvEnrollment` resource to create an enrollment with all the information about your certificate life cycle, from the time you request it, through removal or automatic renewal. You can treat an enrollment as a core container for all the operations you perform within CPS.
 * 
 * You can use this resource with `akamai.DnsRecord` or other third-party DNS provider to create DNS records, and `akamai.CpsDvValidation` to complete the certificate validation.
 * 
 * ## Example Usage
 * 
 * Basic usage:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.akamai.CpsDvEnrollment;
 * import com.pulumi.akamai.CpsDvEnrollmentArgs;
 * import com.pulumi.akamai.inputs.CpsDvEnrollmentAdminContactArgs;
 * import com.pulumi.akamai.inputs.CpsDvEnrollmentTechContactArgs;
 * import com.pulumi.akamai.inputs.CpsDvEnrollmentCsrArgs;
 * import com.pulumi.akamai.inputs.CpsDvEnrollmentNetworkConfigurationArgs;
 * import com.pulumi.akamai.inputs.CpsDvEnrollmentOrganizationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new CpsDvEnrollment(&#34;example&#34;, CpsDvEnrollmentArgs.builder()        
 *             .contractId(&#34;ctr_1-AB123&#34;)
 *             .acknowledgePreVerificationWarnings(true)
 *             .commonName(&#34;cps-test.example.net&#34;)
 *             .sans(            
 *                 &#34;san1.cps-test.example.net&#34;,
 *                 &#34;san2.cps-test.example.net&#34;)
 *             .secureNetwork(&#34;enhanced-tls&#34;)
 *             .sniOnly(true)
 *             .adminContact(CpsDvEnrollmentAdminContactArgs.builder()
 *                 .firstName(&#34;x1&#34;)
 *                 .lastName(&#34;x2&#34;)
 *                 .phone(&#34;123123123&#34;)
 *                 .email(&#34;x1x2@example.net&#34;)
 *                 .addressLineOne(&#34;150 Broadway&#34;)
 *                 .city(&#34;Cambridge&#34;)
 *                 .countryCode(&#34;US&#34;)
 *                 .organization(&#34;Akamai&#34;)
 *                 .postalCode(&#34;02142&#34;)
 *                 .region(&#34;MA&#34;)
 *                 .title(&#34;Administrator&#34;)
 *                 .build())
 *             .techContact(CpsDvEnrollmentTechContactArgs.builder()
 *                 .firstName(&#34;x3&#34;)
 *                 .lastName(&#34;x4&#34;)
 *                 .phone(&#34;123123123&#34;)
 *                 .email(&#34;x3x4@akamai.com&#34;)
 *                 .addressLineOne(&#34;150 Broadway&#34;)
 *                 .city(&#34;Cambridge&#34;)
 *                 .countryCode(&#34;US&#34;)
 *                 .organization(&#34;Akamai&#34;)
 *                 .postalCode(&#34;02142&#34;)
 *                 .region(&#34;MA&#34;)
 *                 .title(&#34;Administrator&#34;)
 *                 .build())
 *             .certificateChainType(&#34;default&#34;)
 *             .csr(CpsDvEnrollmentCsrArgs.builder()
 *                 .countryCode(&#34;US&#34;)
 *                 .city(&#34;cambridge&#34;)
 *                 .organization(&#34;Akamai&#34;)
 *                 .organizationalUnit(&#34;Dev&#34;)
 *                 .state(&#34;MA&#34;)
 *                 .build())
 *             .enableMultiStackedCertificates(false)
 *             .networkConfiguration(CpsDvEnrollmentNetworkConfigurationArgs.builder()
 *                 .disallowedTlsVersions(                
 *                     &#34;TLSv1&#34;,
 *                     &#34;TLSv1_1&#34;)
 *                 .cloneDnsNames(false)
 *                 .geography(&#34;core&#34;)
 *                 .ocspStapling(&#34;on&#34;)
 *                 .preferredCiphers(&#34;ak-akamai-2020q1&#34;)
 *                 .mustHaveCiphers(&#34;ak-akamai-2020q1&#34;)
 *                 .quicEnabled(false)
 *                 .build())
 *             .signatureAlgorithm(&#34;SHA-256&#34;)
 *             .organization(CpsDvEnrollmentOrganizationArgs.builder()
 *                 .name(&#34;Akamai&#34;)
 *                 .phone(&#34;123123123&#34;)
 *                 .addressLineOne(&#34;150 Broadway&#34;)
 *                 .city(&#34;Cambridge&#34;)
 *                 .countryCode(&#34;US&#34;)
 *                 .postalCode(&#34;02142&#34;)
 *                 .region(&#34;MA&#34;)
 *                 .build())
 *             .build());
 * 
 *         ctx.export(&#34;dnsChallenges&#34;, example.dnsChallenges());
 *         ctx.export(&#34;httpChallenges&#34;, example.httpChallenges());
 *         ctx.export(&#34;enrollmentId&#34;, example.id());
 *     }
 * }
 * ```
 * ## Attributes reference
 * 
 * The resource returns these attributes:
 * 
 * * `registration_authority` - (Required) This value populates automatically with the `lets-encrypt` certificate type and is preserved in the `state` file.
 * * `certificate_type` - (Required) This value populates automatically with the `san` certificate type and is preserved in the `state` file.
 * * `validation_type` - (Required) This value populates automatically with the `dv` validation type and is preserved in the `state` file.
 * * `id` - The unique identifier for this enrollment.
 * * `dns_challenges` - The validation challenge for the domains listed in the certificate. To successfully perform the validation, only one challenge for each domain must be complete, either `dns_challenges` or `http_challenges`.
 *   
 *   Returns these additional attributes:
 *   
 *       * `domain` - The domain to validate.
 *       * `full_path` - The URL where Akamai publishes `response_body` for Let&#39;s Encrypt to validate.
 *       * `response_body` - The data Let&#39;s Encrypt expects to find served at `full_path` URL.
 * * `http_challenges` - The validation challenge for the domains listed in the certificate. To successfully perform the validation, only one challenge for each domain must be complete, either `dns_challenges` or `http_challenges`.
 *   
 *   Returns these additional attributes:
 *   
 *       * `domain` - The domain to validate.
 *       * `full_path` - The URL where Akamai publishes `response_body` for Let&#39;s Encrypt to validate.
 *       * `response_body` - The data Let&#39;s Encrypt expects to find served at `full_path` URL.
 * 
 * ## Import
 * 
 * Basic Usagehcl resource &#34;akamai_cps_dv_enrollment&#34; &#34;example&#34; { # (resource arguments) } You can import your Akamai DV enrollment using a comma-delimited string of the enrollment ID and
 * 
 *  contract ID, optionally with the `ctr_` prefix. You have to enter the IDs in this order`enrollment_id,contract_id` For example
 * 
 * ```sh
 *  $ pulumi import akamai:index/cpsDvEnrollment:CpsDvEnrollment example 12345,1-AB123
 * ```
 * 
 */
@ResourceType(type="akamai:index/cpsDvEnrollment:CpsDvEnrollment")
public class CpsDvEnrollment extends com.pulumi.resources.CustomResource {
    /**
     * Whether you want to automatically acknowledge the validation warnings of the current job state and proceed with the execution of a change.
     * 
     */
    @Export(name="acknowledgePreVerificationWarnings", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> acknowledgePreVerificationWarnings;

    /**
     * @return Whether you want to automatically acknowledge the validation warnings of the current job state and proceed with the execution of a change.
     * 
     */
    public Output<Optional<Boolean>> acknowledgePreVerificationWarnings() {
        return Codegen.optional(this.acknowledgePreVerificationWarnings);
    }
    /**
     * Contact information for the certificate administrator at your company.
     * 
     */
    @Export(name="adminContact", type=CpsDvEnrollmentAdminContact.class, parameters={})
    private Output<CpsDvEnrollmentAdminContact> adminContact;

    /**
     * @return Contact information for the certificate administrator at your company.
     * 
     */
    public Output<CpsDvEnrollmentAdminContact> adminContact() {
        return this.adminContact;
    }
    @Export(name="allowDuplicateCommonName", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowDuplicateCommonName;

    public Output<Optional<Boolean>> allowDuplicateCommonName() {
        return Codegen.optional(this.allowDuplicateCommonName);
    }
    /**
     * Certificate trust chain type.
     * 
     */
    @Export(name="certificateChainType", type=String.class, parameters={})
    private Output</* @Nullable */ String> certificateChainType;

    /**
     * @return Certificate trust chain type.
     * 
     */
    public Output<Optional<String>> certificateChainType() {
        return Codegen.optional(this.certificateChainType);
    }
    @Export(name="certificateType", type=String.class, parameters={})
    private Output<String> certificateType;

    public Output<String> certificateType() {
        return this.certificateType;
    }
    /**
     * - (Required) The fully qualified domain name (FQDN) for which you plan to use your certificate. The domain name you specify here must be owned or have legal rights to use the domain by the company you specify as `organization`. The company that owns the domain name must be a legally incorporated entity and be active and in good standing.
     * 
     */
    @Export(name="commonName", type=String.class, parameters={})
    private Output<String> commonName;

    /**
     * @return - (Required) The fully qualified domain name (FQDN) for which you plan to use your certificate. The domain name you specify here must be owned or have legal rights to use the domain by the company you specify as `organization`. The company that owns the domain name must be a legally incorporated entity and be active and in good standing.
     * 
     */
    public Output<String> commonName() {
        return this.commonName;
    }
    /**
     * - (Required) A contract&#39;s ID, optionally with the `ctr_` prefix.
     * 
     */
    @Export(name="contractId", type=String.class, parameters={})
    private Output<String> contractId;

    /**
     * @return - (Required) A contract&#39;s ID, optionally with the `ctr_` prefix.
     * 
     */
    public Output<String> contractId() {
        return this.contractId;
    }
    /**
     * When you create an enrollment, you also generate a certificate signing request (CSR) using CPS. CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
     * 
     */
    @Export(name="csr", type=CpsDvEnrollmentCsr.class, parameters={})
    private Output<CpsDvEnrollmentCsr> csr;

    /**
     * @return When you create an enrollment, you also generate a certificate signing request (CSR) using CPS. CPS signs the CSR with the private key. The CSR contains all the information the CA needs to issue your certificate.
     * 
     */
    public Output<CpsDvEnrollmentCsr> csr() {
        return this.csr;
    }
    @Export(name="dnsChallenges", type=List.class, parameters={CpsDvEnrollmentDnsChallenge.class})
    private Output<List<CpsDvEnrollmentDnsChallenge>> dnsChallenges;

    public Output<List<CpsDvEnrollmentDnsChallenge>> dnsChallenges() {
        return this.dnsChallenges;
    }
    /**
     * Whether to enable an ECDSA certificate in addition to an RSA certificate. CPS automatically performs all certificate operations on both certificates, and uses the best certificate for each client connection to your secure properties. If you are pinning the certificates, you need to pin both the RSA and the ECDSA certificate.
     * 
     */
    @Export(name="enableMultiStackedCertificates", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableMultiStackedCertificates;

    /**
     * @return Whether to enable an ECDSA certificate in addition to an RSA certificate. CPS automatically performs all certificate operations on both certificates, and uses the best certificate for each client connection to your secure properties. If you are pinning the certificates, you need to pin both the RSA and the ECDSA certificate.
     * 
     */
    public Output<Optional<Boolean>> enableMultiStackedCertificates() {
        return Codegen.optional(this.enableMultiStackedCertificates);
    }
    @Export(name="httpChallenges", type=List.class, parameters={CpsDvEnrollmentHttpChallenge.class})
    private Output<List<CpsDvEnrollmentHttpChallenge>> httpChallenges;

    public Output<List<CpsDvEnrollmentHttpChallenge>> httpChallenges() {
        return this.httpChallenges;
    }
    /**
     * The network information and TLS Metadata you want CPS to use to push the completed certificate to the network.
     * 
     */
    @Export(name="networkConfiguration", type=CpsDvEnrollmentNetworkConfiguration.class, parameters={})
    private Output<CpsDvEnrollmentNetworkConfiguration> networkConfiguration;

    /**
     * @return The network information and TLS Metadata you want CPS to use to push the completed certificate to the network.
     * 
     */
    public Output<CpsDvEnrollmentNetworkConfiguration> networkConfiguration() {
        return this.networkConfiguration;
    }
    /**
     * Your organization information.
     * 
     */
    @Export(name="organization", type=CpsDvEnrollmentOrganization.class, parameters={})
    private Output<CpsDvEnrollmentOrganization> organization;

    /**
     * @return Your organization information.
     * 
     */
    public Output<CpsDvEnrollmentOrganization> organization() {
        return this.organization;
    }
    @Export(name="registrationAuthority", type=String.class, parameters={})
    private Output<String> registrationAuthority;

    public Output<String> registrationAuthority() {
        return this.registrationAuthority;
    }
    /**
     * Additional common names to create a Subject Alternative Names (SAN) list.
     * 
     */
    @Export(name="sans", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> sans;

    /**
     * @return Additional common names to create a Subject Alternative Names (SAN) list.
     * 
     */
    public Output<Optional<List<String>>> sans() {
        return Codegen.optional(this.sans);
    }
    /**
     * The type of deployment network you want to use. `standard-tls` deploys your certificate to Akamai&#39;s standard secure network, but it isn&#39;t PCI compliant. `enhanced-tls` deploys your certificate to Akamai&#39;s more secure network with PCI compliance capability.
     * 
     */
    @Export(name="secureNetwork", type=String.class, parameters={})
    private Output<String> secureNetwork;

    /**
     * @return The type of deployment network you want to use. `standard-tls` deploys your certificate to Akamai&#39;s standard secure network, but it isn&#39;t PCI compliant. `enhanced-tls` deploys your certificate to Akamai&#39;s more secure network with PCI compliance capability.
     * 
     */
    public Output<String> secureNetwork() {
        return this.secureNetwork;
    }
    /**
     * The Secure Hash Algorithm (SHA) function, either `SHA-1` or `SHA-256`.
     * 
     */
    @Export(name="signatureAlgorithm", type=String.class, parameters={})
    private Output<String> signatureAlgorithm;

    /**
     * @return The Secure Hash Algorithm (SHA) function, either `SHA-1` or `SHA-256`.
     * 
     */
    public Output<String> signatureAlgorithm() {
        return this.signatureAlgorithm;
    }
    /**
     * Whether you want to enable SNI-only extension for the enrollment. Server Name Indication (SNI) is an extension of the Transport Layer Security (TLS) networking protocol. It allows a server to present multiple certificates on the same IP address. All modern web browsers support the SNI extension. If you have the same SAN on two or more certificates with the SNI-only option set, Akamai may serve traffic using any certificate which matches the requested SNI hostname. You should avoid multiple certificates with overlapping SAN names when using SNI-only. You can&#39;t change this setting once an enrollment is created.
     * 
     */
    @Export(name="sniOnly", type=Boolean.class, parameters={})
    private Output<Boolean> sniOnly;

    /**
     * @return Whether you want to enable SNI-only extension for the enrollment. Server Name Indication (SNI) is an extension of the Transport Layer Security (TLS) networking protocol. It allows a server to present multiple certificates on the same IP address. All modern web browsers support the SNI extension. If you have the same SAN on two or more certificates with the SNI-only option set, Akamai may serve traffic using any certificate which matches the requested SNI hostname. You should avoid multiple certificates with overlapping SAN names when using SNI-only. You can&#39;t change this setting once an enrollment is created.
     * 
     */
    public Output<Boolean> sniOnly() {
        return this.sniOnly;
    }
    /**
     * The technical contact within Akamai. This is the person you work closest with at Akamai and who can verify the certificate request. The CA calls this contact if there are any issues with the certificate and they can&#39;t reach the `admin_contact`.
     * 
     */
    @Export(name="techContact", type=CpsDvEnrollmentTechContact.class, parameters={})
    private Output<CpsDvEnrollmentTechContact> techContact;

    /**
     * @return The technical contact within Akamai. This is the person you work closest with at Akamai and who can verify the certificate request. The CA calls this contact if there are any issues with the certificate and they can&#39;t reach the `admin_contact`.
     * 
     */
    public Output<CpsDvEnrollmentTechContact> techContact() {
        return this.techContact;
    }
    @Export(name="validationType", type=String.class, parameters={})
    private Output<String> validationType;

    public Output<String> validationType() {
        return this.validationType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CpsDvEnrollment(String name) {
        this(name, CpsDvEnrollmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CpsDvEnrollment(String name, CpsDvEnrollmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CpsDvEnrollment(String name, CpsDvEnrollmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/cpsDvEnrollment:CpsDvEnrollment", name, args == null ? CpsDvEnrollmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CpsDvEnrollment(String name, Output<String> id, @Nullable CpsDvEnrollmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("akamai:index/cpsDvEnrollment:CpsDvEnrollment", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CpsDvEnrollment get(String name, Output<String> id, @Nullable CpsDvEnrollmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CpsDvEnrollment(name, id, state, options);
    }
}
