// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (Optional) Whether the match is case sensitive.
        /// </summary>
        [Input("caseSensitive")]
        public bool? CaseSensitive { get; set; }

        /// <summary>
        /// - (Optional) For `clientip`, `continent`, `countrycode`, `proxy`, and `regioncode` match types, this defines the part of the request that determines the IP address to use. Values include the connecting IP address (`CONNECTING_IP`) and the X_Forwarded_For header (`XFF_HEADERS`). To select both, enter the two values separated by a space delimiter. When both values are included, the connecting IP address is evaluated first.
        /// </summary>
        [Input("checkIps")]
        public string? CheckIps { get; set; }

        /// <summary>
        /// - (Optional) Compares a string expression with a pattern, either `contains`, `exists`, or `equals`.
        /// </summary>
        [Input("matchOperator")]
        public string? MatchOperator { get; set; }

        /// <summary>
        /// - (Optional) The type of match used, either `header`, `hostname`, `path`, `extension`, `query`, `cookie`, `deviceCharacteristics`, `clientip`, `continent`, `countrycode`, `regioncode`, `protocol`, `method`, or `proxy`.
        /// </summary>
        [Input("matchType")]
        public string? MatchType { get; set; }

        /// <summary>
        /// - (Optional) This depends on the `match_type`. If the `match_type` is `hostname`, then `match_value` is the fully qualified domain name, like `www.akamai.com`.
        /// </summary>
        [Input("matchValue")]
        public string? MatchValue { get; set; }

        /// <summary>
        /// - (Optional) Whether to negate the match.
        /// </summary>
        [Input("negate")]
        public bool? Negate { get; set; }

        [Input("objectMatchValues")]
        private List<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueArgs>? _objectMatchValues;

        /// <summary>
        /// - (Optional) If `match_value` is empty, this argument is required. An object used when a rule includes more complex match criteria, like multiple value attributes. Includes these sub-arguments:
        /// </summary>
        public List<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueArgs> ObjectMatchValues
        {
            get => _objectMatchValues ?? (_objectMatchValues = new List<Inputs.GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchObjectMatchValueArgs>());
            set => _objectMatchValues = value;
        }

        public GetCloudletsPhasedReleaseMatchRuleMatchRuleMatchArgs()
        {
        }
    }
}