// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecRatePolicies
    {
        /// <summary>
        /// Use the `akamai.getAppSecRatePolicies` data source to retrieve the rate policies for a specific security configuration, or a single rate policy.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Basic usage:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Akamai = Pulumi.Akamai;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var configuration = Output.Create(Akamai.GetAppSecConfiguration.InvokeAsync(new Akamai.GetAppSecConfigurationArgs
        ///         {
        ///             Name = @var.Security_configuration,
        ///         }));
        ///         var ratePolicies = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecRatePolicies.InvokeAsync(new Akamai.GetAppSecRatePoliciesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.RatePoliciesOutput = ratePolicies.Apply(ratePolicies =&gt; ratePolicies.OutputText);
        ///         this.RatePoliciesJson = ratePolicies.Apply(ratePolicies =&gt; ratePolicies.Json);
        ///         var ratePolicy = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecRatePolicies.InvokeAsync(new Akamai.GetAppSecRatePoliciesArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             RatePolicyId = @var.Rate_policy_id,
        ///         })));
        ///         this.RatePolicyJson = ratePolicy.Apply(ratePolicy =&gt; ratePolicy.Json);
        ///         this.RatePolicyOutput = ratePolicy.Apply(ratePolicy =&gt; ratePolicy.OutputText);
        ///     }
        /// 
        ///     [Output("ratePoliciesOutput")]
        ///     public Output&lt;string&gt; RatePoliciesOutput { get; set; }
        ///     [Output("ratePoliciesJson")]
        ///     public Output&lt;string&gt; RatePoliciesJson { get; set; }
        ///     [Output("ratePolicyJson")]
        ///     public Output&lt;string&gt; RatePolicyJson { get; set; }
        ///     [Output("ratePolicyOutput")]
        ///     public Output&lt;string&gt; RatePolicyOutput { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppSecRatePoliciesResult> InvokeAsync(GetAppSecRatePoliciesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecRatePoliciesResult>("akamai:index/getAppSecRatePolicies:getAppSecRatePolicies", args ?? new GetAppSecRatePoliciesArgs(), options.WithVersion());
    }


    public sealed class GetAppSecRatePoliciesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the security configuration to use.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// The ID of the rate policy to use. If this parameter is not supplied, information about all rate policies will be returned.
        /// </summary>
        [Input("ratePolicyId")]
        public int? RatePolicyId { get; set; }

        public GetAppSecRatePoliciesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecRatePoliciesResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A JSON-formatted list of the rate policy information.
        /// </summary>
        public readonly string Json;
        /// <summary>
        /// A tabular display showing the ID and name of all rate policies associated with the specified security configuration.
        /// </summary>
        public readonly string OutputText;
        public readonly int? RatePolicyId;

        [OutputConstructor]
        private GetAppSecRatePoliciesResult(
            int configId,

            string id,

            string json,

            string outputText,

            int? ratePolicyId)
        {
            ConfigId = configId;
            Id = id;
            Json = json;
            OutputText = outputText;
            RatePolicyId = ratePolicyId;
        }
    }
}
