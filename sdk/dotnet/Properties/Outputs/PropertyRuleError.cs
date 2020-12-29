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
    public sealed class PropertyRuleError
    {
        public readonly string? BehaviorName;
        public readonly string? Detail;
        public readonly string? ErrorLocation;
        public readonly string? Instance;
        public readonly int? StatusCode;
        public readonly string? Title;
        public readonly string? Type;

        [OutputConstructor]
        private PropertyRuleError(
            string? behaviorName,

            string? detail,

            string? errorLocation,

            string? instance,

            int? statusCode,

            string? title,

            string? type)
        {
            BehaviorName = behaviorName;
            Detail = detail;
            ErrorLocation = errorLocation;
            Instance = instance;
            StatusCode = statusCode;
            Title = title;
            Type = type;
        }
    }
}
