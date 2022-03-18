// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleInputArgs : Pulumi.ResourceArgs
    {
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// - (Optional) The end time for this match. Specify the value in UTC in seconds since the epoch.
        /// </summary>
        [Input("end")]
        public Input<int>? End { get; set; }

        [Input("forwardSettings", required: true)]
        private InputList<Inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingInputArgs>? _forwardSettings;

        /// <summary>
        /// - (Required) Defines data used to construct a new request URL if all conditions are met. If all of the conditions you set are true, the Edge Server returns an HTTP response from the rewritten URL.
        /// </summary>
        public InputList<Inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingInputArgs> ForwardSettings
        {
            get => _forwardSettings ?? (_forwardSettings = new InputList<Inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingInputArgs>());
            set => _forwardSettings = value;
        }

        /// <summary>
        /// - (Optional) An identifier for Akamai internal use only.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// - (Optional) The URL that the Cloudlet uses to match the incoming request.
        /// </summary>
        [Input("matchUrl")]
        public Input<string>? MatchUrl { get; set; }

        [Input("matches")]
        private InputList<Inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchInputArgs>? _matches;

        /// <summary>
        /// - (Optional) A list of conditions to apply to a Cloudlet, including:
        /// </summary>
        public InputList<Inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchInputArgs> Matches
        {
            get => _matches ?? (_matches = new InputList<Inputs.GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchInputArgs>());
            set => _matches = value;
        }

        /// <summary>
        /// - (Optional) Whether the match supports default rules that apply to all requests.
        /// </summary>
        [Input("matchesAlways")]
        public Input<bool>? MatchesAlways { get; set; }

        /// <summary>
        /// - (Optional) If you're using a `match_type` that supports name attributes, specify the part the incoming request to match on, either `cookie`, `header`, `parameter`, or `query`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// - (Optional) The start time for this match. Specify the value in UTC in seconds since the epoch.
        /// </summary>
        [Input("start")]
        public Input<int>? Start { get; set; }

        /// <summary>
        /// - (Required) The type of the array, either `object`, `range`, or `simple`. Use the `simple` option when adding only an array of string-based values.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleInputArgs()
        {
        }
    }
}