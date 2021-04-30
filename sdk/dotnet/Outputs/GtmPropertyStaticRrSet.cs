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
    public sealed class GtmPropertyStaticRrSet
    {
        public readonly ImmutableArray<string> Rdatas;
        public readonly int? Ttl;
        public readonly string? Type;

        [OutputConstructor]
        private GtmPropertyStaticRrSet(
            ImmutableArray<string> rdatas,

            int? ttl,

            string? type)
        {
            Rdatas = rdatas;
            Ttl = ttl;
            Type = type;
        }
    }
}
