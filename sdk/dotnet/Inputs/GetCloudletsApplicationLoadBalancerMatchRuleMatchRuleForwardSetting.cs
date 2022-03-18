// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai.Inputs
{

    public sealed class GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// - (Required) The ID of the Conditional Origin the requests are forwarded to.
        /// </summary>
        [Input("originId", required: true)]
        public string OriginId { get; set; } = null!;

        public GetCloudletsApplicationLoadBalancerMatchRuleMatchRuleForwardSettingArgs()
        {
        }
    }
}