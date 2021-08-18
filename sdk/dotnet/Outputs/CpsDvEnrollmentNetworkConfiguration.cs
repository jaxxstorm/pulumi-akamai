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
    public sealed class CpsDvEnrollmentNetworkConfiguration
    {
        public readonly Outputs.CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication? ClientMutualAuthentication;
        public readonly bool? CloneDnsNames;
        public readonly ImmutableArray<string> DisallowedTlsVersions;
        public readonly string Geography;
        public readonly string? MustHaveCiphers;
        public readonly string? OcspStapling;
        public readonly string? PreferredCiphers;
        public readonly bool? QuicEnabled;

        [OutputConstructor]
        private CpsDvEnrollmentNetworkConfiguration(
            Outputs.CpsDvEnrollmentNetworkConfigurationClientMutualAuthentication? clientMutualAuthentication,

            bool? cloneDnsNames,

            ImmutableArray<string> disallowedTlsVersions,

            string geography,

            string? mustHaveCiphers,

            string? ocspStapling,

            string? preferredCiphers,

            bool? quicEnabled)
        {
            ClientMutualAuthentication = clientMutualAuthentication;
            CloneDnsNames = cloneDnsNames;
            DisallowedTlsVersions = disallowedTlsVersions;
            Geography = geography;
            MustHaveCiphers = mustHaveCiphers;
            OcspStapling = ocspStapling;
            PreferredCiphers = preferredCiphers;
            QuicEnabled = quicEnabled;
        }
    }
}
