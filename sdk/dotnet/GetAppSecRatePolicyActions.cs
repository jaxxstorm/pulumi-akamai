// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecRatePolicyActions
    {
        /// <summary>
        /// **Scopes**: Security policy; rate policy
        /// 
        /// Returns information about your rate policy actions. Actions specify what happens any time a rate policy is triggered: the issue could be ignored, the request could be denied, or an alert could be generated.
        /// 
        /// **Related API Endpoint:** [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rate-policies](https://techdocs.akamai.com/application-security/reference/get-rate-policies-actions)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var ratePolicyActionsAppSecRatePolicyActions = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecRatePolicyActions.InvokeAsync(new Akamai.GetAppSecRatePolicyActionsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.RatePolicyActions = ratePolicyActionsAppSecRatePolicyActions.Apply(ratePolicyActionsAppSecRatePolicyActions =&gt; ratePolicyActionsAppSecRatePolicyActions.OutputText);
        ///     }
        /// 
        ///     [Output("ratePolicyActions")]
        ///     public Output&lt;string&gt; RatePolicyActions { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `output_text`. Tabular report showing the ID, IPv4 action, and IPv6 action of the rate policies.
        /// </summary>
        public static Task<GetAppSecRatePolicyActionsResult> InvokeAsync(GetAppSecRatePolicyActionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecRatePolicyActionsResult>("akamai:index/getAppSecRatePolicyActions:getAppSecRatePolicyActions", args ?? new GetAppSecRatePolicyActionsArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security policy; rate policy
        /// 
        /// Returns information about your rate policy actions. Actions specify what happens any time a rate policy is triggered: the issue could be ignored, the request could be denied, or an alert could be generated.
        /// 
        /// **Related API Endpoint:** [/appsec/v1/configs/{configId}/versions/{versionNumber}/security-policies/{policyId}/rate-policies](https://techdocs.akamai.com/application-security/reference/get-rate-policies-actions)
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
        ///             Name = "Documentation",
        ///         }));
        ///         var ratePolicyActionsAppSecRatePolicyActions = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecRatePolicyActions.InvokeAsync(new Akamai.GetAppSecRatePolicyActionsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             SecurityPolicyId = "gms1_134637",
        ///         })));
        ///         this.RatePolicyActions = ratePolicyActionsAppSecRatePolicyActions.Apply(ratePolicyActionsAppSecRatePolicyActions =&gt; ratePolicyActionsAppSecRatePolicyActions.OutputText);
        ///     }
        /// 
        ///     [Output("ratePolicyActions")]
        ///     public Output&lt;string&gt; RatePolicyActions { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `output_text`. Tabular report showing the ID, IPv4 action, and IPv6 action of the rate policies.
        /// </summary>
        public static Output<GetAppSecRatePolicyActionsResult> Invoke(GetAppSecRatePolicyActionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecRatePolicyActionsResult>("akamai:index/getAppSecRatePolicyActions:getAppSecRatePolicyActions", args ?? new GetAppSecRatePolicyActionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecRatePolicyActionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the rate policies and rate policy actions.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the rate policy you want to return action information for. If not included, action information is returned for all your rate policies.
        /// </summary>
        [Input("ratePolicyId")]
        public int? RatePolicyId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the rate policies and rate policy actions.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public string SecurityPolicyId { get; set; } = null!;

        public GetAppSecRatePolicyActionsArgs()
        {
        }
    }

    public sealed class GetAppSecRatePolicyActionsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the rate policies and rate policy actions.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the rate policy you want to return action information for. If not included, action information is returned for all your rate policies.
        /// </summary>
        [Input("ratePolicyId")]
        public Input<int>? RatePolicyId { get; set; }

        /// <summary>
        /// . Unique identifier of the security policy associated with the rate policies and rate policy actions.
        /// </summary>
        [Input("securityPolicyId", required: true)]
        public Input<string> SecurityPolicyId { get; set; } = null!;

        public GetAppSecRatePolicyActionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecRatePolicyActionsResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string OutputText;
        public readonly int? RatePolicyId;
        public readonly string SecurityPolicyId;

        [OutputConstructor]
        private GetAppSecRatePolicyActionsResult(
            int configId,

            string id,

            string outputText,

            int? ratePolicyId,

            string securityPolicyId)
        {
            ConfigId = configId;
            Id = id;
            OutputText = outputText;
            RatePolicyId = ratePolicyId;
            SecurityPolicyId = securityPolicyId;
        }
    }
}
