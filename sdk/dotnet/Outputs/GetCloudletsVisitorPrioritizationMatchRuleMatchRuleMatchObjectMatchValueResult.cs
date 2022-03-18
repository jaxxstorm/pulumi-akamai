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
    public sealed class GetCloudletsVisitorPrioritizationMatchRuleMatchRuleMatchObjectMatchValueResult
    {
        /// <summary>
        /// - (Optional) If you're using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// - (Optional) Whether the `name` argument should be evaluated based on case sensitivity.
        /// </summary>
        public readonly bool? NameCaseSensitive;
        /// <summary>
        /// - (Optional) Whether the `name` argument includes wildcards.
        /// </summary>
        public readonly bool? NameHasWildcard;
        /// <summary>
        /// - (Optional) If you set the `type` argument to `object`, use this array to list the values to match on.
        /// </summary>
        public readonly Outputs.GetCloudletsVisitorPrioritizationMatchRuleMatchRuleMatchObjectMatchValueOptionsResult? Options;
        /// <summary>
        /// - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// - (Optional) If you set the `type` argument to `simple`, specify the values in the incoming request to match on.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetCloudletsVisitorPrioritizationMatchRuleMatchRuleMatchObjectMatchValueResult(
            string? name,

            bool? nameCaseSensitive,

            bool? nameHasWildcard,

            Outputs.GetCloudletsVisitorPrioritizationMatchRuleMatchRuleMatchObjectMatchValueOptionsResult? options,

            string type,

            ImmutableArray<string> values)
        {
            Name = name;
            NameCaseSensitive = nameCaseSensitive;
            NameHasWildcard = nameHasWildcard;
            Options = options;
            Type = type;
            Values = values;
        }
    }
}
