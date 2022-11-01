// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.akamai.outputs;

import com.pulumi.akamai.outputs.GetCPSEnrollmentAdminContact;
import com.pulumi.akamai.outputs.GetCPSEnrollmentCsr;
import com.pulumi.akamai.outputs.GetCPSEnrollmentDnsChallenge;
import com.pulumi.akamai.outputs.GetCPSEnrollmentHttpChallenge;
import com.pulumi.akamai.outputs.GetCPSEnrollmentNetworkConfiguration;
import com.pulumi.akamai.outputs.GetCPSEnrollmentOrganization;
import com.pulumi.akamai.outputs.GetCPSEnrollmentTechContact;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCPSEnrollmentResult {
    private List<GetCPSEnrollmentAdminContact> adminContacts;
    private String certificateChainType;
    private String certificateType;
    private String commonName;
    private String contractId;
    private List<GetCPSEnrollmentCsr> csrs;
    private List<GetCPSEnrollmentDnsChallenge> dnsChallenges;
    private Boolean enableMultiStackedCertificates;
    private Integer enrollmentId;
    private List<GetCPSEnrollmentHttpChallenge> httpChallenges;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<GetCPSEnrollmentNetworkConfiguration> networkConfigurations;
    private List<GetCPSEnrollmentOrganization> organizations;
    private String registrationAuthority;
    private List<String> sans;
    private String secureNetwork;
    private String signatureAlgorithm;
    private Boolean sniOnly;
    private List<GetCPSEnrollmentTechContact> techContacts;
    private String validationType;

    private GetCPSEnrollmentResult() {}
    public List<GetCPSEnrollmentAdminContact> adminContacts() {
        return this.adminContacts;
    }
    public String certificateChainType() {
        return this.certificateChainType;
    }
    public String certificateType() {
        return this.certificateType;
    }
    public String commonName() {
        return this.commonName;
    }
    public String contractId() {
        return this.contractId;
    }
    public List<GetCPSEnrollmentCsr> csrs() {
        return this.csrs;
    }
    public List<GetCPSEnrollmentDnsChallenge> dnsChallenges() {
        return this.dnsChallenges;
    }
    public Boolean enableMultiStackedCertificates() {
        return this.enableMultiStackedCertificates;
    }
    public Integer enrollmentId() {
        return this.enrollmentId;
    }
    public List<GetCPSEnrollmentHttpChallenge> httpChallenges() {
        return this.httpChallenges;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<GetCPSEnrollmentNetworkConfiguration> networkConfigurations() {
        return this.networkConfigurations;
    }
    public List<GetCPSEnrollmentOrganization> organizations() {
        return this.organizations;
    }
    public String registrationAuthority() {
        return this.registrationAuthority;
    }
    public List<String> sans() {
        return this.sans;
    }
    public String secureNetwork() {
        return this.secureNetwork;
    }
    public String signatureAlgorithm() {
        return this.signatureAlgorithm;
    }
    public Boolean sniOnly() {
        return this.sniOnly;
    }
    public List<GetCPSEnrollmentTechContact> techContacts() {
        return this.techContacts;
    }
    public String validationType() {
        return this.validationType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCPSEnrollmentResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetCPSEnrollmentAdminContact> adminContacts;
        private String certificateChainType;
        private String certificateType;
        private String commonName;
        private String contractId;
        private List<GetCPSEnrollmentCsr> csrs;
        private List<GetCPSEnrollmentDnsChallenge> dnsChallenges;
        private Boolean enableMultiStackedCertificates;
        private Integer enrollmentId;
        private List<GetCPSEnrollmentHttpChallenge> httpChallenges;
        private String id;
        private List<GetCPSEnrollmentNetworkConfiguration> networkConfigurations;
        private List<GetCPSEnrollmentOrganization> organizations;
        private String registrationAuthority;
        private List<String> sans;
        private String secureNetwork;
        private String signatureAlgorithm;
        private Boolean sniOnly;
        private List<GetCPSEnrollmentTechContact> techContacts;
        private String validationType;
        public Builder() {}
        public Builder(GetCPSEnrollmentResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminContacts = defaults.adminContacts;
    	      this.certificateChainType = defaults.certificateChainType;
    	      this.certificateType = defaults.certificateType;
    	      this.commonName = defaults.commonName;
    	      this.contractId = defaults.contractId;
    	      this.csrs = defaults.csrs;
    	      this.dnsChallenges = defaults.dnsChallenges;
    	      this.enableMultiStackedCertificates = defaults.enableMultiStackedCertificates;
    	      this.enrollmentId = defaults.enrollmentId;
    	      this.httpChallenges = defaults.httpChallenges;
    	      this.id = defaults.id;
    	      this.networkConfigurations = defaults.networkConfigurations;
    	      this.organizations = defaults.organizations;
    	      this.registrationAuthority = defaults.registrationAuthority;
    	      this.sans = defaults.sans;
    	      this.secureNetwork = defaults.secureNetwork;
    	      this.signatureAlgorithm = defaults.signatureAlgorithm;
    	      this.sniOnly = defaults.sniOnly;
    	      this.techContacts = defaults.techContacts;
    	      this.validationType = defaults.validationType;
        }

        @CustomType.Setter
        public Builder adminContacts(List<GetCPSEnrollmentAdminContact> adminContacts) {
            this.adminContacts = Objects.requireNonNull(adminContacts);
            return this;
        }
        public Builder adminContacts(GetCPSEnrollmentAdminContact... adminContacts) {
            return adminContacts(List.of(adminContacts));
        }
        @CustomType.Setter
        public Builder certificateChainType(String certificateChainType) {
            this.certificateChainType = Objects.requireNonNull(certificateChainType);
            return this;
        }
        @CustomType.Setter
        public Builder certificateType(String certificateType) {
            this.certificateType = Objects.requireNonNull(certificateType);
            return this;
        }
        @CustomType.Setter
        public Builder commonName(String commonName) {
            this.commonName = Objects.requireNonNull(commonName);
            return this;
        }
        @CustomType.Setter
        public Builder contractId(String contractId) {
            this.contractId = Objects.requireNonNull(contractId);
            return this;
        }
        @CustomType.Setter
        public Builder csrs(List<GetCPSEnrollmentCsr> csrs) {
            this.csrs = Objects.requireNonNull(csrs);
            return this;
        }
        public Builder csrs(GetCPSEnrollmentCsr... csrs) {
            return csrs(List.of(csrs));
        }
        @CustomType.Setter
        public Builder dnsChallenges(List<GetCPSEnrollmentDnsChallenge> dnsChallenges) {
            this.dnsChallenges = Objects.requireNonNull(dnsChallenges);
            return this;
        }
        public Builder dnsChallenges(GetCPSEnrollmentDnsChallenge... dnsChallenges) {
            return dnsChallenges(List.of(dnsChallenges));
        }
        @CustomType.Setter
        public Builder enableMultiStackedCertificates(Boolean enableMultiStackedCertificates) {
            this.enableMultiStackedCertificates = Objects.requireNonNull(enableMultiStackedCertificates);
            return this;
        }
        @CustomType.Setter
        public Builder enrollmentId(Integer enrollmentId) {
            this.enrollmentId = Objects.requireNonNull(enrollmentId);
            return this;
        }
        @CustomType.Setter
        public Builder httpChallenges(List<GetCPSEnrollmentHttpChallenge> httpChallenges) {
            this.httpChallenges = Objects.requireNonNull(httpChallenges);
            return this;
        }
        public Builder httpChallenges(GetCPSEnrollmentHttpChallenge... httpChallenges) {
            return httpChallenges(List.of(httpChallenges));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder networkConfigurations(List<GetCPSEnrollmentNetworkConfiguration> networkConfigurations) {
            this.networkConfigurations = Objects.requireNonNull(networkConfigurations);
            return this;
        }
        public Builder networkConfigurations(GetCPSEnrollmentNetworkConfiguration... networkConfigurations) {
            return networkConfigurations(List.of(networkConfigurations));
        }
        @CustomType.Setter
        public Builder organizations(List<GetCPSEnrollmentOrganization> organizations) {
            this.organizations = Objects.requireNonNull(organizations);
            return this;
        }
        public Builder organizations(GetCPSEnrollmentOrganization... organizations) {
            return organizations(List.of(organizations));
        }
        @CustomType.Setter
        public Builder registrationAuthority(String registrationAuthority) {
            this.registrationAuthority = Objects.requireNonNull(registrationAuthority);
            return this;
        }
        @CustomType.Setter
        public Builder sans(List<String> sans) {
            this.sans = Objects.requireNonNull(sans);
            return this;
        }
        public Builder sans(String... sans) {
            return sans(List.of(sans));
        }
        @CustomType.Setter
        public Builder secureNetwork(String secureNetwork) {
            this.secureNetwork = Objects.requireNonNull(secureNetwork);
            return this;
        }
        @CustomType.Setter
        public Builder signatureAlgorithm(String signatureAlgorithm) {
            this.signatureAlgorithm = Objects.requireNonNull(signatureAlgorithm);
            return this;
        }
        @CustomType.Setter
        public Builder sniOnly(Boolean sniOnly) {
            this.sniOnly = Objects.requireNonNull(sniOnly);
            return this;
        }
        @CustomType.Setter
        public Builder techContacts(List<GetCPSEnrollmentTechContact> techContacts) {
            this.techContacts = Objects.requireNonNull(techContacts);
            return this;
        }
        public Builder techContacts(GetCPSEnrollmentTechContact... techContacts) {
            return techContacts(List.of(techContacts));
        }
        @CustomType.Setter
        public Builder validationType(String validationType) {
            this.validationType = Objects.requireNonNull(validationType);
            return this;
        }
        public GetCPSEnrollmentResult build() {
            final var o = new GetCPSEnrollmentResult();
            o.adminContacts = adminContacts;
            o.certificateChainType = certificateChainType;
            o.certificateType = certificateType;
            o.commonName = commonName;
            o.contractId = contractId;
            o.csrs = csrs;
            o.dnsChallenges = dnsChallenges;
            o.enableMultiStackedCertificates = enableMultiStackedCertificates;
            o.enrollmentId = enrollmentId;
            o.httpChallenges = httpChallenges;
            o.id = id;
            o.networkConfigurations = networkConfigurations;
            o.organizations = organizations;
            o.registrationAuthority = registrationAuthority;
            o.sans = sans;
            o.secureNetwork = secureNetwork;
            o.signatureAlgorithm = signatureAlgorithm;
            o.sniOnly = sniOnly;
            o.techContacts = techContacts;
            o.validationType = validationType;
            return o;
        }
    }
}