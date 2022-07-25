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
    public sealed class GetIamTimezonesTimezoneResult
    {
        public readonly string Description;
        public readonly string Offset;
        public readonly string Posix;
        public readonly string Timezone;

        [OutputConstructor]
        private GetIamTimezonesTimezoneResult(
            string description,

            string offset,

            string posix,

            string timezone)
        {
            Description = description;
            Offset = offset;
            Posix = posix;
            Timezone = timezone;
        }
    }
}
