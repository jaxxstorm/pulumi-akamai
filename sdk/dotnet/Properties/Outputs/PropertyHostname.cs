// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Properties.Outputs
{

    [OutputType]
    public sealed class PropertyHostname
    {
        /// <summary>
        /// The certificate's provisioning type, either the default `CPS_MANAGED` type for the custom certificates you provision with the [Certificate Provisioning System (CPS)](https://techdocs.akamai.com/cps/docs), or `DEFAULT` for certificates provisioned automatically.
        /// </summary>
        public readonly string CertProvisioningType;
        public readonly ImmutableArray<Outputs.PropertyHostnameCertStatus> CertStatuses;
        /// <summary>
        /// A string containing the original origin's hostname. For example, `"example.org"`.
        /// </summary>
        public readonly string CnameFrom;
        /// <summary>
        /// A string containing the hostname for edge content. For example,  `"example.org.edgesuite.net"`.
        /// </summary>
        public readonly string CnameTo;
        public readonly string? CnameType;
        public readonly string? EdgeHostnameId;

        [OutputConstructor]
        private PropertyHostname(
            string certProvisioningType,

            ImmutableArray<Outputs.PropertyHostnameCertStatus> certStatuses,

            string cnameFrom,

            string cnameTo,

            string? cnameType,

            string? edgeHostnameId)
        {
            CertProvisioningType = certProvisioningType;
            CertStatuses = certStatuses;
            CnameFrom = cnameFrom;
            CnameTo = cnameTo;
            CnameType = cnameType;
            EdgeHostnameId = edgeHostnameId;
        }
    }
}
