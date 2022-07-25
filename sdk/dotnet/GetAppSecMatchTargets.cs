// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Akamai
{
    public static class GetAppSecMatchTargets
    {
        /// <summary>
        /// **Scopes**: Security configuration; match target
        /// 
        /// Returns information about your match targets. Match targets determine which security policy should apply to an API, hostname, or path.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/match-targets{?policyId,includeChildObjectName}](https://techdocs.akamai.com/application-security/reference/get-match-targets)
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
        ///         var matchTargetsAppSecMatchTargets = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecMatchTargets.InvokeAsync(new Akamai.GetAppSecMatchTargetsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.MatchTargets = matchTargetsAppSecMatchTargets.Apply(matchTargetsAppSecMatchTargets =&gt; matchTargetsAppSecMatchTargets.OutputText);
        ///         var matchTarget = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecMatchTargets.InvokeAsync(new Akamai.GetAppSecMatchTargetsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             MatchTargetId = 2712938,
        ///         })));
        ///         this.MatchTargetOutput = matchTarget.Apply(matchTarget =&gt; matchTarget.OutputText);
        ///     }
        /// 
        ///     [Output("matchTargets")]
        ///     public Output&lt;string&gt; MatchTargets { get; set; }
        ///     [Output("matchTargetOutput")]
        ///     public Output&lt;string&gt; MatchTargetOutput { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `output_text`. Tabular report showing the ID and security policy ID of your match targets.
        /// - `json`. JSON-formatted list of the match target information.
        /// </summary>
        public static Task<GetAppSecMatchTargetsResult> InvokeAsync(GetAppSecMatchTargetsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppSecMatchTargetsResult>("akamai:index/getAppSecMatchTargets:getAppSecMatchTargets", args ?? new GetAppSecMatchTargetsArgs(), options.WithDefaults());

        /// <summary>
        /// **Scopes**: Security configuration; match target
        /// 
        /// Returns information about your match targets. Match targets determine which security policy should apply to an API, hostname, or path.
        /// 
        /// **Related API Endpoint**: [/appsec/v1/configs/{configId}/versions/{versionNumber}/match-targets{?policyId,includeChildObjectName}](https://techdocs.akamai.com/application-security/reference/get-match-targets)
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
        ///         var matchTargetsAppSecMatchTargets = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecMatchTargets.InvokeAsync(new Akamai.GetAppSecMatchTargetsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///         })));
        ///         this.MatchTargets = matchTargetsAppSecMatchTargets.Apply(matchTargetsAppSecMatchTargets =&gt; matchTargetsAppSecMatchTargets.OutputText);
        ///         var matchTarget = configuration.Apply(configuration =&gt; Output.Create(Akamai.GetAppSecMatchTargets.InvokeAsync(new Akamai.GetAppSecMatchTargetsArgs
        ///         {
        ///             ConfigId = configuration.ConfigId,
        ///             MatchTargetId = 2712938,
        ///         })));
        ///         this.MatchTargetOutput = matchTarget.Apply(matchTarget =&gt; matchTarget.OutputText);
        ///     }
        /// 
        ///     [Output("matchTargets")]
        ///     public Output&lt;string&gt; MatchTargets { get; set; }
        ///     [Output("matchTargetOutput")]
        ///     public Output&lt;string&gt; MatchTargetOutput { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Output Options
        /// 
        /// The following options can be used to determine the information returned, and how that returned information is formatted:
        /// 
        /// - `output_text`. Tabular report showing the ID and security policy ID of your match targets.
        /// - `json`. JSON-formatted list of the match target information.
        /// </summary>
        public static Output<GetAppSecMatchTargetsResult> Invoke(GetAppSecMatchTargetsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppSecMatchTargetsResult>("akamai:index/getAppSecMatchTargets:getAppSecMatchTargets", args ?? new GetAppSecMatchTargetsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppSecMatchTargetsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the match targets.
        /// </summary>
        [Input("configId", required: true)]
        public int ConfigId { get; set; }

        /// <summary>
        /// . Unique identifier of the match target you want to return information for. If not included, information is returned for all your match targets.
        /// </summary>
        [Input("matchTargetId")]
        public int? MatchTargetId { get; set; }

        public GetAppSecMatchTargetsArgs()
        {
        }
    }

    public sealed class GetAppSecMatchTargetsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// . Unique identifier of the security configuration associated with the match targets.
        /// </summary>
        [Input("configId", required: true)]
        public Input<int> ConfigId { get; set; } = null!;

        /// <summary>
        /// . Unique identifier of the match target you want to return information for. If not included, information is returned for all your match targets.
        /// </summary>
        [Input("matchTargetId")]
        public Input<int>? MatchTargetId { get; set; }

        public GetAppSecMatchTargetsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAppSecMatchTargetsResult
    {
        public readonly int ConfigId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly int? MatchTargetId;
        public readonly string OutputText;

        [OutputConstructor]
        private GetAppSecMatchTargetsResult(
            int configId,

            string id,

            string json,

            int? matchTargetId,

            string outputText)
        {
            ConfigId = configId;
            Id = id;
            Json = json;
            MatchTargetId = matchTargetId;
            OutputText = outputText;
        }
    }
}
