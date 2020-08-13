// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Edgedns.Outputs
{

    [OutputType]
    public sealed class DnsZoneTsigKey
    {
        public readonly string Algorithm;
        /// <summary>
        /// key name
        /// * `algorithm`
        /// * `secret`
        /// </summary>
        public readonly string Name;
        public readonly string Secret;

        [OutputConstructor]
        private DnsZoneTsigKey(
            string algorithm,

            string name,

            string secret)
        {
            Algorithm = algorithm;
            Name = name;
            Secret = secret;
        }
    }
}