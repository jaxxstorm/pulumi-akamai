// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Outputs
{

    [OutputType]
    public sealed class GetCPSEnrollmentsEnrollmentResult
    {
        public readonly ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentAdminContactResult> AdminContacts;
        public readonly string CertificateChainType;
        public readonly string CertificateType;
        public readonly string CommonName;
        public readonly ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentCsrResult> Csrs;
        public readonly bool EnableMultiStackedCertificates;
        public readonly int EnrollmentId;
        public readonly ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentNetworkConfigurationResult> NetworkConfigurations;
        public readonly ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentOrganizationResult> Organizations;
        public readonly bool PendingChanges;
        public readonly string RegistrationAuthority;
        public readonly ImmutableArray<string> Sans;
        public readonly string SecureNetwork;
        public readonly string SignatureAlgorithm;
        public readonly bool SniOnly;
        public readonly ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentTechContactResult> TechContacts;
        public readonly string ValidationType;

        [OutputConstructor]
        private GetCPSEnrollmentsEnrollmentResult(
            ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentAdminContactResult> adminContacts,

            string certificateChainType,

            string certificateType,

            string commonName,

            ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentCsrResult> csrs,

            bool enableMultiStackedCertificates,

            int enrollmentId,

            ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentNetworkConfigurationResult> networkConfigurations,

            ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentOrganizationResult> organizations,

            bool pendingChanges,

            string registrationAuthority,

            ImmutableArray<string> sans,

            string secureNetwork,

            string signatureAlgorithm,

            bool sniOnly,

            ImmutableArray<Outputs.GetCPSEnrollmentsEnrollmentTechContactResult> techContacts,

            string validationType)
        {
            AdminContacts = adminContacts;
            CertificateChainType = certificateChainType;
            CertificateType = certificateType;
            CommonName = commonName;
            Csrs = csrs;
            EnableMultiStackedCertificates = enableMultiStackedCertificates;
            EnrollmentId = enrollmentId;
            NetworkConfigurations = networkConfigurations;
            Organizations = organizations;
            PendingChanges = pendingChanges;
            RegistrationAuthority = registrationAuthority;
            Sans = sans;
            SecureNetwork = secureNetwork;
            SignatureAlgorithm = signatureAlgorithm;
            SniOnly = sniOnly;
            TechContacts = techContacts;
            ValidationType = validationType;
        }
    }
}