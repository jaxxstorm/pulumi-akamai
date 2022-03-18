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
    public sealed class GetCloudletsPhasedReleaseMatchRuleMatchRuleResult
    {
        /// <summary>
        /// - (Optional) Whether to disable a rule so it is not evaluated against incoming requests.
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
        /// </summary>
        public readonly int? End;
        /// <summary>
        /// (Required) The data used to construct a new request URL if all match conditions are met. If all conditions are met, the edge server returns an HTTP response from the rewritten URL.
        /// </summary>
        public readonly Outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsResult ForwardSettings;
        /// <summary>
        /// - (Optional) If you're using a URL match, this specifies the URL that the Cloudlet uses to match the incoming request.
        /// </summary>
        public readonly string? MatchUrl;
        /// <summary>
        /// - (Optional) A list of conditions to apply to a Cloudlet, including:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchResult> Matches;
        /// <summary>
        /// - (Optional) Whether the match supports default rules that apply to all requests.
        /// </summary>
        public readonly bool? MatchesAlways;
        /// <summary>
        /// - (Optional) If you're using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
        /// </summary>
        public readonly int? Start;
        /// <summary>
        /// - (Required) The type of the array, either `object` or `simple`. Use the `simple` option when adding only an array of string-based values.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCloudletsPhasedReleaseMatchRuleMatchRuleResult(
            bool? disabled,

            int? end,

            Outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleForwardSettingsResult forwardSettings,

            string? matchUrl,

            ImmutableArray<Outputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchResult> matches,

            bool? matchesAlways,

            string? name,

            int? start,

            string type)
        {
            Disabled = disabled;
            End = end;
            ForwardSettings = forwardSettings;
            MatchUrl = matchUrl;
            Matches = matches;
            MatchesAlways = matchesAlways;
            Name = name;
            Start = start;
            Type = type;
        }
    }
}
