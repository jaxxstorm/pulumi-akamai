// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (Optional) Whether the `value` argument should be evaluated based on case sensitivity.
        /// </summary>
        [Input("valueCaseSensitive")]
        public bool? ValueCaseSensitive { get; set; }

        /// <summary>
        /// - (Optional) Whether the `value` argument should be compared in an escaped form.
        /// </summary>
        [Input("valueEscaped")]
        public bool? ValueEscaped { get; set; }

        /// <summary>
        /// - (Optional) Whether the `value` argument includes wildcards.
        /// </summary>
        [Input("valueHasWildcard")]
        public bool? ValueHasWildcard { get; set; }

        [Input("values")]
        private List<string>? _values;

        /// <summary>
        /// - (Optional) If you set the `type` argument to `simple` or `range`, specify the values in the incoming request to match on. With `range`, you can only specify an array of integers, for example `[1, 2]`.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleMatchObjectMatchValueOptionsArgs()
        {
        }
    }
}